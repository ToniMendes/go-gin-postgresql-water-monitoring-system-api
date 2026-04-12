// Package postgresql provides PostgreSQL database connectivity and operations.
package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewPgSQLPool(ctx context.Context, connectionString string) (*pgx.Conn, error) {

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
