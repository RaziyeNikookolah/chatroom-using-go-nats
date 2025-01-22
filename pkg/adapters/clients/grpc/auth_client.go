package grpc

import (
	"context"
	"fmt"

	"github.com/onlineTraveling/auth/protobufs"
	"github.com/onlineTraveling/bank/internal/user/domain"
	"github.com/onlineTraveling/bank/pkg/adapters/clients/grpc/mappers"
	"github.com/onlineTraveling/bank/pkg/ports"
	"google.golang.org/grpc"
)

type GRPCAuthClient struct {
	ServiceRegistry ports.IServiceRegistry
	AuthServiceName string
}

func NewGRPCAuthClient(serviceRegistry ports.IServiceRegistry, authServiceName string) *GRPCAuthClient {
	return &GRPCAuthClient{ServiceRegistry: serviceRegistry, AuthServiceName: authServiceName}
}

func (g *GRPCAuthClient) GetUserByToken(token string) (*domain.User, error) {
	// port, ip, err := g.ServiceRegistry.DiscoverService(g.AuthServiceName)
	port := 50052
	ip := "localhost"
	// if err != nil {
	// 	return nil, err
	// }

	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", ip, port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create a new AuthService client
	client := protobufs.NewAuthServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Prepare the request
	request := &protobufs.GetUserByTokenRequest{
		Token: token,
	}

	// Call the GetUserByToken method
	response, err := client.GetUserByToken(ctx, request)
	if err != nil {
		// print("token", request.Token)

		return nil, domain.ErrInvalidToken
	}
	domainUser, err := mappers.UserClaimsToDomain(response)
	if err != nil {
		return nil, err
	}
	print(domainUser)
	return domainUser, nil
}
