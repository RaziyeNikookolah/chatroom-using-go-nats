package app

import (
	"context"

	"github.com/chatroom/auth/config"
	"gorm.io/gorm"
	// userPort "github.com/chatroom/auth/internal/user/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	// UserService(ctx context.Context) userPort.Service
}
