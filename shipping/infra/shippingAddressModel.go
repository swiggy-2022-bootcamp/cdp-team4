package infra

import "time"

type ShippingAddressModel struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	City      string    `json:"city"`
	Address1  string    `json:"address1"`
	Address2  string    `json:"address2"`
	CountryID int       `json:"country_id"`
	PostCode  int       `json:"postcode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
