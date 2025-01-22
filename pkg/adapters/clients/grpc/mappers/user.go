package mappers

import (
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"
)

//	func RegisterResponseDomainToRegisterResponseProto(user *domain.User) *pb.RegisterResponse {
//		return &pb.RegisterResponse{
//			Token: ,
//		}
//	}
func RegisterResponseProtoToRegisterResponseDomain(m *pb.RegisterResponse) (*domain.RegisterResponse, error) {
	return &domain.RegisterResponse{Token: m.Token}, nil
}
func LoginResponseProtoToLoginResponseDomain(m *pb.LoginResponse) (*domain.LoginResponse, error) {
	return &domain.LoginResponse{Token: m.Token}, nil
}
