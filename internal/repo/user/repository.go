package user

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/client/pg"
	"github.com/Arkosh744/auth-service-api/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

var _ Repository = (*repository)(nil)

const tableName = "users"

type Repository interface {
	Create(context.Context, *model.User) error
	Get(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
	ExistsNameEmail(ctx context.Context, user *model.UserIdentifier) (model.ExistsStatus, error)
	Update(ctx context.Context, username string, user *model.UpdateUser) error
	Delete(ctx context.Context, username string) error
}

type repository struct {
	client pg.Client
}

func NewRepository(client pg.Client) *repository {
	return &repository{
		client: client,
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

	q := pg.Query{
		Name:     "user.Create",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) ExistsNameEmail(ctx context.Context, user *model.UserIdentifier) (model.ExistsStatus, error) {
	builder := sq.Select().PlaceholderFormat(sq.Dollar).
		Column(sq.Expr(`
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

	q := pg.Query{
		Name:     "user.ExistsNameEmail",
		QueryRaw: query,
	}

	var existsCode model.ExistsStatus
	if err = r.client.PG().QueryRowContext(ctx, q, v...).Scan(&existsCode); err != nil {
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

	q := pg.Query{
		Name:     "user.Get",
		QueryRaw: query,
	}

	var user model.User
	if err = r.client.PG().GetContext(ctx, &user, q, v...); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) List(ctx context.Context) ([]*model.User, error) {
	builder := sq.Select("username", "email", "password", "role", "created_at", "updated_at").
		From(tableName).
		Where("deleted_at IS NULL").
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "user.List",
		QueryRaw: query,
	}

	var users = make([]*model.User, 0)
	if err = r.client.PG().ScanAllContext(ctx, &users, q, v...); err != nil {
		return nil, err
	}

	return users, nil
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

	q := pg.Query{
		Name:     "user.Update",
		QueryRaw: query,
	}

	row, err := r.client.PG().ExecContext(ctx, q, v...)
	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
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

	q := pg.Query{
		Name:     "user.Delete",
		QueryRaw: query,
	}

	row, err := r.client.PG().ExecContext(ctx, q, v...)
	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
