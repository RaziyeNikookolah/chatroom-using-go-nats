package user

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/port"
	"github.com/google/uuid"
)

var (
	ErrUserOnCreate           = errors.New("error on creating new user")
	ErrUserCreationValidation = errors.New("validation failed")
	ErrUserNotFound           = errors.New("user not found")
)

type service struct {
	repo port.Repo
}

// GetUserByUUID implements port.Service.
func (s *service) GetUserByUUID(ctx context.Context, userID *domain.UserID) (*domain.User, error) {
	user, err := s.repo.GetByUUID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(ctx context.Context, user domain.User) (domain.UserID, error) {
	if err := user.Validate(); err != nil {
		return domain.UserID(uuid.Nil), fmt.Errorf("%w %w", ErrUserCreationValidation, err)
	}

	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		log.Println("error on creating new user : ", err.Error())
		return domain.UserID(uuid.Nil), ErrUserOnCreate
	}

	return userID, nil
}

func (s *service) GetUserByUsernamePassword(ctx context.Context, username domain.Username, password domain.Password) (*domain.User, error) {
	user, err := s.repo.GetUserByUsernamePassword(ctx, username, password)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
