package access

import (
	"github.com/biryanim/auth/internal/service"
	descAccess "github.com/biryanim/auth/pkg/access_v1"
)

type Implementation struct {
	descAccess.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
