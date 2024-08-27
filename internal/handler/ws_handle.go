package handler

import (
	"context"
	"encoding/json"
	"goSql/internal/queries"
	"log"
	"net/http"
	"sync"
	"time"

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
	clients   map[string]*User
	broadcast chan *EventMessage
	mu        sync.RWMutex
	sqlc      *queries.Queries
	rdb       *redis.Client
}

type User struct {
	ID           string
	Conn         *websocket.Conn
	ChannelID    int
	PodchannelID int
}

type EventMessage struct {
	Event        string      `json:"event"`
	Message      interface{} `json:"message"`
	MessageID    string      `json:"message_id"`
	CreatedAt    time.Time   `json:"created_at"`
	ChannelID    int         `json:"channel_id,omitempty"`
	PodchannelID int         `json:"podchannel_id,omitempty"`
	AuthorID     string      `json:"author_id,omitempty"`
}

func NewWebServer(sqlc *queries.Queries, rdb *redis.Client) *Server {
	return &Server{
		clients:   make(map[string]*User),
		broadcast: make(chan *EventMessage),
		sqlc:      sqlc,
		rdb:       rdb,
	}
}

func (ws *Server) Run() {
	go func() {
		for message := range ws.broadcast {
			ws.mu.RLock()
			for _, user := range ws.clients {
				if err := user.Conn.WriteJSON(message); err != nil {
					log.Printf("Write error to user %s: %s\n", user.ID, err)
				}
			}
			ws.mu.RUnlock()

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
	user := &User{
		ID:   userID,
		Conn: conn,
	}
	ws.mu.Lock()
	ws.clients[userID] = user
	ws.mu.Unlock()

	defer func() {
		ws.mu.Lock()
		delete(ws.clients, userID)
		ws.mu.Unlock()
		conn.Close()
	}()

	for {
		var eventMessage EventMessage
		if err := conn.ReadJSON(&eventMessage); err != nil {
			log.Printf("Read error: %s\n", err)
			break
		}

		eventMessage.AuthorID = userID
		ws.broadcast <- &eventMessage
	}
}
