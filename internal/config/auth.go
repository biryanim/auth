package config

import (
	"errors"
	"os"
	"time"
)

const (
	jwtRefreshTokenSecretKey = "JWT_REFRESH_TOKEN_SECRET"
	jwtAccessTokenSecretKey  = "JWT_ACCESS_TOKEN_SECRET"
	refreshTokenExpiration   = "REFRESH_TOKEN_EXPIRATION"
	accessTokenExpiration    = "ACCESS_TOKEN_EXPIRATION"
)

type AuthConfig interface {
	RefreshTokenSecret() []byte
	AccessTokenSecret() []byte
	RefreshTokenExpiration() time.Duration
	AccessTokenExpiration() time.Duration
}

type authConfig struct {
	refreshTokenSecret []byte
	accessTokenSecret  []byte
	refreshTokenExp    time.Duration
	accessTokenExp     time.Duration
}

func NewJWTConfig() (AuthConfig, error) {
	refreshTokenSecret := []byte(os.Getenv(jwtRefreshTokenSecretKey))
	if len(refreshTokenSecret) == 0 {
		return nil, errors.New("missing JWT refresh token secret")
	}

	accessTokenSecret := []byte(os.Getenv(jwtAccessTokenSecretKey))
	if len(accessTokenSecret) == 0 {
		return nil, errors.New("missing JWT access token secret")
	}

	refreshTokenExp, err := time.ParseDuration(os.Getenv(refreshTokenExpiration))
	if err != nil {
		return nil, err
	}

	accessTokenExp, err := time.ParseDuration(os.Getenv(accessTokenExpiration))
	if err != nil {
		return nil, err
	}

	return &authConfig{
		refreshTokenSecret: refreshTokenSecret,
		accessTokenSecret:  accessTokenSecret,
		refreshTokenExp:    refreshTokenExp,
		accessTokenExp:     accessTokenExp,
	}, nil
}

func (j *authConfig) RefreshTokenSecret() []byte {
	return j.refreshTokenSecret
}

func (j *authConfig) AccessTokenSecret() []byte {
	return j.accessTokenSecret
}

func (j *authConfig) RefreshTokenExpiration() time.Duration {
	return j.refreshTokenExp
}

func (j *authConfig) AccessTokenExpiration() time.Duration {
	return j.accessTokenExp
}
