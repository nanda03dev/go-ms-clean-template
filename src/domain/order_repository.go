package domain

type OrderRepository interface {
	Save(user *Order) error
	FindById(id string) (*Order, error)
	// GetOrdersByUserID(userID int) ([]Order, error)
}
