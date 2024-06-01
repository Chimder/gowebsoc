package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


type WsHandler struct {
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
func NewWsHandler() *WsHandler {
	return &WsHandler{}
}

func (ws *WsHandler) WsConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		http.Error(w, "Could not connect to websocket", http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("WebSocket close error: %v", err)
		}
		log.Println("Client disconnected")
	}()

	/////////////////////////////
	log.Println("Client Connected")
	err = conn.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	///////////////////
	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		// print out that message for clarity
		log.Println(string(p))
		log.Println(string(p))

		// if err := ws.WriteMessage(messageType, p); err != nil {
		// 	log.Println("Write error:", err)
		// 	break
		// }

	}
}
