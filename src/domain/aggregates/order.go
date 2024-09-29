package aggregates

import (
	"github.com/nanda03dev/go-ms-template/src/helper"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

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

func NewOrder(createOrderDTO dto.CreateOrderDTO) *Order {
	return &Order{
		ID:       helper.Generate16DigitUUID(), // Generate unique ID (UUID or similar)
		UserID:   createOrderDTO.UserID,
		ItemName: createOrderDTO.ItemName,
		Amount:   createOrderDTO.Amount,
	}
}
