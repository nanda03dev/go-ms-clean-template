package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/application/service"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/repositories"
	"github.com/nanda03dev/go-ms-template/src/interface/controllers"
)

func InitializeRoutes(ginRouter *gin.Engine, dbs *db.Databases) {

	healthService := service.NewHealthService(dbs)
	healthController := controllers.NewHealthController(healthService)

	userRepository := repositories.NewUserRepository(dbs)
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	orderRepository := repositories.NewOrderRepository(dbs)
	orderService := service.NewOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)

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
