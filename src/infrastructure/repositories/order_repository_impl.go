package repositories

import (
	"context"

	"github.com/nanda03dev/go-ms-template/src/domain/aggregate"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/entity"
)

type OrderRepositoryImpl struct {
	dbs *db.Databases
}

func NewOrderRepository(dbs *db.Databases) aggregate.OrderRepository {
	return &OrderRepositoryImpl{dbs: dbs}
}

func (r *OrderRepositoryImpl) Save(order *aggregate.Order) error {
	query := `INSERT INTO orders (order_id, user_id, item_name, amount) VALUES ($1, $2, $3, $4)`
	_, err := r.dbs.PostgresDB.DB.ExecContext(context.TODO(), query, order.ID, order.UserID, order.ItemName, order.Amount)
	return err
}

func (r *OrderRepositoryImpl) FindById(id string) (*aggregate.Order, error) {
	var order entity.Order

	query := `SELECT order_id, user_id, item_name, amount FROM orders WHERE order_id = $1`
	err := r.dbs.PostgresDB.DB.QueryRowContext(context.TODO(), query, id).Scan(&order.OrderID, &order.UserID, &order.Amount, &order.ItemName)

	if err != nil {
		return nil, err
	}

	// Convert entity.Order to domain.Order
	return convertEntityOrderToDomainOrder(&order), nil
}

// Convert entity.Order to domain.Order
func convertEntityOrderToDomainOrder(eo *entity.Order) *aggregate.Order {
	return &aggregate.Order{
		ID:       eo.OrderID,
		UserID:   eo.UserID,
		Amount:   eo.Amount,
		ItemName: eo.ItemName,
	}
}
