// Package postgresql provides PostgreSQL database pool.
package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgSQLPool(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {

	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
