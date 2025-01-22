package domain

import (
	userDomain "github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
)

type Chatroom struct {
	users       []userDomain.User
	activeUsers []userDomain.User
}
