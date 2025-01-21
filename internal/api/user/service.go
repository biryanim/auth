package user

import (
	"github.com/biryanim/auth/internal/service"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
)

type Implementation struct {
	desc.UnimplementedUserAPIV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
