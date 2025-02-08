package model

import (
	"database/sql"
	"time"
)

type AccessInfo struct {
	Id              int64        `db:"id"`
	EndpointAddress string       `db:"endpoint_address"`
	Role            int32        `db:"role"`
	CreatedAt       time.Time    `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
}
