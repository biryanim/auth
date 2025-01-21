package user

import (
	"context"
	"github.com/biryanim/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, updateInfo *model.UpdateUserInfo) error {
	err := s.userRepository.Update(ctx, id, updateInfo)
	if err != nil {
		return err
	}

	return nil
}
