package access

import (
	"context"
	descAccess "github.com/biryanim/auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Check(ctx context.Context, req *descAccess.CheckRequest) (*emptypb.Empty, error) {

	_, err := i.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
