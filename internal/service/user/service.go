package user

import (
	"github.com/biryanim/auth/internal/client/db"
	"github.com/biryanim/auth/internal/repository"
	"github.com/biryanim/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(userRepository repository.UserRepository, txManager db.TxManager) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
