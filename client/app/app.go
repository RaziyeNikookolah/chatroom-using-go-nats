package app

import (
	"log"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/config"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/pkg/nats"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/pkg/ports"
)

type app struct {
	cfg config.Config

	messageBroker ports.IMessageBroker
	// redisProvider   cache.Provider
}

func (a *app) Config() config.Config {
	return a.cfg
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	a.setMessageBroker()

	return a, nil
}
func (a *app) MessageBroker() ports.IMessageBroker {
	return a.messageBroker
}

func (a *app) setMessageBroker() {
	// natsCfg := a.cfg.Nats
	if a.messageBroker != nil {
		return
	}
	natsClient, err := nats.GetInstance()
	if err != nil {
		log.Fatalf("Error creating NATS Server: %v", err)
	}
	a.messageBroker = natsClient
}
func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
