package service

import (
	"context"
	"errors"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
	userPort "github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/port"
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

	token, err := s.createTokens(userID)
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

	token, err := s.createTokens(user.ID)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (s *UserService) createTokens(userID domain.UserID) (token string, err error) {
	token = generateHMAC(userID.ConvStr(), s.authSecret)
	return
}

func generateHMAC(data, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func verifyHMAC(data, secretKey, expectedMAC string) bool {
	calculatedMAC := generateHMAC(data, secretKey)
	return hmac.Equal([]byte(calculatedMAC), []byte(expectedMAC))
}
