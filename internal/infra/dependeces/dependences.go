// Package dependences orchestrates the application's dependency injection.
// It is responsible for initializing infrastructure resources (such as the database),
// instantiating repositories, use cases, and handlers, and assembling the
// final App structure for execution.
package dependences

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

type App struct {
	Worker     *watermonitoring.WaterMonitoring
	WebHandler *web.Handler
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

func NewApp() *App {
	db, err := startDataBase()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	repo := postgresql.NewPgSQLRepo(db.ClientPgSQL)

	usecaseCreate := writeonly.NewCreateUseCase(repo)
	usecaseUpdate := writeonly.NewUpdateUseCase(repo)

	type handler struct {
		*writeonly.CreateUseCase
		*writeonly.UpdateUseCase
	}

	hub := handler{
		usecaseCreate,
		usecaseUpdate,
	}

	wm := watermonitoring.NewWaterMonitoring(repo)
	hdl := web.NewHandler(hub)

	return &App{
		Worker:     wm,
		WebHandler: hdl,
	}
}
