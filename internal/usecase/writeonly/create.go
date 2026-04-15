// Package writeonly provides repository interfaces for write-only operations.
package writeonly

import (
	"go-gin-postgresql-water-monitoring-system-api/internal/domain"
	"go-gin-postgresql-water-monitoring-system-api/internal/domain/entities"
	"go-gin-postgresql-water-monitoring-system-api/internal/infra/services/viacep"
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

func (c *CreateUseCase) Execute(input dto.WaterMonitoringInput) (dto.WaterMonitoringOutput, error) {
	owner := entities.NewOwner(input.OwnerName, input.Email, input.Phone)

	resp, err := viacep.NewQuery(input.CEP)
	if err != nil {
		return dto.WaterMonitoringOutput{}, err
	}

	address := entities.NewAddress(input.CEP, resp.PublicPlace, resp.Neighborhood, resp.State, resp.City, resp.Region)

	err = c.conn.Save(owner, address)
	if err != nil {
		return dto.WaterMonitoringOutput{}, err
	}

	response := dto.WaterMonitoringOutput{
		OwnerName:    owner.OwnerName,
		Email:        owner.Email,
		Phone:        owner.Phone,
		CEP:          address.CEP,
		PublicPlace:  address.PublicPlace,
		Neighborhood: address.Neighborhood,
		State:        address.Uf,
		City:         address.City,
		Region:       address.Region,
	}

	return response, nil
}
