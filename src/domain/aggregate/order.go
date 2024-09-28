package aggregate

type Order struct {
	ID       string
	UserID   string
	ItemName string
	Amount   float64
}

type OrderRepository interface {
	Save(user *Order) error
	FindById(id string) (*Order, error)
	// GetOrdersByUserID(userID int) ([]Order, error)
}
