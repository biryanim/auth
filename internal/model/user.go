package model

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	ID        int64        `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserInfo struct {
	Name     string `db:"name"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Role     int32  `db:"role"`
	Password string `db:"password"`
}

type UpdateUserInfo struct {
	Name  *string `db:"name"`
	Email *string `db:"email"`
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     int32  `json:"role"`
}
