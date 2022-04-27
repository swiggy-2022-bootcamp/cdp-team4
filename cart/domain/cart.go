package domain

import (
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/utils/errs"
)

type Item struct {
	Name     string `json:"name"`
	Cost     int `json:"cost"`
	Quantity int `json:"quantity"`
}
type Cart struct {
	Id     string          `json:"id"`
	UserID string          `json:"user_id"`
	Items  map[string]Item `json:"items"`
}

type CartRepository interface {
	InsertCart(Cart) (string, *errs.AppError)
	UpdateCartByUserId(string,map[string]Item) (bool, *errs.AppError)
	DeleteCartItemByUserId(string,[]string) (bool, *errs.AppError)
	FindAllCarts() ([]Cart, *errs.AppError)
	FindCartById(string) (*Cart, *errs.AppError)
	FindCartByUserId(string) (*Cart, *errs.AppError)
	DeleteCartByUserId(string) (bool, *errs.AppError)
	DeleteCartById(string) (bool, *errs.AppError)
}

func NewCart(userId string, productsList map[string]Item) *Cart {
	return &Cart{
		UserID: userId,
		Items:  productsList,
	}
}
