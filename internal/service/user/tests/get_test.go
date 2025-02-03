package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/biryanim/auth/internal/model"
	"github.com/biryanim/auth/internal/repository"
	repoMock "github.com/biryanim/auth/internal/repository/mocks"
	"github.com/biryanim/auth/internal/service/user"
	"github.com/biryanim/platform_common/pkg/db"
	txMock "github.com/biryanim/platform_common/pkg/db/mocks"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"

	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txManagerMock func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		repoErr = fmt.Errorf("repo error")

		res = &model.User{
			ID: id,
			Info: model.UserInfo{
				Name:  name,
				Email: email,
				Role:  1,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{Time: updatedAt, Valid: true},
		}
	)
	defer t.Cleanup(mc.Finish)

	testCases := []struct {
		name                string
		args                args
		want                *model.User
		err                 error
		userRepoositoryMock userRepositoryMockFunc
		txManagerMock       txManagerMock
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: res,
			err:  nil,
			userRepoositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(res, nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMock.NewTxManagerMock(mc)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: nil,
			err:  repoErr,
			userRepoositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, repoErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMock.NewTxManagerMock(mc)
				return mock
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepoositoryMock(mc)
			txManagerMock := tt.txManagerMock(mc)
			service := user.NewService(userRepoMock, txManagerMock)

			resp, err := service.Get(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resp)
		})
	}
}
