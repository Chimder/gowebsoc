package model

import (
	"time"

	"gorm.io/gorm"
)

type Timestamps struct {
	CreatedAt time.Time  `db:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `db:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `db:"deleted_at" gorm:"column:deleted_at"`
}
type User struct {
	gorm.Model
	Name     string    `db:"name" gorm:"column:name;unique;not null"`
	Email    string    `db:"email" gorm:"column:email;unique;not null"`
	Messages []Message `gorm:"foreignKey:AuthorID"`
}

type Channel struct {
	gorm.Model
	Name        string       `db:"name" gorm:"column:name;not null;unique"`
	CreatedAt   time.Time    `db:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time    `db:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *time.Time   `db:"deleted_at" gorm:"column:deleted_at"`
	PodChannels []Podchannel `db:"podchannels" gorm:"foreignKey:ChannelID"`
}

type Podchannel struct {
	gorm.Model
	Name      string     `db:"name" gorm:"column:name;not null"`
	Type      string     `db:"type" gorm:"column:type;default:text"`
	CreatedAt time.Time  `db:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `db:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `db:"deleted_at" gorm:"column:deleted_at"`
	ChannelID uint       `db:"channel_id" gorm:"column:channel_id;not null"`
	Channel   Channel    `db:"channel" gorm:"foreignKey:ChannelID"`
	Messages  []Message  `db:"name" gorm:"foreignKey:PodchannelID"`
}

type Message struct {
	gorm.Model
	Content      string     `gorm:"column:content;not null"`
	AuthorID     uint       `gorm:"column:author_id;not null"`
	Author       User       `gorm:"foreignKey:AuthorID"`
	PodchannelID uint       `gorm:"column:podchannel_id;not null"`
	Podchannel   Podchannel `gorm:"foreignKey:PodchannelID"`
}
