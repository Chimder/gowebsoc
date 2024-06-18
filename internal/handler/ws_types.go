package handler

import (
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	ID           string
	Conn         *websocket.Conn
	ChannelID    string
	PodchannelID string
}

type EventMessage struct {
	Event        string      `json:"event"`
	Data         interface{} `json:"data"`
	ChannelID    string      `json:"channel_id,omitempty"`
	PodchannelID string      `json:"podchannel_id,omitempty"`
}

type Message struct {
	Content      string    `json:"content"`
	AuthorID     int       `json:"author_id"`
	PodchannelID int       `json:"podchannel_id"`
	CreatedAt    time.Time `json:"created_at"`
}
