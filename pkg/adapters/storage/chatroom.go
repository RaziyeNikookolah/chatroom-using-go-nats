package storage

import (
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/port"
	"gorm.io/gorm"
)

type chatroomRepo struct {
	db *gorm.DB
}

func NewChatroomRepo(db *gorm.DB) port.Repo {
	return &chatroomRepo{
		db: db,
	}
}
