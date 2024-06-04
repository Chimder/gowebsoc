package handler

import (
	"github.com/gorilla/websocket"
)

type User struct {
	ID   string
	Conn *websocket.Conn
	Send chan string
}

type EventMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

// func (e *EventMessage) DecodeData(target interface{}) error {
// 	dataBytes, err := json.Marshal(e.Data)
// 	if err != nil {
// 		return err
// 	}
// 	return json.Unmarshal(dataBytes, target)
// }
