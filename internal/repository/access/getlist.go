package access

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/auth/internal/model"
	"github.com/biryanim/platform_common/pkg/db"
)

func (r *repo) GetList(ctx context.Context) ([]*model.AccessInfo, error) {
	builder := sq.Select(idColumn, endpointAddressColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(accessesTableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "access_repository_list",
		QueryRaw: query,
	}

	var result []*model.AccessInfo
	err = r.db.DB().ScanAllContext(ctx, result, q, args)
	if err != nil {
		return nil, err
	}

	return result, nil
}
