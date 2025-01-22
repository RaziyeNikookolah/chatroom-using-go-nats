package storage

import (
	"context"
	"errors"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/port"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/adapters/storage/mapper"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/adapters/storage/types"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.Repo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, userDomain domain.User) (domain.UserID, error) {
	user := mapper.UserDomain2Storage(userDomain)
	return domain.UserID(user.ID), r.db.Table("users").WithContext(ctx).Create(user).Error
}

func (r *userRepo) GetByUUID(ctx context.Context, userID *domain.UserID) (*domain.User, error) {
	var user types.User

	// Query the users table using the provided UUID
	q := r.db.Table("users").Debug().WithContext(ctx).Where("uuid = ?", userID.ConvStr())

	err := q.First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}
func (r *userRepo) GetUserByUsernamePassword(ctx context.Context, username domain.Username, password domain.Password) (*domain.User, error) {
	var user types.User

	// Query the users table using the provided UUID
	q := r.db.Table("users").Debug().WithContext(ctx).Where("username = ?", username).Where("password=?", password)

	err := q.First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}
