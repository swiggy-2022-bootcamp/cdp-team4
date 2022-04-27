package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNewShippingAddress(t *testing.T) {
	firstname := "Naveen"
	lastname := "Kumar"
	city := "Banglore"
	address1 := "Address1"
	address2 := "Address2"
	countryid := 81
	postcode := 560063

	newShippingAddress := NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)

	assert.Equal(t, firstname, newShippingAddress.FirstName)
	assert.Equal(t, lastname, newShippingAddress.LastName)
	assert.Equal(t, city, newShippingAddress.City)
	assert.Equal(t, address1, newShippingAddress.Address1)
	assert.Equal(t, address2, newShippingAddress.Address2)
	assert.Equal(t, countryid, newShippingAddress.CountryID)
	assert.Equal(t, postcode, newShippingAddress.PostCode)
}
