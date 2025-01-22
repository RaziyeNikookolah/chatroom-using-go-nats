package port

import (
	"context"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/domain"
)

type Service interface {
	Send(ctx context.Context, chatroom *domain.Chatroom) error
}
