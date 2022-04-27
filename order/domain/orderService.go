package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/order/utils/errs"

type OrderService interface {
	CreateOrder(string, string, map[string]int, map[string]int, int) (string, *errs.AppError)
	GetOrderById(string) (*Order, *errs.AppError)
	GetOrderByUserId(string) ([]Order, *errs.AppError)
	GetOrderByStatus(string) ([]Order, *errs.AppError)
	GetAllOrders() ([]Order, *errs.AppError)
	DeleteOrderById(string) (bool, *errs.AppError)
	UpdateOrderStatus(string, string) (bool, *errs.AppError)
}

type OrderOverviewService interface {
	CreateOrderOverview(OrderOverview) (bool, *errs.AppError)
	GetOrderOverviewByOrderID(string) (*OrderOverview, *errs.AppError)
}

type service struct {
	orderRepository OrderRepository
}

type service1 struct {
	orderOverviewrepo OrderOverviewRepository
}

var order_status []string = []string{"confirmed", "declined", "cancelled", "pending", "delivered"}

func (s service) CreateOrder(userId string, status string, products_quantity map[string]int, products_cost map[string]int, totalcost int) (string, *errs.AppError) {
	order := NewOrder(userId, status, products_quantity, products_cost, totalcost)
	resultid, err := s.orderRepository.InsertOrder(*order)
	if err != nil {
		return "", err
	}
	return resultid, nil
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

func (s service) GetOrderByStatus(status string) ([]Order, *errs.AppError) {
	if !stringInSlice(status, order_status) {
		return nil, errs.NewNotFoundError("Not a valid status")
	}
	res, err := s.orderRepository.FindOrderByStatus(status)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) GetAllOrders() ([]Order, *errs.AppError) {
	res, err := s.orderRepository.FindAllOrders()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) DeleteOrderById(orderId string) (bool, *errs.AppError) {
	_, err := s.orderRepository.DeleteOrderById(orderId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s service) UpdateOrderStatus(id string, status string) (bool, *errs.AppError) {
	if !stringInSlice(status, order_status) {
		return false, errs.NewNotFoundError("Not a valid status")
	}
	_, err := s.orderRepository.UpdateOrderStatus(id, status)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s service1) CreateOrderOverview(ov OrderOverview) (bool, *errs.AppError) {
	ovn := Neworderoverview(ov.OrderID, ov.ProductsIdQuantity)
	_, err := s.orderOverviewrepo.InsertOrderOverview(*ovn)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s service1) GetOrderOverviewByOrderID(orderid string) (*OrderOverview, *errs.AppError) {
	result, err := s.orderOverviewrepo.GetOrderOverview(orderid)
	if err != nil {
		return &OrderOverview{}, err
	}
	return result, nil
}

func NewOrderService(orderRepository OrderRepository) OrderService {
	return &service{
		orderRepository: orderRepository,
	}
}

func NewOrderOverviewService(orderOverviewrepo OrderOverviewRepository) OrderOverviewService {
	return &service1{
		orderOverviewrepo: orderOverviewrepo,
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
