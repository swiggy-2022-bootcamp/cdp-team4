package infra_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/infra"
)

var testShippingAddressService = infra.NewDynamoShippingAddressRepository()
var insertedid string

func TestShouldCreateNewShippingAddresssinDynamoDB(t *testing.T) {
	firstname := "Naveen"
	lastname := "Kumar"
	city := "Banglore"
	address1 := "Address1"
	address2 := "Address2"
	countryid := 81
	postcode := 560063

	newShippingAddress := domain.NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)
	res, err := testShippingAddressService.InsertShippingAddress(*newShippingAddress)
	insertedid = res
	t.Logf("Inserted Id is %s\n", insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldGetShippingAddresssByShippingAddresssIdDynamoDB(t *testing.T) {
	firstname := "Naveen"
	lastname := "Kumar"
	city := "Banglore"
	address1 := "Address1"
	address2 := "Address2"
	countryid := 81
	postcode := 560063
	//t.Logf("Inserted Id is %s Reading\n", insertedid)
	res, err := testShippingAddressService.FindShippingAddressById(insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, firstname, res.FirstName)
	assert.Equal(t, lastname, res.LastName)
	assert.Equal(t, city, res.City)
	assert.Equal(t, address1, res.Address1)
	assert.Equal(t, address2, res.Address2)
	assert.Equal(t, countryid, res.CountryID)
	assert.Equal(t, postcode, res.PostCode)
}

func TestShouldUpdateShippingAddresssStatusDynamoDB(t *testing.T) {
	firstname := "Naveen"
	lastname := "Kumar"
	city := "Banglore"
	address1 := "Address1"
	address2 := "Address2"
	countryid := 81
	postcode := 560012

	newShippingAddress := domain.NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)
	res, err := testShippingAddressService.UpdateShippingAddressById(insertedid, *newShippingAddress)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

// func TestShouldDeleteShippingAddresssByShippingAddresssIdDynamoDB(t *testing.T) {
// 	res, err := testShippingAddressService.DeleteShippingAddressById(insertedid)
// 	assert.NotNil(t, res)
// 	assert.Nil(t, err)
// }
