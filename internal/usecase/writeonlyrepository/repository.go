// Package writeonlyrepository provides write-only repository use cases for water monitoring.
package writeonlyrepository

import "go-gin-postgresql-water-monitoring-system-api/internal/usecase/dto"

type ICreateUseCase interface {
	Execute(dto.WaterMonitoringInput) (dto.WaterMonitoringOutput, error)
}
