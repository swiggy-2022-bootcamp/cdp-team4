package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

type ShippingCost struct {
	City         string `json:"city"`
	ShippingCost int    `json:"shipping_cost"`
}

func NewShippingCost(city string, shippingcost int) *ShippingCost {
	return &ShippingCost{
		City:         city,
		ShippingCost: shippingcost,
	}
}

type ShippingCostRepository interface {
	InsertShippingCost(ShippingCost) (bool, *errs.AppError)
	FindShippingCostByCity(string) (*ShippingCost, *errs.AppError)
	DeleteShippingCostByCity(string) (bool, *errs.AppError)
	UpdateShippingCost(ShippingCost) (bool, *errs.AppError)
}
