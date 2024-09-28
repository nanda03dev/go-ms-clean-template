package dto

type CreateOrderDTO struct {
	UserID   string  `json:"userID"`
	ItemName string  `json:"itemName"`
	Amount   float64 `json:"amount"`
}
