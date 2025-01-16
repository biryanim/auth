package converter

import (
	"github.com/biryanim/auth/internal/repository/user/model"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromRepo(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromRepo(user model.Info) *desc.UserInfo {
	var role desc.Role

	switch user.Role {
	case "admin":
		role = desc.Role_admin
	case "user":
		role = desc.Role_user
	default:

		role = desc.Role_UNKNOWN_ROLE_TYPE
	}
	return &desc.UserInfo{
		Name:  user.Name,
		Email: user.Email,
		Role:  role,
	}
}
