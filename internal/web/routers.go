// Package web provides HTTP routing and handlers for the water monitoring system API.
package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(endpoint *Handler) {
	r := gin.Default()

	r.POST("/api/water-monitoring/add-new-residence", endpoint.AddNewResidence)
	r.PUT("/api/water-monitoring/update-residence/:id", endpoint.UpdateOwner)
	
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

	r.Run()
}
