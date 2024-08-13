package handler

import (
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	ID           string
	Conn         *websocket.Conn
	ChannelID    int
	PodchannelID int
}

type EventMessage struct {
	Event        string      `json:"event"`
	Message      interface{} `json:"message"`
	CreatedAt    time.Time   `json:"created_at"`
	ChannelID    int         `json:"channel_id,omitempty"`
	PodchannelID int         `json:"podchannel_id,omitempty"`
	AuthorID     string      `json:"author_id,omitempty"`
}

// type Message struct {
// 	Content      string    `json:"content"`
// 	AuthorID     int       `json:"author_id"`
// 	PodchannelID int       `json:"podchannel_id"`
// 	CreatedAt    time.Time `json:"created_at"`
// }
