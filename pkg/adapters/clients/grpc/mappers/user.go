package mappers

import (
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
)

func RegisterResponseProtoToRegisterResponseDomain(m *pb.RegisterResponse) (*domain.RegisterResponse, error) {
	return &domain.RegisterResponse{Token: m.Token}, nil
}
func LoginResponseProtoToLoginResponseDomain(m *pb.LoginResponse) (*domain.LoginResponse, error) {
	return &domain.LoginResponse{Token: m.Token}, nil
}
func UserClaimResponseProtoToUserResponseDomain(m *pb.UserClaimResponse) (*domain.UserClaim, error) {
	return &domain.UserClaim{
		Username: m.Username,
		Email:    m.Email,
		ID:       m.Id,
	}, nil
}
