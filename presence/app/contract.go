package app

import (
	"context"

	"github.com/chatroom/presence/config"
	"gorm.io/gorm"
	// userPort "github.com/chatroom/presence/internal/user/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	// UserService(ctx context.Context) userPort.Service
}
