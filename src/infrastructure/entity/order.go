package entity

import "time"

// Order represents the PostgreSQL Order entity
type Order struct {
	OrderID   string    `json:"order_id"`
	UserID    string    `json:"user_id"`
	ItemName  string    `json:"item_name"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
