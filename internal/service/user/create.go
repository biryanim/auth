package user

import (
	"context"
	"github.com/biryanim/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, userInfo *model.UserInfo) (int64, error) {
	id, err := s.userRepository.Create(ctx, userInfo)
	if err != nil {
		return 0, err
	}
	return id, nil
}
