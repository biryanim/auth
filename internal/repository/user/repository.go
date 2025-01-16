package user

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/auth/internal/repository"
	"github.com/biryanim/auth/internal/repository/user/converter"
	"github.com/biryanim/auth/internal/repository/user/model"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UsersRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, userInfo *desc.UserInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, roleColumn).
		Values(userInfo.Name, userInfo.Email, userInfo.Role.String()).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*desc.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user model.User
	err = r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Info.Name, &user.Info.Email, &user.Info.Role, &user.CreatedAt, &user.UpdatedAt)
	fmt.Println("HEEEEEEREEEE")
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Update(ctx context.Context, id int64, updateInfo *desc.UpdateUserInfo) error {
	var name, email *string
	if updateInfo.Name != nil {
		nameValue := updateInfo.Name.GetValue()
		name = &nameValue
	}
	if updateInfo.Email != nil {
		emailValue := updateInfo.Email.GetValue()
		email = &emailValue
	}
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, name).
		Set(emailColumn, email).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}
