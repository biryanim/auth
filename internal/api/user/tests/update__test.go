package tests

import (
	"context"
	"fmt"
	"github.com/biryanim/auth/internal/api/user"
	"github.com/biryanim/auth/internal/model"
	"github.com/biryanim/auth/internal/service"
	serviceMock "github.com/biryanim/auth/internal/service/mocks"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()

		serviceErr = fmt.Errorf("service error")

		serviceReq = &model.UpdateUserInfo{
			Name:  &name,
			Email: &email,
		}

		req = &desc.UpdateRequest{
			Id: id,
			Info: &desc.UpdateUserInfo{
				Name:  &wrapperspb.StringValue{Value: name},
				Email: &wrapperspb.StringValue{Value: email},
			},
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, id, serviceReq).Return(nil)
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
				mock.UpdateMock.Expect(ctx, id, serviceReq).Return(serviceErr)
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

			resp, err := api.Update(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resp)
		})
	}
}
