package dto

type CreateOrderDTO struct {
	UserID   int     `json:"userID"`
	ItemName string  `json:"item_name"`
	Amount   float64 `json:"amount"`
}
