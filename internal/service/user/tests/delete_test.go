package tests

import (
	"context"
	"fmt"
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

func TestDelete(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txManagerMock func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		repoErr = fmt.Errorf("repo error")
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		err                error
		userRepositoryMock userRepositoryMockFunc
		txManagerMock      txManagerMock
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			err: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
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
				req: id,
			},
			err: repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(repoErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMock.NewTxManagerMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(mc)
			service := user.NewService(userRepoMock, txManagerMock)

			err := service.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
