package user

import (
	"github.com/biryanim/auth/internal/repository"
	"github.com/biryanim/auth/internal/service"
	"github.com/biryanim/platform_common/pkg/db"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(userRepository repository.UserRepository, tx db.TxManager) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      tx,
	}
}
