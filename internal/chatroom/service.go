package chatroom

import (
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/port"
	userPort "github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/port"
)

type service struct {
	repo     port.Repo
	userPort userPort.Service
}
