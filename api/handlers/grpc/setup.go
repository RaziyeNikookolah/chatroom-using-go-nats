package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/app"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/config"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/handlers/grpc/user"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func Run(cfg config.Config, app app.App) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.GrpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	healthServer := &user.HealthServer{}
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	reflection.Register(grpcServer)

	log.Println("User | GRPC server started..")

	// vHandler := service.NewVehicleService(app.VehicleService())
	// d := vehicle.NewGRPCVehicleHandler(*vHandler)

	userHandler := service.NewUserService(app.UserService(context.Background()), cfg.Server.Secret)
	userGrpcHandler := user.NewGRPCUserHandler(*userHandler)
	pb.RegisterUserServiceServer(grpcServer, userGrpcHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
