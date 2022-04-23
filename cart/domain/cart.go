package domain

import (
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/utils/errs"
)	

type Cart struct {
	Id               string         `json:"id"`
	UserID           string         `json:"user_id"`
	ProductsQuantity map[string]int `json:"products_quantity"`
}

type CartRepository interface {
	InsertCart(Cart) (string, *errs.AppError)
	FindAllCarts() ([]Cart, *errs.AppError)
	FindCartById(string) (*Cart, *errs.AppError)
	DeleteCartById(string) (bool, *errs.AppError)
}

func NewCart(userId string, products_quantity map[string]int) *Cart {
	return &Cart{
		UserID:           userId,
		ProductsQuantity: products_quantity,
	}
}
