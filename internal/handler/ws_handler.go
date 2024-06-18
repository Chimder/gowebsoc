package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WsHandler struct{}

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
					if err := user.Conn.WriteJSON(message); err != nil {
						log.Printf("Write error: %s\n", err)
					}
				}
				ws.mu.Unlock()
			}
		}
	}()
}

func (ws *Server) broadcastUserList() {
	message := EventMessage{
		Event: "user_list",
		Data:  map[string]int{"connected_users": len(ws.users)},
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

		log.Println("all", eventMessage)
		log.Println("event", eventMessage.Event)
		log.Println("mess", eventMessage.Data)

		ws.broadcast <- &EventMessage{
			Event: "message",
			Data:  eventMessage.Data,
		}
	}
}
