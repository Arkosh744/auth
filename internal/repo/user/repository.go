package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ Repository = (*repository)(nil)

const tableName = "users"

type Repository interface {
	Create(context.Context, *model.User) error
	Get(ctx context.Context, username string) (*model.User, error)
	ExistsNameEmail(ctx context.Context, user *model.UserIdentifier) (model.ExistsStatus, error)
	Update(ctx context.Context, username string, user *model.UpdateUser) error
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

func (r *repository) ExistsNameEmail(ctx context.Context, user *model.UserIdentifier) (model.ExistsStatus, error) {
	builder := sq.Select().PlaceholderFormat(sq.Dollar).Column(sq.Expr(`
			CASE 
				WHEN EXISTS(SELECT 1 FROM users WHERE (username = $1) AND (email = $2)) THEN 3 
				WHEN EXISTS(SELECT 1 FROM users WHERE username = $1) THEN 1 
				WHEN EXISTS(SELECT 1 FROM users WHERE email = $2) THEN 2 
				ELSE 0 
			END 
			as exists_code`,
		user.Username.String, user.Email.String))

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var existsCode model.ExistsStatus
	err = r.pool.QueryRow(ctx, query, v...).Scan(&existsCode)
	if err != nil {
		return 0, err
	}

	return existsCode, nil
}

func (r *repository) Get(ctx context.Context, username string) (*model.User, error) {
	builder := sq.Select("username", "email", "password", "role", "created_at", "updated_at").
		From(tableName).
		Where(sq.Eq{"username": username}).
		Where("deleted_at IS NULL").
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	user := &model.User{}
	roleStr := ""
	if err = r.pool.QueryRow(ctx, query, v...).
		Scan(&user.Username, &user.Email, &user.Password, &roleStr, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	user.Role = model.StringToRole(roleStr)

	return user, nil
}

func (r *repository) Update(ctx context.Context, username string, user *model.UpdateUser) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("updated_at", sq.Expr("NOW()")).
		Where(sq.Eq{"username": username}).
		Where("deleted_at IS NULL")

	if user.Username.Valid {
		builder = builder.Set("username", user.Username.String)
	}

	if user.Email.Valid {
		builder = builder.Set("email", user.Email.String)
	}

	if user.Password.Valid {
		builder = builder.Set("password", user.Password.String)
	}

	if user.Role.Valid {
		builder = builder.Set("role", user.Role.String)
	}

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := r.pool.Exec(ctx, query, v...)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, username string) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("deleted_at", sq.Expr("NOW()")).
		Where(sq.Eq{"username": username})

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	row, err := r.pool.Exec(ctx, query, v...)
	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
