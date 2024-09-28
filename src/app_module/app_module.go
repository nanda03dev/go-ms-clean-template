package app_module

import (
	"sync"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/nanda03dev/go-ms-template/src/application/service"
	"github.com/nanda03dev/go-ms-template/src/domain/aggregate"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/repositories"
	"github.com/nanda03dev/go-ms-template/src/interface/handlers"
)

type Repository struct {
	UserRepository  aggregate.UserRepository
	OrderRepository aggregate.OrderRepository
}

type Service struct {
	HealthService service.HealthService
	UserService   service.UserService
	OrderService  service.OrderService
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

func GetAppModule() *AppModule {
	once.Do(func() {
		var databases = db.GetDBConnection()

		healthService := service.NewHealthService(databases)
		healthHandler := handlers.NewHealthHandler(healthService)

		userRepository := repositories.NewUserRepository(databases)
		userService := service.NewUserService(userRepository)
		userHandler := handlers.NewUserHandler(userService)

		orderRepository := repositories.NewOrderRepository(databases)
		orderService := service.NewOrderService(orderRepository)
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
