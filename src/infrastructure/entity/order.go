package entity

import "time"

// Order represents the PostgreSQL Order entity
type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ItemName  string    `json:"item_name"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
