package infra

import "time"

type ShippingAddressModel struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	City      string    `json:"city"`
	Address1  string    `json:"address_1"`
	Address2  string    `json:"address_2"`
	CountryID int       `json:"country_id"`
	PostCode  int       `json:"postcode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
