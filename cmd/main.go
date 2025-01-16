package main

import (
	"context"
	"flag"
	"github.com/biryanim/auth/internal/config"
	"github.com/biryanim/auth/internal/repository"
	"github.com/biryanim/auth/internal/repository/user"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

var configPath string

type server struct {
	desc.UnimplementedUserAPIV1Server
	usersRepository repository.UsersRepository
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req.GetPassword() != req.GetPassword() {
		return nil, status.Error(codes.InvalidArgument, "password does not match")
	}

	id, err := s.usersRepository.Create(ctx, req.GetInfo())
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	err := s.usersRepository.Update(ctx, req.GetId(), req.GetInfo())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	err := s.usersRepository.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userObj, err := s.usersRepository.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d,	name: %s,	email: %s,	role: %v,	created_at: %v,	updated_at: %v,", userObj.Id, userObj.Info.Name, userObj.Info.Email, userObj.Info.Role, userObj.CreatedAt, userObj.UpdatedAt)

	return &desc.GetResponse{
		User: userObj,
	}, nil
}

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

	userRepo := user.NewRepository(pool)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserAPIV1Server(s, &server{usersRepository: userRepo})

	log.Printf("server listening on %s", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
