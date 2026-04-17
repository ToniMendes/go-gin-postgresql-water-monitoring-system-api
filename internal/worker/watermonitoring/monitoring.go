// Package watermonitoring handles water monitoring and consumption tracking operations.
package watermonitoring

import (
	"go-gin-postgresql-water-monitoring-system-api/internal/domain"
	"log"
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
	tasks := make(chan int64, 1000)

	go func() {
		if err := r.repo.GetAllID(tasks); err != nil {
			log.Printf("Error: %v", err)
		}

		close(tasks)
	}()

	var wg sync.WaitGroup

	workes := 100

	for i := 0; i < workes; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range tasks {
				totalExpense := expenses()

				err := r.repo.UpadteWaterConsumption(totalExpense, int64(id))
				if err != nil {
					log.Printf("Error: %v", err)
				}
			}
		}()
	}

	wg.Wait()

	log.Printf("Ciclo de monitoramento finalizado: residências processadas.")
	return nil
}

func expenses() float64 {
	min := 0.12
	max := 0.37
	return min + rand.Float64()*(max-min)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
