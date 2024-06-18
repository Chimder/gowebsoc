package handler

import (
	"context"
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
	users      map[string]*User
	broadcast  chan *EventMessage
	register   chan *User
	unregister chan *User
	mu         sync.Mutex
	rdb        *redis.Client
}

func NewWebServer() *Server {
	return &Server{
		users:      make(map[string]*User),
		broadcast:  make(chan *EventMessage),
		register:   make(chan *User),
		unregister: make(chan *User),
	}
}

func (ws *Server) Run() {
	go func() {
		for {
			select {
			case user := <-ws.register:
				ws.mu.Lock()
				ws.users[user.ID] = user
				ws.broadcastUserList()
				ws.mu.Unlock()

			case user := <-ws.unregister:
				ws.mu.Lock()
				if _, ok := ws.users[user.ID]; ok {
					delete(ws.users, user.ID)
					ws.broadcastUserList()
				}
				ws.mu.Unlock()

			case message := <-ws.broadcast:
				ws.mu.Lock()
				for _, user := range ws.users {
					if user.PodchannelID == message.PodchannelID {
						if err := user.Conn.WriteJSON(message); err != nil {
							log.Printf("Write error: %s\n", err)
						}
					}
				}
				ws.mu.Unlock()
				if err := ws.rdb.RPush(context.Background(), "messageQueue", message).Err(); err != nil {
					log.Printf("Redis push error: %s\n", err)
				}

			}
		}
	}()
	go processMessages(context.Background(), ws.pgdb, ws.redisClient)
}

func (ws *Server) broadcastUserList() {
	message := EventMessage{
		Event: "users",
		Data:  map[string]int{"users": len(ws.users)},
	}

	for _, user := range ws.users {
		if err := user.Conn.WriteJSON(&message); err != nil {
			log.Printf("Error broadcasting user list update: %s\n", err)
		}
	}
}

func (ws *Server) WsConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		http.Error(w, "Could not connect to websocket", http.StatusInternalServerError)
		return
	}

	userID := uuid.New().String()
	user := &User{ID: userID, Conn: conn}
	ws.register <- user

	defer func() {
		ws.unregister <- user
		conn.Close()
	}()

	for {
		var eventMessage EventMessage
		if err := conn.ReadJSON(&eventMessage); err != nil {
			log.Printf("Read error: %s\n", err)
			break
		}

		log.Println("MESS", eventMessage)
		// Handle setting channel and podchannel
		if eventMessage.Event == "join_podchannel" {
			user.ChannelID = eventMessage.ChannelID
			user.PodchannelID = eventMessage.PodchannelID
			continue
		}

		ws.broadcast <- &EventMessage{
			Event:        eventMessage.Event,
			Data:         eventMessage.Data,
			ChannelID:    user.ChannelID,
			PodchannelID: user.PodchannelID,
		}
	}
}
