package service

import (
	"context"
	"github.com/biryanim/auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, userInfo *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, id int64, updateInfo *model.UpdateUserInfo) error
	Delete(ctx context.Context, id int64) error
}
