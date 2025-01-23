package storage

import (
	"context"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/domain"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/port"
	"gorm.io/gorm"
)

type chatroomRepo struct {
	db *gorm.DB
}

// GetActiveUsers implements port.Repo.
func NewChatroomRepo(db *gorm.DB) port.Repo {
	return &chatroomRepo{
		db: db,
	}
}
func (c *chatroomRepo) GetActiveUsers(ctx context.Context) (*domain.ActiveUsers, error) {
	panic("unimplemented")
}

// SendMessage implements port.Repo.
func (c *chatroomRepo) SendMessage(ctx context.Context, message *domain.MessageToSend) error {
	panic("unimplemented")
}

// ShowMessages implements port.Repo.
func (c *chatroomRepo) ShowMessages(ctx context.Context, UserID string) (*domain.MessagesToShow, error) {
	panic("unimplemented")
}

// SubscribeUser implements port.Repo.
func (c *chatroomRepo) SubscribeUser(ctx context.Context, UserID string) error {
	panic("unimplemented")
}
