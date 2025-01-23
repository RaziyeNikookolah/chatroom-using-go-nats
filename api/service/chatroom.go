package service

import (
	"context"
	"fmt"

	// jwt2 "github.com/golang-jwt/jwt/v5"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/domain"
	chatroomPort "github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/port"
)

type ChatroomService struct {
	svc chatroomPort.Service
}

func NewChatroomService(svc chatroomPort.Service) *ChatroomService {
	//, chatroomSvc chatroomPort.Service) *ChatroomService {
	return &ChatroomService{
		svc: svc,
	}
}

func (s *ChatroomService) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	err := s.svc.SendMessage(ctx, &domain.MessageToSend{
		Username: req.GetUsername(),
		UserID:   req.GetUserID(),
		Message:  req.GetMessage(),
	})

	if err != nil {
		return &pb.SendResponse{Response: false}, err
	}
	return &pb.SendResponse{
		Response: true,
	}, nil
}
func (s *ChatroomService) SubscribeUser(ctx context.Context, req *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {
	err := s.svc.SubscribeUser(ctx, req.GetUserID())

	if err != nil {
		return &pb.SubscribeResponse{Response: false}, err
	}
	return &pb.SubscribeResponse{
		Response: true,
	}, nil
}

func (s *ChatroomService) Show(ctx context.Context, req *pb.ShowRequest) (*pb.ShowResponse, error) {
	messages, err := s.svc.ShowMessages(ctx, req.GetUserID())
	if err != nil {
		return nil, err
	}

	if messages == nil {
		return nil, fmt.Errorf("%w: %w", ErrUserNotFound, ErrNoMessage)
	}

	return &pb.ShowResponse{
		Messages: messages.Messages,
	}, nil
}
func (s *ChatroomService) GetActiveUsers(ctx context.Context, req *pb.GetActiveUsersRequest) (*pb.GetActiveUsersResponse, error) {
	usernames, err := s.svc.GetActiveUsers(ctx)
	if err != nil {
		return nil, err
	}
	if usernames == nil {
		return nil, ErrNoOneIsActiveNow
	}

	return &pb.GetActiveUsersResponse{
		Usernames: usernames.Usernames}, nil
}
