package app

import (
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/config"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/pkg/ports"
)

type App interface {
	Config() config.Config
	MessageBroker() ports.IMessageBroker
}
