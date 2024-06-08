package model

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `gorm:"column:;unique;not null"`
	Email    string    `gorm:"column:;unique;not null"`
	Messages []Message `gorm:"foreignKey:AuthorID;references:ID"`
}

type Channel struct {
	gorm.Model
	Name        string       `gorm:"column:name;not null;unique"`
	PodChannels []Podchannel `gorm:"foreignKey:ChannelID;references:ID"`
}

type Podchannel struct {
	gorm.Model
	Name      string  `gorm:"column:name;not null"`
	Type      string  `gorm:"default:text"`
	ChannelID string  `gorm:"column:channel_id;not null"`
	Channel   Channel `gorm:"foreignKey:ChannelID"`
}

type Message struct {
	gorm.Model
	Content   string `gorm:"column:content;not null"`
	AuthorID  string `gorm:"column:author_id;not null"`
	ChannelID string `gorm:"column:channel_id;not null"`
	Author    User   `gorm:"foreignKey:AuthorID;references:ID"`
}
