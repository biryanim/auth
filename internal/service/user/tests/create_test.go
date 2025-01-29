package tests

import (
	"context"
	"fmt"
	"github.com/biryanim/auth/internal/model"
	"github.com/biryanim/auth/internal/service/user"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"

	"github.com/biryanim/auth/internal/repository"
	repoMock "github.com/biryanim/auth/internal/repository/mocks"
	"github.com/gojuno/minimock/v3"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req *model.UserInfo
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()

		repoErr = fmt.Errorf("repo error")

		req = &model.UserInfo{
			Name:  name,
			Email: email,
			Role:  1,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(id, nil)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(0, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewMockService(userRepoMock)

			reps, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, reps)
		})
	}
}
