package service

import (
	domain "github.com/nanda03dev/go-ms-template/src/domain/order"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type OrderService interface {
	CreateOrder(order dto.CreateOrderDTO) (*domain.Order, error)
	GetOrderById(id string) (*domain.Order, error)
}

type orderService struct {
	orderRepo domain.OrderRepository
}

func NewOrderService(orderRepo domain.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (s *orderService) CreateOrder(createOrderDTO dto.CreateOrderDTO) (*domain.Order, error) {
	newUser := &domain.Order{
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

func (s *orderService) GetOrderById(id string) (*domain.Order, error) {
	return s.orderRepo.FindById(id)
}
