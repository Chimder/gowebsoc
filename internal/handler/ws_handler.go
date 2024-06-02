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
	broadcast  chan *string
	register   chan *User
	unregister chan *User
	mu         sync.Mutex
}

func NewWebServer() *Server {
	return &Server{
		users:      make(map[string]*User),
		broadcast:  make(chan *string),
		register:   make(chan *User),
		unregister: make(chan *User),
	}
}

func (s *Server) Run() {
	go func() {
		for {
			select {
			case user := <-s.register:
				s.mu.Lock()
				s.users[user.ID] = user
				s.mu.Unlock()
				log.Printf("User %s registered", user.ID)

			case user := <-s.unregister:
				s.mu.Lock()
				if _, ok := s.users[user.ID]; ok {
					delete(s.users, user.ID)
					user.Conn.Close()
					log.Printf("User %s unregistered", user.ID)
				}
				s.mu.Unlock()

			case message := <-s.broadcast:
				s.handleMessage(message)
			}
		}
	}()
}

func (s *Server) handleMessage(mess string) {
	s.mu.Lock()

	for _, user := range s.users {
		if err := user.Conn.WriteJSON(mess); err != nil {
			log.Printf("Error sending message to user %s: %v", user.ID, err)
			user.Conn.Close()
			delete(s.users, user.ID)
		}
	}
	s.mu.Unlock()
}

func (s *Server) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "could not open websocket connection: "+err.Error(), http.StatusBadRequest)
		return
	}

	user := &User{
		ID:   uuid.New().String(),
		Conn: conn,
	}
	s.register <- user

	log.Printf("user conn")
	err = conn.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	go func() {
		defer func() {
			s.unregister <- user
		}()
		for {
			// var eventMessage EventMessage
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Printf("error reading JSON message: %v", err)
				break
			}

			log.Println(string(p))
			s.broadcast <- string(p)
		}
	}()
}
