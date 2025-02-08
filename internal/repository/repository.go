package repository

import (
	"context"
	"github.com/biryanim/auth/internal/model"
	"github.com/biryanim/platform_common/pkg/filter"
)

type UserRepository interface {
	Create(ctx context.Context, userInfo *model.UserCreate) (int64, error)
	Get(ctx context.Context, filter *filter.Filter) (*model.User, error)
	Update(ctx context.Context, id int64, updateInfo *model.UpdateUserInfo) error
	Delete(ctx context.Context, id int64) error
}

type AccessRepository interface {
	GetList(ctx context.Context) ([]*model.AccessInfo, error)
}
