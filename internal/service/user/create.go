package user

import (
	"context"
	"github.com/biryanim/auth/internal/model"
	"github.com/biryanim/auth/internal/utils"
)

func (s *serv) Create(ctx context.Context, userInfo *model.UserInfo) (int64, error) {
	var id int64

	hashedPassword, err := utils.HashPassword(userInfo.Password)
	if err != nil {
		return 0, err
	}
	userInfo.Password = hashedPassword

	id, err = s.userRepository.Create(ctx, userInfo)

	if err != nil {
		return 0, err
	}
	return id, nil
}
