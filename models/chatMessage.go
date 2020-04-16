package models

import (
	"time"
)

type ChatMessage struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	ChatID uint `json:"chat_id,omitempty"`
	UserID uint `json:"user_id,omitempty"`
	Text string `json:"text"`
	User *User `gorm:"foreignkey:UserID" json:"user,omitempty"`
	Chat *Chat `gorm:"foreignkey:ChatID" json:"chat,omitempty"`
}