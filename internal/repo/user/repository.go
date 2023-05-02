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
	Get(ctx context.Context, username string) (user *model.User, err error)
	ExistsName(ctx context.Context, username string) (exists bool, err error)
	ExistsEmail(ctx context.Context, email string) (exists bool, err error)
	Update(ctx context.Context, username string, user *model.User) error
	Delete(ctx context.Context, username string) error
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

func (r *repository) Get(ctx context.Context, username string) (user *model.User, err error) {
	builder := sq.Select("username", "email", "password", "role", "created_at", "updated_at").
		From(tableName).
		Where("username = ?", username).
		Where("deleted_at IS NULL").
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	user = &model.User{}
	err = r.pool.QueryRow(ctx, query, v...).
		Scan(&user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Update(ctx context.Context, username string, user *model.User) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("updated_at", user.UpdatedAt).
		Where("username = ?", username).
		Where("deleted_at IS NULL")

	if user.Username != "" {
		builder = builder.Set("username", user.Username)
	}

	if user.Email != "" {
		builder = builder.Set("email", user.Email)
	}

	if user.Password != "" {
		builder = builder.Set("password", user.Password)
	}

	if user.Role != "" {
		builder = builder.Set("role", user.Role)
	}

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

func (r *repository) Delete(ctx context.Context, username string) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("deleted_at", "NOW()").
		Where("username = ?", username)

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
