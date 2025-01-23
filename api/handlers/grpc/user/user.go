package user

import (
	"context"
	"errors"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/service"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type GRPCUserHandler struct {
	pb.UnimplementedUserServiceServer
	userService *service.UserService
}

func NewGRPCUserHandler(userService service.UserService) *GRPCUserHandler {
	return &GRPCUserHandler{userService: &userService}
}

func (g *GRPCUserHandler) Register(ctx context.Context, regUsr *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	response, err := g.userService.SignUp(ctx, regUsr)
	if err != nil {
		if errors.Is(err, port.ErrUserAlreadyExist) {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return response, nil
}
func (g *GRPCUserHandler) Login(ctx context.Context, regUsr *pb.LoginRequest) (*pb.LoginResponse, error) {
	response, err := g.userService.SignIn(ctx, regUsr)
	if err != nil {
		if errors.Is(err, port.ErrInvalidCredential) {
			return nil, status.Errorf(codes.Unknown, "invalid credential")
		}
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return response, nil
}
func (g *GRPCUserHandler) GetUserClaimWithToken(ctx context.Context, t *pb.TokenRequest) (*pb.UserClaimResponse, error) {
	response, err := g.userService.GetUserClaimWithToken(ctx, t)
	if err != nil {
		if errors.Is(err, port.ErrInvalidToken) {
			return nil, status.Errorf(codes.Unknown, "invalid token")
		}
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return response, nil
}

type HealthServer struct {
	grpc_health_v1.HealthServer
}

// Check implements Health.Check
func (s *HealthServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
