package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/application/services"
)

type HealthHandler interface {
	Health(ctx *gin.Context)
}

type healthHandler struct {
	healthService services.HealthService
}

func NewHealthHandler(healthService services.HealthService) HealthHandler {
	return &healthHandler{
		healthService: healthService,
	}
}

func (c *healthHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.healthService.Health())
}
