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

//func NewMockService(deps ...interface{}) service.UserService {
//	srv := serv{}
//
//	for _, dep := range deps {
//		switch s := dep.(type) {
//		case repository.UserRepository:
//			srv.userRepository = s
//		}
//	}
//
//	return &srv
//}
