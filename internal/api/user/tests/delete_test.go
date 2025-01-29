package tests

import (
	"context"
	"fmt"
	"github.com/biryanim/auth/internal/api/user"
	serviceMock "github.com/biryanim/auth/internal/service/mocks"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/biryanim/auth/internal/service"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"github.com/gojuno/minimock/v3"
	"testing"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceError = fmt.Errorf("service error")

		req = &desc.DeleteRequest{
			Id: id,
		}
	)

	tests := []struct {
		name                string
		args                args
		want                *emptypb.Empty
		err                 error
		userServiceMockFunc func(mc *minimock.Controller) service.UserService
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			userServiceMockFunc: func(mc *minimock.Controller) service.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.DeleteMock.Expect(ctx, req.GetId()).Return(nil)
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
			err:  serviceError,
			userServiceMockFunc: func(mc *minimock.Controller) service.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.DeleteMock.Expect(ctx, req.GetId()).Return(serviceError)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMockFunc(mc)
			api := user.NewImplementation(userServiceMock)

			resp, err := api.Delete(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resp)
		})
	}
}
