package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/application/service"
)

type HealthHandler interface {
	Health(ctx *gin.Context)
}

type healthHandler struct {
	healthService service.HealthService
}

func NewHealthHandler(healthService service.HealthService) HealthHandler {
	return &healthHandler{
		healthService: healthService,
	}
}

func (c *healthHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.healthService.Health())
}
