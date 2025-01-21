package app

import (
	"context"
	"github.com/biryanim/auth/internal/api/user"
	"github.com/biryanim/auth/internal/closer"
	"github.com/biryanim/auth/internal/config"
	"github.com/biryanim/auth/internal/repository"
	userRepository "github.com/biryanim/auth/internal/repository/user"
	"github.com/biryanim/auth/internal/service"
	userService "github.com/biryanim/auth/internal/service/user"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool         *pgxpool.Pool
	userRepository repository.UserRepository

	userService service.UserService

	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	return s.pgPool
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.PgPool(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
