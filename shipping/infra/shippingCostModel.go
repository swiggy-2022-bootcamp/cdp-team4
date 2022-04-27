package infra

import "time"

type ShippingCostModel struct {
	City      string    `json:"city"`
	Cost      int       `json:"shipping_cost"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
