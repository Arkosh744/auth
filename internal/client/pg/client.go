package pg

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ Client = (*client)(nil)

type Client interface {
	Close() error
	PG() PG
}

type client struct {
	pg PG
}

func NewClient(ctx context.Context, pgCfg *pgxpool.Config) (Client, error) {
	dbc, err := pgxpool.ConnectConfig(ctx, pgCfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err.Error())
	}

	return &client{pg: &pg{pgxPool: dbc}}, nil
}

func (c *client) PG() PG {
	return c.pg
}

func (c *client) Close() error {
	if c.pg != nil {
		return c.pg.Close()
	}

	return nil
}
