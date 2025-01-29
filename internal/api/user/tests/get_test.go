package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/biryanim/auth/internal/api/user"
	"github.com/biryanim/auth/internal/model"
	serviceMock "github.com/biryanim/auth/internal/service/mocks"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/biryanim/auth/internal/service"
	"github.com/gojuno/minimock/v3"
	"testing"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		serviceRes = &model.User{
			ID: id,
			Info: model.UserInfo{
				Name:  name,
				Email: email,
				Role:  1,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Valid: true,
				Time:  updatedAt,
			},
		}

		res = &desc.GetResponse{
			User: &desc.User{
				Id: id,
				Info: &desc.UserInfo{
					Name:  name,
					Email: email,
					Role:  1,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(serviceRes, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewImplementation(userServiceMock)
			resp, err := api.Get(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resp)
		})
	}
}
