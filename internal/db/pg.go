package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func newPg(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)

	return pool, err
}
