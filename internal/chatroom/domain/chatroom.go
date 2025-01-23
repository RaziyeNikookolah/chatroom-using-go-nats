package domain

import (
	userDomain "github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
)

type Chatroom struct {
	users       []userDomain.User
	activeUsers []userDomain.User
}
type MessageToSend struct {
	UserID   string
	Username string
	Message  string
}
type MessagesToShow struct {
	Messages []string
}
type ActiveUsers struct {
	Usernames []string
}
