package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

type ShippingAddressService interface {
	CreateShippingAddress(string, string, string, string, string, int, int) (string, *errs.AppError)
	GetShippingAddressById(string) (*ShippingAddress, *errs.AppError)
	DeleteShippingAddressById(string) (bool, *errs.AppError)
	UpdateShippingAddressById(string, ShippingAddress) (bool, *errs.AppError)
}

type service1 struct {
	shippingAddressRepository ShippingAddressRepository
}

func (s service1) CreateShippingAddress(firstname, lastname, city, address1, address2 string, countryid, postcode int) (string, *errs.AppError) {
	shippingAddress := NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)
	resid, err := s.shippingAddressRepository.InsertShippingAddress(*shippingAddress)
	if err != nil {
		return "", err
	}
	return resid, nil
}
func (s service1) GetShippingAddressById(shippingAddressId string) (*ShippingAddress, *errs.AppError) {
	res, err := s.shippingAddressRepository.FindShippingAddressById(shippingAddressId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service1) DeleteShippingAddressById(shippingAddressId string) (bool, *errs.AppError) {
	_, err := s.shippingAddressRepository.DeleteShippingAddressById(shippingAddressId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s service1) UpdateShippingAddressById(shippingAddressId string, newShippingAddress ShippingAddress) (bool, *errs.AppError) {
	_, err := s.shippingAddressRepository.UpdateShippingAddressById(shippingAddressId, newShippingAddress)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewShippingAddressService(shippingAddressRepository ShippingAddressRepository) ShippingAddressService {
	return &service1{
		shippingAddressRepository: shippingAddressRepository,
	}
}
