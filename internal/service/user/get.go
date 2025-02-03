package user

import (
	"context"
	"github.com/biryanim/auth/internal/model"
	"github.com/biryanim/platform_common/pkg/filter"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	idFilter := filter.New(filter.Condition{
		Key:   "id",
		Value: id,
	})
	user, err := s.userRepository.Get(ctx, idFilter)
	if err != nil {
		return nil, err
	}

	return user, nil
}
