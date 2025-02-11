package user

import (
	"context"
	"github.com/biryanim/auth/internal/model"
	"github.com/biryanim/platform_common/pkg/filter"
	"github.com/biryanim/platform_common/pkg/sys/validate"
	"time"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	err := validate.Validate(ctx, validateID(id))
	if err != nil {
		return nil, err
	}

	idFilter := filter.New(filter.Condition{
		Key:   "id",
		Value: id,
	})

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := s.userRepository.Get(ctx, idFilter)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func validateID(id int64) validate.Condition {
	return func(ctx context.Context) error {
		if id <= 0 {
			return validate.NewValidateErrors("id must be greater than 0")
		}
		return nil
	}
}
