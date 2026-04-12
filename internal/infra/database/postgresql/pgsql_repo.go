package postgresql

import (
	"go-gin-postgresql-water-monitoring-system-api/internal/domain"
	"go-gin-postgresql-water-monitoring-system-api/internal/domain/entities"

	"github.com/jackc/pgx/v5"
)

type PgSQLRepo struct {
	client *pgx.Conn
}

func NewPgSQLRepo(client *pgx.Conn) *PgSQLRepo {
	return &PgSQLRepo{
		client: client,
	}
}

func (r *PgSQLRepo) Save(owner *entities.Owner, address *entities.Address) error {
	return nil
}

