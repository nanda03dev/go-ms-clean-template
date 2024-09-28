package service

import (
	"github.com/nanda03dev/go-ms-template/src/domain/aggregate"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type OrderService interface {
	CreateOrder(order dto.CreateOrderDTO) (*aggregate.Order, error)
	GetOrderById(id string) (*aggregate.Order, error)
}

type orderService struct {
	orderRepo aggregate.OrderRepository
}

func NewOrderService(orderRepo aggregate.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (s *orderService) CreateOrder(createOrderDTO dto.CreateOrderDTO) (*aggregate.Order, error) {
	newUser := &aggregate.Order{
		ID:       Generate16DigitUUID(), // Generate unique ID (UUID or similar)
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

func (s *orderService) GetOrderById(id string) (*aggregate.Order, error) {
	return s.orderRepo.FindById(id)
}
