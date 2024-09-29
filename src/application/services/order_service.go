package services

import (
	"github.com/nanda03dev/go-ms-template/src/domain/aggregates"
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
	newOrder := aggregates.NewOrder(createOrderDTO)

	err := s.orderRepo.Save(newOrder)
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}

func (s *orderService) GetOrderById(id string) (*aggregates.Order, error) {
	return s.orderRepo.FindById(id)
}
