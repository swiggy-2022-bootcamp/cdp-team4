package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/cart/utils/errs"

type CartService interface {
	CreateCart(string,map[string]int,map[string]int) (string, *errs.AppError)
	GetCartById(string) (*Cart, *errs.AppError)
	GetAllCarts() ([]Cart, *errs.AppError)
	DeleteCartById(string) (bool, *errs.AppError)
}

type service struct {
	cartRepository CartRepository
}

func (s service) CreateCart(userId string, products_quantity map[string]int, products_cost map[string]int) (string, *errs.AppError) {
	cart := NewCart(userId, products_quantity, products_cost)
	resultid, err := s.cartRepository.InsertCart(*cart)
	if err != nil {
		return "", err
	}
	return resultid, nil
}

func (s service) GetCartById(cartId string) (*Cart, *errs.AppError) {
	res, err := s.cartRepository.FindCartById(cartId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) GetAllCarts() ([]Cart, *errs.AppError) {
	res, err := s.cartRepository.FindAllCarts()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) DeleteCartById(cartId string) (bool, *errs.AppError) {
	_, err := s.cartRepository.DeleteCartById(cartId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewCartService(cartRepository CartRepository) CartService {
	return &service{
		cartRepository: cartRepository,
	}
}
