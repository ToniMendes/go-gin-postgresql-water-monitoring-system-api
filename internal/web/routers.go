// Package web provides HTTP routing and handlers for the water monitoring system API.
package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(endpoint *Handler) {
	r := gin.Default()

	r.POST("/api/water-monitoring/add-new-address", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	r.Run()
}
