package infra

import (
	"time"
)

type OrderModel struct {
	ID               string         `json:"id"`
	UserID           string         `json:"user_id"`
	Status           string         `json:"order_status"`
	DateTime         time.Time      `json:"dateTime"`
	ProductsQuantity map[string]int `json:"products_quantity"`
	ProductsCost     map[string]int `json:"products_cost"`
	TotalCost        int            `json:"total_cost"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}
