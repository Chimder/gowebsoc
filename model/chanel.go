package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `gorm:"column:name;unique;not null"`
	Email    string    `gorm:"column:email;unique;not null"`
	Messages []Message `gorm:"foreignKey:AuthorID"`
}

type Channel struct {
	gorm.Model
	Name        string       `gorm:"column:name;not null;unique"`
	PodChannels []Podchannel `gorm:"foreignKey:ChannelID"`
}

type Podchannel struct {
	gorm.Model
	Name      string    `gorm:"column:name;not null"`
	Type      string    `gorm:"column:type;default:text"`
	ChannelID uint      `gorm:"column:channel_id;not null"`
	Channel   Channel   `gorm:"foreignKey:ChannelID"`
	Messages  []Message `gorm:"foreignKey:PodchannelID"`
}

type Message struct {
	gorm.Model
	Content      string     `gorm:"column:content;not null"`
	AuthorID     uint       `gorm:"column:author_id;not null"`
	Author       User       `gorm:"foreignKey:AuthorID"`
	PodchannelID uint       `gorm:"column:podchannel_id;not null"`
	Podchannel   Podchannel `gorm:"foreignKey:PodchannelID"`
}
