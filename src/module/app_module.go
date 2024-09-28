package module

import (
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/nanda03dev/go-ms-template/src/application/service"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
)

type AppModule struct {
	HealthService service.HealthService
}

// NewAppModule initializes all the repositories and services
func NewAppModule(databases *db.Databases) *AppModule {
	// Initialize repositories
	// userRepo := infrastructure.NewUserRepository(db)
	// orderRepo := infrastructure.NewOrderRepository(db)

	// Initialize services
	// orderService := application.NewOrderService(orderRepo)
	healthService := service.NewHealthService(databases)

	return &AppModule{
		HealthService: healthService,
	}
}
