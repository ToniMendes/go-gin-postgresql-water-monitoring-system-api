// Package watermonitoring handles water monitoring and consumption tracking operations.
package watermonitoring

import (
	"context"
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

func (r *WaterMonitoring) RecordWaterConsumption(ctx context.Context) error {
	id := make(chan int64, 1000)

	go func() {
		if err := r.repo.GetAllID(id); err != nil {
			log.Printf("Error: %v", err)
		}

		close(id)
	}()

	var wg sync.WaitGroup

	workers := 100

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case id, ok := <-id:
					if !ok {
						return
					}
					inv := invoice()
					if err := r.repo.UpdateWaterConsumption(inv, id); err != nil {
						log.Printf("Error: %v", err)
					}
				}
			}
		}()
	}

	wg.Wait()

	log.Printf("Ciclo de monitoramento finalizado: residências processadas.")
	return nil
}

func (r *WaterMonitoring) NewMonitoring() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ticker := time.NewTicker(90 * time.Second)
	for range ticker.C {
		err := r.RecordWaterConsumption(ctx)
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}
}

func invoice() float64 {
	min := 0.12
	max := 0.37
	return min + rand.Float64()*(max-min)
}
