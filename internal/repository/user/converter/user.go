package converter

import (
	"github.com/biryanim/auth/internal/model"
	modelRepo "github.com/biryanim/auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Password:  user.Password,
	}
}

func ToUserInfoFromRepo(user modelRepo.Info) model.UserInfo {
	return model.UserInfo{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
}
