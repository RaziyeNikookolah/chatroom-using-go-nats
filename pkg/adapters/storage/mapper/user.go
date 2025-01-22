package mapper

import (
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/adapters/storage/types"
	"github.com/google/uuid"
)

func UserDomain2Storage(userDomain domain.User) *types.User {
	return &types.User{
		ID:        uuid.UUID(userDomain.ID),
		CreatedAt: userDomain.CreatedAt,
		Username:  string(userDomain.Username),
		Email:     string(userDomain.Email),
		Password:  string(userDomain.Password),
	}
}

func UserStorage2Domain(user types.User) *domain.User {
	return &domain.User{
		ID:        domain.UserID(user.ID),
		CreatedAt: user.CreatedAt,
		Username:  domain.Username(user.Username),
		Email:     domain.Email(user.Email),
		Password:  domain.Password(user.Password),
	}
}
