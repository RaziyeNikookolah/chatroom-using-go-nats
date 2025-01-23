package service

import (
	"context"
	"errors"

	// jwt2 "github.com/golang-jwt/jwt/v5"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
	userPort "github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/port"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/jwt"
	"github.com/google/uuid"
)

type UserService struct {
	svc userPort.Service
	// chatroomSvc chatroomPort.Service
	authSecret string
}

func NewUserService(svc userPort.Service, authSecret string) *UserService {
	//, chatroomSvc chatroomPort.Service) *UserService {
	return &UserService{
		svc:        svc,
		authSecret: authSecret,
		// chatroomSvc: chatroomSvc,
	}
}

var (
	ErrUserCreationValidation = user.ErrUserCreationValidation
	ErrUserOnCreate           = user.ErrUserOnCreate
	ErrUserNotFound           = user.ErrUserNotFound
	ErrInvalidUserPassword    = errors.New("invalid password")
	ErrNoMessage              = errors.New("no message found")
	ErrNoOneIsActiveNow       = errors.New("no one is active now")
)

func (s *UserService) SignUp(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userID, err := s.svc.CreateUser(ctx, domain.User{
		Username: domain.Username(req.GetUsername()),
		Email:    domain.Email(req.GetEmail()),
		Password: domain.NewPassword(req.GetPassword()),
	})

	if err != nil {
		return nil, err
	}

	token, err := jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		// RegisteredClaims: jwt2.RegisteredClaims{
		// 	// ExpiresAt: jwt2.NewNumericDate(helperTime.AddMinutes(s.expMin, true)),
		// },

		UserID:   uuid.UUID(userID),
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Token: token,
	}, nil
}

func (s *UserService) SignIn(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.svc.GetUserByUsernamePassword(ctx, domain.Username(req.GetUsername()),
		domain.NewPassword(req.GetPassword()),
	)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if !user.PasswordIsCorrect(req.GetPassword()) {
		return nil, ErrInvalidUserPassword
	}

	token, err := jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		// RegisteredClaims: jwt2.RegisteredClaims{
		// 	// ExpiresAt: jwt2.NewNumericDate(helperTime.AddMinutes(s.expMin, true)),
		// },

		UserID:   uuid.UUID(user.ID),
		Username: req.Username,
		Email:    string(user.Email),
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token: token,
	}, nil
}
func (s *UserService) GetUserClaimWithToken(ctx context.Context, req *pb.TokenRequest) (*pb.UserClaimResponse, error) {
	user, err := s.svc.GetUserClaimWithToken(ctx, req.GetToken(), s.authSecret)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return &pb.UserClaimResponse{
		Username: string(user.Username),
		Email:    string(user.Email),
		Id:       user.UserID.String(),
	}, nil
}
