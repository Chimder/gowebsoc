package main

import (
	"time"

	"github.com/uptrace/bun"
)

type Timestamps struct {
	CreatedAt time.Time  `bun:"created_at"`
	UpdatedAt time.Time  `bun:"updated_at"`
	DeletedAt *time.Time `bun:"deleted_at"`
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int64     `bun:",pk,autoincrement"`
	Name     string    `bun:"name,unique,notnull"`
	Email    string    `bun:"email,unique,notnull"`
	Messages []Message `bun:"rel:has-many,join:id=author_id"`
}

type Channel struct {
	bun.BaseModel `bun:"table:channels,alias:c"`
	ID            int64        `bun:",pk,autoincrement"`
	Name          string       `bun:"name,unique,notnull"`
	PodChannels   []Podchannel `bun:"rel:has-many,join:id=channel_id"`
}

type Podchannel struct {
	bun.BaseModel `bun:"table:podchannels,alias:pc"`
	ID            int64     `bun:",pk,autoincrement"`
	Name          string    `bun:"name,notnull"`
	Type          string    `bun:"type,default:'text'"`
	ChannelID     int64     `bun:"channel_id,notnull"`
	Channel       Channel   `bun:"rel:belongs-to,join:channel_id=id"`
	Messages      []Message `bun:"rel:has-many,join:id=podchannel_id"`
}

type Message struct {
	bun.BaseModel `bun:"table:messages,alias:m"`
	ID            int64      `bun:",pk,autoincrement"`
	Content       string     `bun:"content,notnull"`
	AuthorID      int64      `bun:"author_id,notnull"`
	Author        User       `bun:"rel:belongs-to,join:author_id=id"`
	PodchannelID  int64      `bun:"podchannel_id,notnull"`
	Podchannel    Podchannel `bun:"rel:belongs-to,join:podchannel_id=id"`
	Timestamps
}
