package access

import (
	"github.com/biryanim/auth/internal/repository"
	"github.com/biryanim/platform_common/pkg/db"
)

const (
	accessesTableName = "accesses"

	idColumn              = "id"
	endpointAddressColumn = "endpoint_address"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
	roleColumn            = "role"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AccessRepository {
	return &repo{
		db: db,
	}
}
