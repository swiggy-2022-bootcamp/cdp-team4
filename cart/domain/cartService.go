package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/cart/utils/errs"

type CartService interface {
	CreateCart(string,map[string]Item) (string, *errs.AppError)
	UpdateCartItemsByUserId(string,map[string]Item) (bool, *errs.AppError)
	DeleteCartItemByUserId(string,[]string) (bool, *errs.AppError)
	GetCartById(string) (*Cart, *errs.AppError)
	GetCartByUserId(string) (*Cart, *errs.AppError)
	GetAllCarts() ([]Cart, *errs.AppError)
	DeleteCartById(string) (bool, *errs.AppError)
	DeleteCartByUserId(string) (bool, *errs.AppError)
}

type service struct {
	cartRepository CartRepository
}

func (s service) CreateCart(userId string, domainItemMap map[string]Item) (string, *errs.AppError) {
	cart := NewCart(userId, domainItemMap)
	resultid, err := s.cartRepository.InsertCart(*cart)
	if err != nil {
		return "", err
	}
	return resultid, nil
}

func (s service) UpdateCartItemsByUserId(userId string, domainItemMap map[string]Item) (bool, *errs.AppError) {
	res, err := s.cartRepository.UpdateCartByUserId(userId,domainItemMap)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (s service) DeleteCartItemByUserId(userId string, productIdList []string) (bool, *errs.AppError) {
	res, err := s.cartRepository.DeleteCartItemByUserId(userId,productIdList)
	if err != nil {
		return false, err
	}
	return res, nil
}


func (s service) GetCartById(cartId string) (*Cart, *errs.AppError) {
	res, err := s.cartRepository.FindCartById(cartId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) GetCartByUserId(userId string) (*Cart, *errs.AppError) {
	res, err := s.cartRepository.FindCartByUserId(userId)
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

func (s service) DeleteCartByUserId(userId string) (bool, *errs.AppError) {
	_, err := s.cartRepository.DeleteCartByUserId(userId)
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
