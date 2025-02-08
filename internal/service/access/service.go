package access

import (
	"github.com/biryanim/auth/internal/config"
	"github.com/biryanim/auth/internal/repository"
	"github.com/biryanim/auth/internal/service"
)

type serv struct {
	authConfig       config.AuthConfig
	accessRepository repository.AccessRepository
}

func NewService(authConfig config.AuthConfig, accessRepository repository.AccessRepository) service.AccessService {
	return &serv{
		authConfig:       authConfig,
		accessRepository: accessRepository,
	}
}
