// Package database provides database connection and configuration management.
package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	ClientPgSQL *pgxpool.Pool
}

func NewDatabase(pgSQLPool *pgxpool.Pool) *Database {
	return &Database{
		ClientPgSQL: pgSQLPool,
	}
}
