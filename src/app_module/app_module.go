package app_module

import (
	"sync"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/nanda03dev/go-ms-template/src/application/services"
	"github.com/nanda03dev/go-ms-template/src/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/repositories"
	"github.com/nanda03dev/go-ms-template/src/interface/handlers"
)

type Repository struct {
	UserRepository  aggregates.UserRepository
	OrderRepository aggregates.OrderRepository
}

type Service struct {
	HealthService services.HealthService
	UserService   services.UserService
	OrderService  services.OrderService
}

type Handler struct {
	HealthHandler handlers.HealthHandler
	UserHandler   handlers.UserHandler
	OrderHandler  handlers.OrderHandler
}

type AppModule struct {
	Service    Service
	Repository Repository
	Handler    Handler
}

var (
	once      sync.Once
	appModule *AppModule
)

func InitModules() *AppModule {
	return GetAppModule()
}

func GetAppModule() *AppModule {
	once.Do(func() {
		var databases = db.GetDBConnection()

		healthService := services.NewHealthService(databases)
		healthHandler := handlers.NewHealthHandler(healthService)

		userRepository := repositories.NewUserRepository(databases)
		userService := services.NewUserService(userRepository)
		userHandler := handlers.NewUserHandler(userService)

		orderRepository := repositories.NewOrderRepository(databases)
		orderService := services.NewOrderService(orderRepository)
		orderHandler := handlers.NewOrderHandler(orderService)

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
			Handler: Handler{
				HealthHandler: healthHandler,
				UserHandler:   userHandler,
				OrderHandler:  orderHandler},
		}
	})
	return appModule
}
