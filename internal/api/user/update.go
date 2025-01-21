package user

import (
	"context"
	"github.com/biryanim/auth/internal/converter"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	err := i.userService.Update(ctx, req.GetId(), converter.ToUpdatedUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
