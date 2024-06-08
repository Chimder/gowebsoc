package handler

import (
	"github.com/gorilla/websocket"
)

type User struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
}

type EventMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

