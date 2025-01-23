package app

import (
	"context"
	"fmt"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/config"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom"
	chatroomPort "github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/port"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user"
	userPort "github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/port"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/adapters/nats"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/adapters/storage"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/ports"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/postgres"
	"gorm.io/gorm"

	appCtx "github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/context"
)

type app struct {
	db              *gorm.DB
	cfg             config.Config
	userService     userPort.Service
	chatroomService chatroomPort.Service
	messageBroker   ports.IMessageBroker
	// redisProvider   cache.Provider
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}

	return a.userServiceWithDB(db)
}

func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db))
}

func (a *app) chatroomServiceWithDB(db *gorm.DB) chatroomPort.Service {
	return chatroom.NewChatroomService(storage.NewChatroomRepo(db))
	// user.NewService(storage.NewUserRepo(db)))
}

func (a *app) ChatroomService(ctx context.Context) chatroomPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.chatroomService == nil {
			a.chatroomService = a.chatroomServiceWithDB(a.db)
		}
		return a.chatroomService
	}

	return a.chatroomServiceWithDB(db)
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db

	postgres.AddUuidExtension(db)
	postgres.GormMigrations(db)
	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}
	a.setMessageBroker()

	return a, nil
}
func (a *app) MessageBroker() ports.IMessageBroker {
	return a.messageBroker
}

func (a *app) setMessageBroker() {
	natsCfg := a.cfg.Nats
	if a.messageBroker != nil {
		return
	}
	nats, err := nats.NewNATS(fmt.Sprintf("%s:%d", natsCfg.Host, natsCfg.Port))
	if err != nil {
		return
	}
	a.messageBroker = nats

}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
