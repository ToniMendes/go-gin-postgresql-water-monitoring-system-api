// Package database provides database connection and configuration management.
package database

import "github.com/jackc/pgx/v5"

type Database struct {
	ClientPgSQL *pgx.Conn
}

func NewDatabase(pgSQLConn *pgx.Conn) *Database {
	return &Database{
		ClientPgSQL: pgSQLConn,
	}
}
