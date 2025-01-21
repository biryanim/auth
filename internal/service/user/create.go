package user

import (
	"context"
	"github.com/biryanim/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, userInfo *model.UserInfo) (int64, error) {
	var id int64
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Create(ctx, userInfo)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}
	return id, nil
}
