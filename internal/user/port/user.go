package port

import (
	"context"
	"errors"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
)

var (
	ErrUserAlreadyExist  = errors.New("user already exists")
	ErrInvalidCredential = errors.New("invalid username or password")
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserID, error)
	GetByUUID(ctx context.Context, uuid *domain.UserID) (*domain.User, error)
	GetUserByUsernamePassword(ctx context.Context, username domain.Username, password domain.Password) (*domain.User, error)
}
