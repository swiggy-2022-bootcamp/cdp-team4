package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

type ShippingAddress struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	CountryID int    `json:"country_id"`
	PostCode  int    `json:"postcode"`
}

func NewShippingAddress(firstname, lastname, city, address1, address2 string, countryid, postcode int) *ShippingAddress {
	return &ShippingAddress{
		FirstName: firstname,
		LastName:  lastname,
		City:      city,
		Address1:  address1,
		Address2:  address2,
		CountryID: countryid,
		PostCode:  postcode,
	}
}

type ShippingAddressRepository interface {
	InsertShippingAddress(ShippingAddress) (string, *errs.AppError)
	FindShippingAddressById(string) (*ShippingAddress, *errs.AppError)
	DeleteShippingAddressById(string) (bool, *errs.AppError)
	UpdateShippingAddressById(string, ShippingAddress) (bool, *errs.AppError)
}
