package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

type ShippingCostService interface {
	CreateShippingCost(string, int) (ShippingCost, *errs.AppError)
	GetShippingCostById(int) (*ShippingCost, *errs.AppError)
	DeleteShippingCostById(int) *errs.AppError
	UpdateShippingCost(ShippingCost) (*ShippingCost, *errs.AppError)
}

type service2 struct {
	shippingCostRepository ShippingCostRepository
}

func (s service2) CreateShippingCost(city string, cost int) (ShippingCost, *errs.AppError) {
	shippingAddress := NewShippingCost(city, cost)
	persistedShippingCost, err := s.shippingCostRepository.InsertShippingCost(*shippingAddress)
	if err != nil {
		return ShippingCost{}, err
	}
	return persistedShippingCost, nil
}
func (s service2) GetShippingCostById(shippingCostId int) (*ShippingCost, *errs.AppError) {
	res, err := s.shippingCostRepository.FindShippingCostById(shippingCostId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service2) DeleteShippingCostById(shippingCostId int) *errs.AppError {
	err := s.shippingCostRepository.DeleteShippingCostById(shippingCostId)
	if err != nil {
		return err
	}
	return nil
}

func (s service2) UpdateShippingCost(sc ShippingCost) (*ShippingCost, *errs.AppError) {
	res, err := s.shippingCostRepository.UpdateShippingCost(sc)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewShippingCostService(shippingCostRepository ShippingCostRepository) ShippingCostService {
	return &service2{
		shippingCostRepository: shippingCostRepository,
	}
}
