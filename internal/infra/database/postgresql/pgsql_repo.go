package postgresql

import (
	"context"
	"go-gin-postgresql-water-monitoring-system-api/internal/domain/entities"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgSQLRepo struct {
	pool *pgxpool.Pool
}

func NewPgSQLRepo(pool *pgxpool.Pool) *PgSQLRepo {
	return &PgSQLRepo{
		pool: pool,
	}
}

func (r *PgSQLRepo) Save(owners *entities.Owner, address *entities.Address) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO residences (owners_name, email, phone, cep, publicplace, neighborhood, uf, city, region)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	_, err := r.pool.Exec(ctx, query, owners.OwnerName, owners.Email, owners.Phone, address.CEP, address.PublicPlace, address.Neighborhood, address.Uf, address.City, address.Region)
	if err != nil {
		return err
	}

	return nil
}

func (r *PgSQLRepo) UpadteWaterConsumption(value float64, id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE residences
		SET invoice = invoice + $1
		WHERE id = $2;
	`

	_, err := r.pool.Exec(ctx, query, value, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PgSQLRepo) GetAllID(ids chan<- int64) (error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.pool.Query(ctx, "SELECT id FROM residences")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return err
		}
		ids <- id
	}
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}
