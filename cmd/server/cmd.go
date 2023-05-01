package main

import (
	"context"
	"log"
	"net"

	userV1 "github.com/Arkosh744/auth-grpc/internal/api/user_v1"
	userRepo "github.com/Arkosh744/auth-grpc/internal/repo/user"
	userService "github.com/Arkosh744/auth-grpc/internal/service/user"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = ":50051"

func main() {
	ctx := context.Background()
	list, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pgConfig, err := pgxpool.ParseConfig("host=localhost port=54325 dbname=grpc user=grpc-user password=grpc-password sslmode=disable")
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	dbc, err := pgxpool.ConnectConfig(ctx, pgConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbc.Close()

	repo := userRepo.NewRepository(dbc)
	service := userService.NewService(repo)
	desc.RegisterUserServer(s, userV1.NewImplementation(service))

	err = s.Serve(list)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
