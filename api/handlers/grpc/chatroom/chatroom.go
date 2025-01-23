package chatroom

import (
	"context"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCChatroomHandler struct {
	pb.UnimplementedChatroomServiceServer
	chatroomService *service.ChatroomService
}

func NewGRPCChatroomHandler(chatroomService service.ChatroomService) *GRPCChatroomHandler {
	return &GRPCChatroomHandler{chatroomService: &chatroomService}
}

func (g *GRPCChatroomHandler) Send(ctx context.Context, sReq *pb.SendRequest) (*pb.SendResponse, error) {
	response, err := g.chatroomService.Send(ctx, sReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return response, nil
}
func (g *GRPCChatroomHandler) Subscribe(ctx context.Context, sReq *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {
	response, err := g.chatroomService.SubscribeUser(ctx, sReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return response, nil
}
func (g *GRPCChatroomHandler) Show(ctx context.Context, regUsr *pb.ShowRequest) (*pb.ShowResponse, error) {
	response, err := g.chatroomService.Show(ctx, regUsr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return response, nil
}
func (g *GRPCChatroomHandler) GetActiveUsers(ctx context.Context, t *pb.GetActiveUsersRequest) (*pb.GetActiveUsersResponse, error) {
	response, err := g.chatroomService.GetActiveUsers(ctx, t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return response, nil
}
