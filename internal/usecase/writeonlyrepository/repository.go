// Package writeonlyrepository provides write-only repository use cases for water monitoring.
package writeonlyrepository


import "go-gin-postgresql-water-monitoring-system-api/internal/usecase/dto"

type ICreateUseCase interface {
	ExecCreate(dto.WaterMonitoringInput) (dto.WaterMonitoringOutput, error)
}
type IUpdateOwnerUseCase interface {
	ExecUpdateOwner(dto.WaterMonitoringInput, int64) (dto.WaterMonitoringOutput,error)
}
