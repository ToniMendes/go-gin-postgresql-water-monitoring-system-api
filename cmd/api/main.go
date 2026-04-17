package main

import (
	"context"
	"go-gin-postgresql-water-monitoring-system-api/internal/configs"
	"go-gin-postgresql-water-monitoring-system-api/internal/infra/database"
	"go-gin-postgresql-water-monitoring-system-api/internal/infra/database/postgresql"
	"go-gin-postgresql-water-monitoring-system-api/internal/usecase/writeonly"
	"go-gin-postgresql-water-monitoring-system-api/internal/web"
	"go-gin-postgresql-water-monitoring-system-api/internal/worker/watermonitoring"
	"log"
	"time"
)

func main() {
	err := configs.StartConfig()
	returnFatalError(err)

	db, err := startDataBase()
	returnFatalError(err)

	repo := postgresql.NewPgSQLRepo(db.ClientPgSQL)

	usecaseCreate := writeonly.NewCreateUseCase(repo)
	usecaseUpdate := writeonly.NewUpdateUseCase(repo)


	wm := watermonitoring.NewWaterMonitoring(repo)

	go func() {
		ticker := time.NewTicker(90 * time.Second)
		for range ticker.C {
			err := wm.RecordWaterConsumption()
			if err != nil {
				log.Printf("Error: %v", err)
			}
		}
	}()

	type handler struct {
		*writeonly.CreateUseCase
		*writeonly.UpdateUseCase
	}

	hub := handler{
		usecaseCreate,
		usecaseUpdate,
	}

	endpoints := web.NewHandler(hub)

	web.Routers(endpoints)
}

func startDataBase() (*database.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := postgresql.NewPgSQLPool(ctx, configs.Env.DBURL)
	if err != nil {
		return nil, err
	}

	db := database.NewDatabase(pool)

	return db, nil
}

func returnFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
