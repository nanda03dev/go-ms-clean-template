package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/application/service"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/interface/controllers"
)

func InitializeRoutes(ginRouter *gin.Engine, databases *db.Databases) {

	healthService := service.NewHealthService(databases)
	healthController := controllers.NewHealthController(healthService)

	ginRouter.GET("/health", healthController.Health)
}
