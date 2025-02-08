package auth

import (
	"github.com/biryanim/auth/internal/config"
	"github.com/biryanim/auth/internal/repository"
	"github.com/biryanim/auth/internal/service"
)

type serv struct {
	userRepo   repository.UserRepository
	authConfig config.AuthConfig
}

func NewService(userRepo repository.UserRepository, cfg config.AuthConfig) service.AuthService {
	return &serv{
		userRepo:   userRepo,
		authConfig: cfg,
	}
}
