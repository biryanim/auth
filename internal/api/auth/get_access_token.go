package auth

import (
	"context"
	descAuth "github.com/biryanim/auth/pkg/auth_v1"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *descAuth.GetAccessTokenRequest) (*descAuth.GetAccessTokenResponse, error) {
	token, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &descAuth.GetAccessTokenResponse{
		AccessToken: token,
	}, nil
}
