package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pg *pgxpool.Pool
}

func New(ctx context.Context, pgDsn string) DB {
	pool, err := newPg(ctx, pgDsn)
	if err != nil {
		panic(err)
	}

	if pool.Ping(ctx) != nil {
		panic(err)
	}

	return DB{
		Pg: pool,
	}
}
