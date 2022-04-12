package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

type ShippingCost struct {
	Id           int    `json:"id"`
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
	InsertShippingCost(ShippingCost) (ShippingCost, *errs.AppError)
	FindShippingCostById(int) (*ShippingCost, *errs.AppError)
	DeleteShippingCostById(int) *errs.AppError
	UpdateShippingCost(ShippingCost) (*ShippingCost, *errs.AppError)
}
