package repository

import (
	"context"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
)

type UsersRepository interface {
	Create(ctx context.Context, userInfo *desc.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*desc.User, error)
	Update(ctx context.Context, id int64, updateInfo *desc.UpdateUserInfo) error
	Delete(ctx context.Context, id int64) error
}
