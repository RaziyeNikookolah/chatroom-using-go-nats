package port

import (
	"context"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/jwt"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	GetUserByUUID(ctx context.Context, userID *domain.UserID) (*domain.User, error)
	GetUserClaimWithToken(ctx context.Context, token string, secret string) (*jwt.UserClaims, error)
	GetUserByUsernamePassword(ctx context.Context, username domain.Username, password domain.Password) (*domain.User, error)
}
