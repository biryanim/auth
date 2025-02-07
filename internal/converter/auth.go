package converter

import (
	"github.com/biryanim/auth/internal/model"
	descAuth "github.com/biryanim/auth/pkg/auth_v1"
)

func ToLoginDTOFromDesc(info *descAuth.UserInfo) *model.LoginDTO {
	return &model.LoginDTO{
		Username: info.Username,
		Password: info.Password,
	}
}
