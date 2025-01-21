package user

import (
	"github.com/biryanim/auth/internal/repository"
	"github.com/biryanim/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) service.UserService {
	return &serv{
		userRepository: userRepository,
	}
}
