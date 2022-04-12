package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

type ShippingAddressService interface {
	CreateShippingAddress(string, string, string, string, string, int, int) (ShippingAddress, *errs.AppError)
	GetShippingAddressById(int) (*ShippingAddress, *errs.AppError)
	DeleteShippingAddressById(int) *errs.AppError
	UpdateShippingAddress(ShippingAddress) (*ShippingAddress, *errs.AppError)
}

type service1 struct {
	shippingAddressRepository ShippingAddressRepository
}

func (s service1) CreateShippingAddress(firstname, lastname, city, address1, address2 string, countryid, postcode int) (ShippingAddress, *errs.AppError) {
	shippingAddress := NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)
	persistedShippingAddress, err := s.shippingAddressRepository.InsertShippingAddress(*shippingAddress)
	if err != nil {
		return ShippingAddress{}, err
	}
	return persistedShippingAddress, nil
}
func (s service1) GetShippingAddressById(shippingAddressId int) (*ShippingAddress, *errs.AppError) {
	res, err := s.shippingAddressRepository.FindShippingAddressById(shippingAddressId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service1) DeleteShippingAddressById(shippingAddressId int) *errs.AppError {
	err := s.shippingAddressRepository.DeleteShippingAddressById(shippingAddressId)
	if err != nil {
		return err
	}
	return nil
}

func (s service1) UpdateShippingAddress(sa ShippingAddress) (*ShippingAddress, *errs.AppError) {
	res, err := s.shippingAddressRepository.UpdateShippingAddress(sa)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewShippingAddressService(shippingAddressRepository ShippingAddressRepository) ShippingAddressService {
	return &service1{
		shippingAddressRepository: shippingAddressRepository,
	}
}
