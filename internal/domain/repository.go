// Package domain contains the domain models and interfaces for the water monitoring system.
package domain

import "go-gin-postgresql-water-monitoring-system-api/internal/domain/entities"

type PgSQLRepository interface {
	Save(*entities.Owner, *entities.Address) error
	UpdateOwner(*entities.Owner, *entities.Address, int64) error
	UpadteWaterConsumption(float64, int64) error
	GetAllID(chan<- int64) error
	GetByID(id int64) (entities.Residence, error)
	GetByEmail(email string) (entities.Residence, error)
}
