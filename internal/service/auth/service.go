package auth

import (
	"github.com/biryanim/auth/internal/repository"
	"github.com/biryanim/auth/internal/service"
)

type serv struct {
	userRepo repository.UserRepository
}

func NewService(userRepo repository.UserRepository) service.AuthService {
	return &serv{
		userRepo: userRepo,
	}
}
