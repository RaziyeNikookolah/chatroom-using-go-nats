package ports

import (
	"context"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/domain"
)

type IMessageBroker interface {
	Publish(ctx context.Context, username, msg string) string
	SubscribeToChat(ctx context.Context, username string)
	GetAllMessages(username string) ([]domain.Message, error)
	GetActiveUsers() ([]string, error)
	ConsumeUnreadMessages(consumerName string) ([]string, error)
}
