package app

import (
	"context"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/config"
	chatroomPort "github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/port"
	userPort "github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/port"

	"gorm.io/gorm"
)

type App interface {
	UserService(ctx context.Context) userPort.Service
	ChatroomService(ctx context.Context) chatroomPort.Service
	DB() *gorm.DB
	Config() config.Config
}
