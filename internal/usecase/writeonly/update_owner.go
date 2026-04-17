package writeonly

import (
	"go-gin-postgresql-water-monitoring-system-api/internal/domain"
	"go-gin-postgresql-water-monitoring-system-api/internal/domain/entities"
	"go-gin-postgresql-water-monitoring-system-api/internal/infra/services/viacep"
	"go-gin-postgresql-water-monitoring-system-api/internal/usecase/dto"
)

type UpdateUseCase struct {
	pool domain.PgSQLRepository
}

func NewUpdateUseCase(p domain.PgSQLRepository) *UpdateUseCase {
	return &UpdateUseCase{
		pool: p,
	}
}

func (r *UpdateUseCase) ExecUpdateOwner(input dto.WaterMonitoringInput, id int64) (dto.WaterMonitoringOutput, error) {
	owner, err := entities.NewOwner(input.OwnerName, input.Email, input.Phone)
	if err != nil {
		return dto.WaterMonitoringOutput{}, err
	}

	resp, err := viacep.NewQuery(input.CEP)
	if err != nil {
		return dto.WaterMonitoringOutput{}, err
	}

	address, err := entities.NewAddress(input.CEP, resp.PublicPlace, resp.Neighborhood, resp.State, resp.City, resp.Region)
	if err != nil {
		return dto.WaterMonitoringOutput{}, err
	}

	err = r.pool.UpdateOwner(owner, address, id)
	if err != nil {
		return dto.WaterMonitoringOutput{}, err
	}

	result, err := r.pool.GetByID(id)
	if err != nil {
		return dto.WaterMonitoringOutput{}, err
	}

	response := dto.WaterMonitoringOutput{
		ID:           result.ID,
		OwnerName:    result.OwnerName,
		Email:        result.Email,
		Phone:        result.Phone,
		CEP:          result.CEP,
		PublicPlace:  result.PublicPlace,
		Neighborhood: result.Neighborhood,
		State:        result.Uf,
		City:         result.City,
		Region:       result.Region,
	}

	return response, nil
}
