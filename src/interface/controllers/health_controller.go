package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/application/service"
)

type HealthController interface {
	Health(ctx *gin.Context)
}

type healthController struct {
	healthService service.HealthService
}

func NewHealthController(healthService service.HealthService) HealthController {
	return &healthController{
		healthService: healthService,
	}
}

func (c *healthController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.healthService.Health())
}
