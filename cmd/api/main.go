package main

import (
	"go-gin-postgresql-water-monitoring-system-api/internal/configs"
	dependences "go-gin-postgresql-water-monitoring-system-api/internal/infra/dependeces"
	"go-gin-postgresql-water-monitoring-system-api/internal/web"
	"log"
)

func main() {
	err := configs.StartConfig()
	returnFatalError(err)

	app := dependences.NewApp()

	go app.Worker.NewMonitoring()

	web.Routers(app.WebHandler)
}

func returnFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
