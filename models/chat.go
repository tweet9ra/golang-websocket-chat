package models

import (
	"time"
)

type Chat struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name string `json:"name"`
	Users []User `gorm:"many2many:user_chats" json:"users"`
	Messages []*ChatMessage `json:"messages"`
}