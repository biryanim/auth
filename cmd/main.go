package main

import (
	"context"
	"flag"
	userAPI "github.com/biryanim/auth/internal/api/user"
	"github.com/biryanim/auth/internal/config"
	userRepository "github.com/biryanim/auth/internal/repository/user"
	userService "github.com/biryanim/auth/internal/service/user"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer pool.Close()

	userRepo := userRepository.NewRepository(pool)
	userSrv := userService.NewService(userRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserAPIV1Server(s, userAPI.NewImplementation(userSrv))

	log.Printf("server listening on %s", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
