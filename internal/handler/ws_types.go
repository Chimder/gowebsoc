package handler

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type User struct {
	ID   string
	Conn *websocket.Conn
}

type EventMessage struct {
	// Event string `json:"event"`
	Mess  string `json:"mess"`
}



func (e *EventMessage) DecodeData(target interface{}) error {
	dataBytes, err := json.Marshal(e.Mess)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataBytes, target)
}
