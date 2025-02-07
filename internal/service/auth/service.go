package auth

import (
	"github.com/biryanim/auth/internal/config"
	"github.com/biryanim/auth/internal/repository"
	"github.com/biryanim/auth/internal/service"
)

type serv struct {
	userRepo repository.UserRepository
	jwtCfg   config.JWTConfig
}

func NewService(userRepo repository.UserRepository, cfg config.JWTConfig) service.AuthService {
	return &serv{
		userRepo: userRepo,
		jwtCfg:   cfg,
	}
}
