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
	}
}

func ToUserInfoFromRepo(user modelRepo.Info) model.UserInfo {
	//var role desc.Role
	//
	//switch user.Role {
	//case "admin":
	//	role = desc.Role_admin
	//case "user":
	//	role = desc.Role_user
	//default:
	//
	//	role = desc.Role_UNKNOWN_ROLE_TYPE
	//}
	//return &desc.UserInfo{
	//	Name:  user.Name,
	//	Email: user.Email,
	//	Role:  desc.Role(user.Role),
	//}

	return model.UserInfo{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Password: user.Password,
	}
}
