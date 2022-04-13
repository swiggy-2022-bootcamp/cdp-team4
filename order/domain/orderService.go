package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/order/utils/errs"

type OrderService interface {
	CreateOrder(int, string, map[string]int, map[string]int, int) (Order, *errs.AppError)
	GetOrderById(int) (*Order, *errs.AppError)
	GetOrderByUserId(int) (*Order, *errs.AppError)
	DeleteOrderById(int) *errs.AppError
	UpdateOrder(Order) (*Order, *errs.AppError)
}

type service struct {
	orderRepository OrderRepository
}

func (s service) CreateOrder(userId int, status string, products_quantity map[string]int, products_cost map[string]int, totalcost int) (Order, *errs.AppError) {
	order := NewOrder(userId, status, products_quantity, products_cost, totalcost)
	persistedOrder, err := s.orderRepository.InsertOrder(*order)
	if err != nil {
		return Order{}, err
	}
	return persistedOrder, nil
}
func (s service) GetOrderById(orderId int) (*Order, *errs.AppError) {
	res, err := s.orderRepository.FindOrderById(orderId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) GetOrderByUserId(userId int) (*Order, *errs.AppError) {
	res, err := s.orderRepository.FindOrderByUserId(userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) DeleteOrderById(orderId int) *errs.AppError {
	err := s.orderRepository.DeleteOrderById(orderId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateOrder(sa Order) (*Order, *errs.AppError) {
	res, err := s.orderRepository.UpdateOrder(sa)
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
