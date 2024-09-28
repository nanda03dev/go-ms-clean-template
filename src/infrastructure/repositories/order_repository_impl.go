package repositories

import (
	"errors"

	domain "github.com/nanda03dev/go-ms-template/src/domain/order"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
)

// In-memory storage for simplicity; replace with actual DB connection
var orders = map[string]*domain.Order{}

type OrderRepositoryImpl struct {
	dbs *db.Databases
}

func NewOrderRepository(dbs *db.Databases) domain.OrderRepository {
	return &OrderRepositoryImpl{dbs: dbs}
}

func (r *OrderRepositoryImpl) Save(order *domain.Order) error {
	orders[order.ID] = order
	return nil
}

func (r *OrderRepositoryImpl) FindById(id string) (*domain.Order, error) {
	order, exists := orders[id]
	if !exists {
		return nil, errors.New("order not found")
	}
	return order, nil
}
