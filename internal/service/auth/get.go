package auth

import "golang.org/x/net/context"

func (s *serv) GetRefreshToken(ctx context.Context, oldToken string) (string, error) {
	return oldToken, nil
}

func (s *serv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	return refreshToken, nil
}
