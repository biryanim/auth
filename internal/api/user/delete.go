package user

import (
	"context"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
