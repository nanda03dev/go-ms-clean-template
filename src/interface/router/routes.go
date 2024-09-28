package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/app_module"
)

func InitializeRoutes(ginRouter *gin.Engine) {

	var appModule = app_module.GetAppModule()

	healthController := appModule.Controller.HealthController
	userController := appModule.Controller.UserController
	orderController := appModule.Controller.OrderController

	ginRouter.GET("/health", healthController.Health)

	userRoutes := ginRouter.Group("/user")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/:id", userController.GetUserByID)
	}

	orderRoutes := ginRouter.Group("/order")
	{
		orderRoutes.POST("/", orderController.CreateOrder)
		orderRoutes.GET("/:id", orderController.GetOrderByID)
	}
}
