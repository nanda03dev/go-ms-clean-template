package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/app_module"
)

func InitializeRoutes(ginRouter *gin.Engine) {

	var appModule = app_module.GetAppModule()

	healthHandler := appModule.Handler.HealthHandler
	userHandler := appModule.Handler.UserHandler
	orderHandler := appModule.Handler.OrderHandler

	ginRouter.GET("/health", healthHandler.Health)

	userRoutes := ginRouter.Group("/user")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
	}

	orderRoutes := ginRouter.Group("/order")
	{
		orderRoutes.POST("/", orderHandler.CreateOrder)
		orderRoutes.GET("/:id", orderHandler.GetOrderByID)
	}
}
