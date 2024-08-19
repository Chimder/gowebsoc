package handler

import (
	"context"
	"encoding/json"
	"goSql/internal/queries"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Server struct {
	usersByChannel map[int]map[string]*User
	broadcast      chan *EventMessage
	register       chan *User
	unregister     chan *User
	mu             sync.Mutex
	sqlc           *queries.Queries
	rdb            *redis.Client
}

func NewWebServer(sqlc *queries.Queries, rdb *redis.Client) *Server {
	return &Server{
		usersByChannel: make(map[int]map[string]*User),
		broadcast:      make(chan *EventMessage),
		register:       make(chan *User),
		unregister:     make(chan *User),
		sqlc:           sqlc,
		rdb:            rdb,
	}
}

func (ws *Server) Run() {
	go func() {
		for {
			select {
			case user := <-ws.register:
				log.Printf("Register")
				ws.mu.Lock()
				if _, ok := ws.usersByChannel[user.ChannelID]; !ok {
					ws.usersByChannel[user.ChannelID] = make(map[string]*User)
				}
				ws.usersByChannel[user.ChannelID][user.ID] = user
				ws.mu.Unlock()

			case user := <-ws.unregister:
				log.Printf("Unregister")
				ws.mu.Lock()
				if users, ok := ws.usersByChannel[user.ChannelID]; ok {
					delete(users, user.ID)
					if len(users) == 0 {
						delete(ws.usersByChannel, user.ChannelID)
					}
				}
				ws.mu.Unlock()

			case message := <-ws.broadcast:
				ws.mu.Lock()
				if users, ok := ws.usersByChannel[message.ChannelID]; ok {
					for _, user := range users {
						log.Printf("Send", message)
						if err := user.Conn.WriteJSON(message); err != nil {
							log.Printf("Write error: %s\n", err)
						}
					}
				}
				ws.mu.Unlock()

				messageData, err := json.Marshal(message)
				if err != nil {
					log.Printf("JSON marshal error: %s\n", err)
					continue
				}

				if err := ws.rdb.RPush(context.Background(), "messageQueue", messageData).Err(); err != nil {
					log.Printf("Redis push error: %s\n", err)
					return
				}
			}
		}
	}()
	go ProcessMessages(ws.sqlc, ws.rdb)
}

func (ws *Server) WsConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		http.Error(w, "Could not connect to websocket", http.StatusInternalServerError)
		return
	}

	userID := uuid.New().String()
	defer conn.Close()

	var user *User

	for {
		var eventMessage EventMessage
		if err := conn.ReadJSON(&eventMessage); err != nil {
			log.Printf("Read error: %s\n", err)
			if user != nil {
				ws.unregister <- user
			}
			break
		}

		if user == nil {
			user = &User{
				ID:        userID,
				Conn:      conn,
				ChannelID: eventMessage.ChannelID,
			}
			ws.register <- user
		}

		if eventMessage.Event == "join_podchannel" {
			ws.unregister <- user
			user = nil
			continue
		}
		log.Printf("Received message: %+v\n", eventMessage)
		ws.broadcast <- &EventMessage{
			AuthorID:     userID,
			MessageID:    eventMessage.MessageID,
			Message:      eventMessage.Message,
			Event:        eventMessage.Event,
			CreatedAt:    eventMessage.CreatedAt,
			ChannelID:    eventMessage.ChannelID,
			PodchannelID: eventMessage.PodchannelID,
		}

	}
}
