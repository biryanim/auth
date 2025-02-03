package auth

import (
	"github.com/biryanim/auth/internal/model"
	"golang.org/x/net/context"
)

func (s *serv) Login(ctx context.Context, login *model.LoginDTO) (string, error) {
	return "", nil
}
