package user

import (
	"context"
	"github.com/biryanim/auth/internal/converter"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"log"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userObj, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d,	name: %s,	email: %s,	role: %v,	created_at: %v,	updated_at: %v,", userObj.ID, userObj.Info.Name, userObj.Info.Email, userObj.Info.Role, userObj.CreatedAt, userObj.UpdatedAt)

	return &desc.GetResponse{
		User: converter.ToUserFromService(userObj),
	}, nil
}
