package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

type ShippingCostService interface {
	CreateShippingCost(string, int) (bool, *errs.AppError)
	GetShippingCostByCity(string) (*ShippingCost, *errs.AppError)
	DeleteShippingCostByCity(string) (bool, *errs.AppError)
	UpdateShippingCost(ShippingCost) (bool, *errs.AppError)
}

type service2 struct {
	shippingCostRepository ShippingCostRepository
}

func (s service2) CreateShippingCost(city string, cost int) (bool, *errs.AppError) {
	shippingAddress := NewShippingCost(city, cost)
	_, err := s.shippingCostRepository.InsertShippingCost(*shippingAddress)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (s service2) GetShippingCostByCity(city string) (*ShippingCost, *errs.AppError) {
	res, err := s.shippingCostRepository.FindShippingCostByCity(city)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service2) DeleteShippingCostByCity(city string) (bool, *errs.AppError) {
	_, err := s.shippingCostRepository.DeleteShippingCostByCity(city)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s service2) UpdateShippingCost(sc ShippingCost) (bool, *errs.AppError) {
	_, err := s.shippingCostRepository.UpdateShippingCost(sc)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewShippingCostService(shippingCostRepository ShippingCostRepository) ShippingCostService {
	return &service2{
		shippingCostRepository: shippingCostRepository,
	}
}
