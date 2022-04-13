package domain

import (
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team4/order/utils/errs"
)

type Order struct {
	Id               int            `json:"id"`
	UserID           int            `json:"user_id"`
	Status           string         `json:"status"`
	DateTime         time.Time      `json:"date_time"`
	ProductsQuantity map[string]int `json:"product_quantity"`
	ProductsCost     map[string]int `json:"product_cost"`
	TotalCost        int            `json:"total_cost"`
}

type OrderRepository interface {
	InsertOrder(Order) (Order, *errs.AppError)
	FindOrderById(int) (*Order, *errs.AppError)
	FindOrderByUserId(int) (*Order, *errs.AppError)
	DeleteOrderById(int) *errs.AppError
	UpdateOrder(Order) (*Order, *errs.AppError)
}

func NewOrder(userId int, status string, products_quantity map[string]int, products_cost map[string]int, total_cost int) *Order {
	return &Order{
		UserID:           userId,
		Status:           status,
		DateTime:         time.Now(),
		ProductsQuantity: products_quantity,
		ProductsCost:     products_cost,
		TotalCost:        total_cost,
	}
}
