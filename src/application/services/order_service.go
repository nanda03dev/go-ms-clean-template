package services

import (
	"github.com/nanda03dev/go-ms-template/src/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/helper"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type OrderService interface {
	CreateOrder(order dto.CreateOrderDTO) (*aggregates.Order, error)
	GetOrderById(id string) (*aggregates.Order, error)
}

type orderService struct {
	orderRepo aggregates.OrderRepository
}

func NewOrderService(orderRepo aggregates.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (s *orderService) CreateOrder(createOrderDTO dto.CreateOrderDTO) (*aggregates.Order, error) {
	newUser := &aggregates.Order{
		ID:       helper.Generate16DigitUUID(), // Generate unique ID (UUID or similar)
		UserID:   createOrderDTO.UserID,
		ItemName: createOrderDTO.ItemName,
		Amount:   createOrderDTO.Amount,
	}

	err := s.orderRepo.Save(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *orderService) GetOrderById(id string) (*aggregates.Order, error) {
	return s.orderRepo.FindById(id)
}
