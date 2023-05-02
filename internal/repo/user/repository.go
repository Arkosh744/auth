package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ Repository = (*repository)(nil)

const tableName = "users"

type Repository interface {
	Create(context.Context, *model.User) error
	ExistsName(ctx context.Context, username string) (exists bool, err error)
	ExistsEmail(ctx context.Context, email string) (exists bool, err error)
}

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

func (r *repository) Create(ctx context.Context, user *model.User) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("username", "email", "password", "role").
		Values(user.Username, user.Email, user.Password, user.Role)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Exec(ctx, query, v...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ExistsName(ctx context.Context, username string) (exists bool, err error) {
	err = r.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *repository) ExistsEmail(ctx context.Context, email string) (exists bool, err error) {
	err = r.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
