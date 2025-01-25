package chatroom

import (
	"context"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/domain"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/port"
)

type service struct {
	repo port.Repo
	port port.Service
}

func NewChatroomService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

// implements port.Service.
func (s *service) SubscribeUser(ctx context.Context, userID string) error {
	return s.repo.SubscribeUser(ctx, userID)
}

// GetUserByUUID implements port.Service.
func (s *service) SendMessage(ctx context.Context, message *domain.MessageToSend) error {
	return s.repo.SendMessage(ctx, message)
}

func (s *service) ShowMessages(ctx context.Context, userID string) (*domain.MessagesToShow, error) {
	return s.repo.ShowMessages(ctx, userID)
}

func (s *service) GetActiveUsers(ctx context.Context) (*domain.ActiveUsers, error) {
	return s.repo.GetActiveUsers(ctx)
}
