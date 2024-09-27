package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/application/service"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/interface/controllers"
)

func InitializeRoutes(ginRouter *gin.Engine, mongoDB *db.MongoDB, postgresDB *db.PostgresDB) {

	healthService := service.NewHealthService(mongoDB, postgresDB)
	healthController := controllers.NewHealthController(healthService)

	ginRouter.GET("/health", healthController.Health)
}
