package app_module

import (
	"sync"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/nanda03dev/go-ms-template/src/application/service"
	"github.com/nanda03dev/go-ms-template/src/domain"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/repositories"
	"github.com/nanda03dev/go-ms-template/src/interface/controllers"
)

type Repository struct {
	UserRepository  domain.UserRepository
	OrderRepository domain.OrderRepository
}

type Service struct {
	HealthService service.HealthService
	UserService   service.UserService
	OrderService  service.OrderService
}

type Controller struct {
	HealthController controllers.HealthController
	UserController   controllers.UserController
	OrderController  controllers.OrderController
}

type AppModule struct {
	Service    Service
	Repository Repository
	Controller Controller
}

var (
	once      sync.Once
	appModule *AppModule
)

func GetAppModule() *AppModule {
	once.Do(func() {
		var databases = db.GetDBConnection()

		healthService := service.NewHealthService(databases)
		healthController := controllers.NewHealthController(healthService)

		userRepository := repositories.NewUserRepository(databases)
		userService := service.NewUserService(userRepository)
		userController := controllers.NewUserController(userService)

		orderRepository := repositories.NewOrderRepository(databases)
		orderService := service.NewOrderService(orderRepository)
		orderController := controllers.NewOrderController(orderService)

		appModule = &AppModule{
			Repository: Repository{
				UserRepository:  userRepository,
				OrderRepository: orderRepository,
			},
			Service: Service{
				HealthService: healthService,
				UserService:   userService,
				OrderService:  orderService,
			},
			Controller: Controller{
				HealthController: healthController,
				UserController:   userController,
				OrderController:  orderController},
		}
	})
	return appModule
}
