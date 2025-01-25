package grpc

import (
	"log"

	"google.golang.org/grpc"
)

func GetGrpcConnection() *grpc.ClientConn {
	// Create a gRPC connection to the user service
	grpcConn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client cannot connect to server: %v", err)
	}
	return grpcConn
}
