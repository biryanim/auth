package auth

import (
	"github.com/biryanim/auth/internal/model"
	"github.com/biryanim/auth/internal/utils"
	"golang.org/x/net/context"
)

func (s *serv) GetRefreshToken(ctx context.Context, oldToken string) (string, error) {
	claims, err := utils.VerifyToken(oldToken, s.authConfig.RefreshTokenSecret())
	if err != nil {
		return "", err
	}

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		Role:     claims.Role,
	}, s.authConfig.RefreshTokenSecret(), s.authConfig.RefreshTokenExpiration())
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *serv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, s.authConfig.RefreshTokenSecret())
	if err != nil {
		return "", err
	}

	accessToken, err := utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		Role:     claims.Role,
	}, s.authConfig.AccessTokenSecret(), s.authConfig.AccessTokenExpiration())
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
