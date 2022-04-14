package domain

import (
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team4/order/utils/errs"
)

type Order struct {
	ID               string         `json:"id"`
	UserID           string         `json:"user_id"`
	Status           string         `json:"status"`
	DateTime         time.Time      `json:"date_time"`
	ProductsQuantity map[string]int `json:"product_quantity"`
	ProductsCost     map[string]int `json:"product_cost"`
	TotalCost        int            `json:"total_cost"`
}

type OrderRepository interface {
	InsertOrder(Order) (Order, *errs.AppError)
	FindOrderById(string) (*Order, *errs.AppError)
	FindOrderByUserId(string) ([]Order, *errs.AppError)
	DeleteOrderById(string) *errs.AppError
	UpdateOrderStatus(string, string) (*Order, *errs.AppError)
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
