package handler

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/pgtype"
)

type User struct {
	ID           string
	Conn         *websocket.Conn
	ChannelID    int
	PodchannelID int
}

type EventMessage struct {
	Event        string `json:"event"`
	Data         string `json:"data"`
	ChannelID    int    `json:"channel_id,omitempty"`
	PodchannelID int    `json:"podchannel_id,omitempty"`
	AuthorID     string `json:"author_id,omitempty"`
}

type Message struct {
	Content      string    `json:"content"`
	AuthorID     int       `json:"author_id"`
	PodchannelID int       `json:"podchannel_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type Channel struct {
	ID        int32            `db:"id" json:"id"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	Name      string           `db:"name" json:"name"`
}
type SwaggerChannel struct {
	Channel
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type SwaggerChannel struct {
	ID        int32     `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `db:"name" json:"name"`
}
