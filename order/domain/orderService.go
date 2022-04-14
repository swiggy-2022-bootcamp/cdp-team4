package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/order/utils/errs"

type OrderService interface {
	CreateOrder(string, string, map[string]int, map[string]int, int) (Order, *errs.AppError)
	GetOrderById(string) (*Order, *errs.AppError)
	GetOrderByUserId(string) ([]Order, *errs.AppError)
	DeleteOrderById(string) *errs.AppError
	UpdateOrderStatus(string, string) (*Order, *errs.AppError)
}

type service struct {
	orderRepository OrderRepository
}

func (s service) CreateOrder(userId string, status string, products_quantity map[string]int, products_cost map[string]int, totalcost int) (Order, *errs.AppError) {
	order := NewOrder(userId, status, products_quantity, products_cost, totalcost)
	persistedOrder, err := s.orderRepository.InsertOrder(*order)
	if err != nil {
		return Order{}, err
	}
	return persistedOrder, nil
}
func (s service) GetOrderById(orderId string) (*Order, *errs.AppError) {
	res, err := s.orderRepository.FindOrderById(orderId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) GetOrderByUserId(userId string) ([]Order, *errs.AppError) {
	res, err := s.orderRepository.FindOrderByUserId(userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) DeleteOrderById(orderId string) *errs.AppError {
	err := s.orderRepository.DeleteOrderById(orderId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateOrderStatus(id string, status string) (*Order, *errs.AppError) {
	res, err := s.orderRepository.UpdateOrderStatus(id, status)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewOrderService(orderRepository OrderRepository) OrderService {
	return &service{
		orderRepository: orderRepository,
	}
}
