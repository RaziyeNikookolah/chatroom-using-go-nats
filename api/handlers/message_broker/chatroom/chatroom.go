package chatroom

import (
	"context"
	"encoding/json"
	"log"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/service"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/chatroom/domain"
)

type ChatroomHandler struct {
	chatroomService *service.ChatroomService
}

func NewChatroomHandler(chatroomService *service.ChatroomService) *ChatroomHandler {
	return &ChatroomHandler{chatroomService: chatroomService}
}

func (h *ChatroomHandler) Send(message []byte) {
	var req domain.MessageToSend
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Printf("Failed to deserialize message: %v", err)
	}

	h.chatroomService.Send(context.Background(), &pb.SendRequest{
		UserID:   req.UserID,
		Username: req.Username,
		Message:  req.Message,
	})
}
func (h *ChatroomHandler) GetActiveUsers(data []byte) {

	h.chatroomService.GetActiveUsers(context.Background(), &pb.GetActiveUsersRequest{})
}

func (h *ChatroomHandler) Show(data []byte) {

	h.chatroomService.Show(context.Background(), &pb.ShowRequest{
		UserID: string(data),
	})
}
func (h *ChatroomHandler) Subscribe(data []byte) {
	h.chatroomService.SubscribeUser(context.Background(), &pb.SubscribeRequest{
		UserID: string(data),
	})
}
