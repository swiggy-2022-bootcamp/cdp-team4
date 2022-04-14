package infra

import (
	"time"
)

type UserModel struct {
	ID               int64          `json:"id"`
	UserID           int64          `json:"user_id"`
	Status           string         `json:"status"`
	DateTime         time.Time      `json:"date_time"`
	ProductsQuantity map[string]int `json:"product_quantity"`
	ProductsCost     map[string]int `json:"product_cost"`
	TotalCost        int            `json:"total_cost"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}
