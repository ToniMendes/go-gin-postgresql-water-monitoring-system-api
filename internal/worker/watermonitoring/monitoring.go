// Package watermonitoring handles water monitoring and consumption tracking operations.
package watermonitoring

import (
	"go-gin-postgresql-water-monitoring-system-api/internal/domain"
	"math/rand"
	"sync"
	"time"
)

type WaterMonitoring struct {
	repo domain.PgSQLRepository
}

func NewWaterMonitoring(repo domain.PgSQLRepository) *WaterMonitoring {
	return &WaterMonitoring{
		repo: repo,
	}
}

func (r *WaterMonitoring) RecordWaterConsumption() error {
	ids, err := r.repo.GetAllID() 
	if err != nil {
		return err
	}

	tasks := make(chan int, len(ids))
	var wg sync.WaitGroup

	workes := 100

	for i := 0; i < workes; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			for id := range tasks {
				totalExpense := expenses()
				
				// falta implementar o restante da logica 
				_ = totalExpense
				_ = id

				
			}
		}()
	}

	return nil
}

func expenses() float64 {
	min := 0.43
	max := 2.91
	
	random := min + rand.Float64() * (max - min)

	return random
}
