package app

import (
	"context"

	"github.com/chatroom/chat/config"
	"gorm.io/gorm"
	// userPort "github.com/chatroom/chat/internal/user/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	// UserService(ctx context.Context) userPort.Service
}
