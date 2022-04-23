package domain

import (
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team4/order/utils/errs"
)

type Order struct {
	ID               string         `json:"id"`
	UserID           string         `json:"user_id"`
	Status           string         `json:"order_status"`
	DateTime         time.Time      `json:"date_time"`
	ProductsQuantity map[string]int `json:"product_quantity"`
	ProductsCost     map[string]int `json:"product_cost"`
	TotalCost        int            `json:"total_cost"`
}

type OrderOverview struct {
	OrderID            string         `json:"order_id"`
	ProductsIdQuantity map[string]int `json:"products"`
}

type OrderRepository interface {
	InsertOrder(Order) (string, *errs.AppError)
	FindOrderById(string) (*Order, *errs.AppError)
	FindOrderByUserId(string) ([]Order, *errs.AppError)
	FindOrderByStatus(string) ([]Order, *errs.AppError)
	FindAllOrders() ([]Order, *errs.AppError)
	DeleteOrderById(string) (bool, *errs.AppError)
	UpdateOrderStatus(string, string) (bool, *errs.AppError)
}

type OrderOverviewRepository interface {
	InsertOrderOverview(OrderOverview) (bool, *errs.AppError)
	GetOrderOverview(string) (*OrderOverview, *errs.AppError)
}

func NewOrder(userId string, status string, products_quantity map[string]int, products_cost map[string]int, total_cost int) *Order {
	return &Order{
		UserID:           userId,
		Status:           status,
		DateTime:         time.Now(),
		ProductsQuantity: products_quantity,
		ProductsCost:     products_cost,
		TotalCost:        total_cost,
	}
}

func Neworderoverview(orderId string, products map[string]int) *OrderOverview {
	return &OrderOverview{
		OrderID:            orderId,
		ProductsIdQuantity: products,
	}
}
