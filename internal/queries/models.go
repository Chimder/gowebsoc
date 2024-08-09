// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package queries

import (
	"time"
)

type Channel struct {
	ID        int32     `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `db:"name" json:"name"`
}

type Message struct {
	ID           int32     `db:"id" json:"id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Content      string    `db:"content" json:"content"`
	AuthorID     string    `db:"author_id" json:"author_id"`
	PodchannelID int32     `db:"podchannel_id" json:"podchannel_id"`
}

type Podchannel struct {
	ID        int32     `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `db:"name" json:"name"`
	Types     string    `db:"types" json:"types"`
	ChannelID int32     `db:"channel_id" json:"channel_id"`
}

type User struct {
	ID        int32     `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `db:"name" json:"name"`
}
