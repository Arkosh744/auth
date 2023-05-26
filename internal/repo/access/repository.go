package user

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/client/pg"
	"github.com/Arkosh744/auth-service-api/internal/model"
	sq "github.com/Masterminds/squirrel"
)

type Repository interface {
	GetList(ctx context.Context) ([]*model.AccessInfo, error)
}

type repository struct {
	client pg.Client
}

func NewRepository(client pg.Client) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) GetList(ctx context.Context) ([]*model.AccessInfo, error) {
	builder := sq.Select("id", "endpoint_address", "role").
		PlaceholderFormat(sq.Dollar).
		From("accesses")

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "accesses.GetList",
		QueryRaw: query,
	}

	var accessInfo []*model.AccessInfo
	err = r.client.PG().ScanAllContext(ctx, &accessInfo, q, v...)
	if err != nil {
		return nil, err
	}

	return accessInfo, nil
}
