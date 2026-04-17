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

func (r *PgSQLRepo) UpdateOwner(owner *entities.Owner, address *entities.Address, id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE residences
		SET Owners_name = $1, email = $2, phone = $3, cep = $4, publicplace = $5, neighborhood = $6, uf = $7, city = $8, region = $9
		WHERE id = $10;
	`

	_, err := r.pool.Exec(ctx, query, owner.OwnerName, owner.Email, owner.Phone, address.CEP, address.City, address.Neighborhood, address.Uf, address.City, address.Region, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PgSQLRepo) GetAllID(ids chan<- int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
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

		select {
		case ids <- id:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r *PgSQLRepo) GetByID(id int64) (entities.Residence, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, owners_name, email, phone, cep, publicplace, 
               neighborhood, uf, city, region 
        FROM residences 
        WHERE id = $1;
	`

	var res entities.Residence

	err := r.pool.QueryRow(ctx, query, id).Scan(&res.ID, &res.OwnerName, &res.Email, &res.Phone, &res.CEP, &res.PublicPlace, &res.Neighborhood, &res.Uf, &res.City, &res.Region)
	if err != nil {
		return entities.Residence{}, err
	}

	return res, nil
}

func (r *PgSQLRepo) GetByEmail(email string) (entities.Residence, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, owners_name, email, phone, cep, publicplace, 
               neighborhood, uf, city, region 
        FROM residences 
        WHERE email = $1;`

	var res entities.Residence

	err := r.pool.QueryRow(ctx, query, email).Scan(&res.ID, &res.OwnerName, &res.Email, &res.Phone, &res.CEP, &res.PublicPlace, &res.Neighborhood, &res.Uf, &res.City, &res.Region)
	if err != nil {
		return entities.Residence{}, err
	}

	return res, nil
}
