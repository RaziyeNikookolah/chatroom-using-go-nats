package mappers

// import (
// 	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
// 	"github.com/RaziyeNikookolah/chatroom-using-go-nats/internal/user/domain"

// 	"github.com/google/uuid"
// )

// func UserClaimsToDomain(p *pb.GetUserByTokenResponse) (*domain.User, error) {
// 	u, err := uuid.Parse(p.UserId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &domain.User{
// 		ID:      u,
// 		IsAdmin: p.IsAdmin,
// 	}, nil
// }
