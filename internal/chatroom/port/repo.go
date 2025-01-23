package port

import (
	"context"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/domain"
)

type Repo interface {
	SendMessage(ctx context.Context, message *domain.MessageToSend) error
	SubscribeUser(ctx context.Context, UserID string) error
	ShowMessages(ctx context.Context, UserID string) (*domain.MessagesToShow, error)
	GetActiveUsers(ctx context.Context) (*domain.ActiveUsers, error)
}
