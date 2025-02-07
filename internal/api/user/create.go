package user

import (
	"context"
	"github.com/biryanim/auth/internal/converter"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req.GetPassword() != req.GetPasswordConfirm() {
		return nil, status.Error(codes.InvalidArgument, "password does not match")
	}

	id, err := i.userService.Create(ctx, converter.ToUserCreateFromDesc(req.GetInfo(), req.GetPassword()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
