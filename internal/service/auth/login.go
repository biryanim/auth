package auth

import (
	"errors"
	"github.com/biryanim/auth/internal/model"
	"github.com/biryanim/auth/internal/utils"
	"github.com/biryanim/platform_common/pkg/filter"
	"golang.org/x/net/context"
)

func (s *serv) Login(ctx context.Context, login *model.LoginDTO) (string, error) {
	filt := filter.New(filter.Condition{
		Key:   "username",
		Value: login.Username,
	})

	user, err := s.userRepo.Get(ctx, filt)
	if err != nil {
		return "", err
	}

	if !utils.VerifyPassword(user.Password, login.Password) {
		return "", errors.New("invalid password")
	}

	refreshToken, err := utils.GenerateToken(user.Info, s.authConfig.RefreshTokenSecret(), s.authConfig.RefreshTokenExpiration())
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
