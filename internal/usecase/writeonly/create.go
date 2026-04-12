// Package writeonly provides repository interfaces for write-only operations.
package writeonly

import (
	"go-gin-postgresql-water-monitoring-system-api/internal/domain"
	"go-gin-postgresql-water-monitoring-system-api/internal/domain/entities"
	"go-gin-postgresql-water-monitoring-system-api/internal/usecase/dto"
)

type CreateUseCase struct {
	conn domain.PgSQLRepository
}

func NewCreateUseCase(c domain.PgSQLRepository) *CreateUseCase {
	return &CreateUseCase{
		conn: c,
	}
}

func (c *CreateUseCase) Execute(dto.WaterMonitoringInput) (dto.WaterMonitoringOutput, error) {
	owner := entities.NewOwner(dto.OwnerName, dto.Email, dto.Phone)

}
