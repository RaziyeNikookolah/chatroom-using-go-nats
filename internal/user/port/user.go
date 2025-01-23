package port

import (
	"context"
	"errors"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/jwt"
)

var (
	ErrUserAlreadyExist  = errors.New("user already exists")
	ErrInvalidCredential = errors.New("invalid username or password")
	ErrInvalidToken      = errors.New("invalid token")
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserID, error)
	GetByUUID(ctx context.Context, uuid *domain.UserID) (*domain.User, error)
	GetUserClaimWithToken(ctx context.Context, token string, secret string) (*jwt.UserClaims, error)
	GetUserByUsernamePassword(ctx context.Context, username domain.Username, password domain.Password) (*domain.User, error)
}
