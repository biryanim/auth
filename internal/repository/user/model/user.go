package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64        `db:"id"`
	Info      Info         `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	Password  string       `db:"password"`
}

type Info struct {
	Name     string `db:"name"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Role     int32  `db:"role"`
}
