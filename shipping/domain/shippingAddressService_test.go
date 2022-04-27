package domain_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

	"github.com/stretchr/testify/assert"
)

var mockShippingAddressRepo = mocks.ShippingAddressRepository{}
var shippingAddresService = domain.NewShippingAddressService(&mockShippingAddressRepo)

func TestShouldReturnNewShippingAddressService(t *testing.T) {
	userService := domain.NewShippingAddressService(nil)
	assert.NotNil(t, userService)
}

func TestShouldCreateNewShippingAddress(t *testing.T) {
	resultid := "94713094"
	firstname := "Naveen"
	lastname := "Kumar"
	city := "Banglore"
	address1 := "Address1"
	address2 := "Address2"
	countryid := 81
	postcode := 560063

	//newShippingAddress := domain.NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)

	mockShippingAddressRepo.On("InsertShippingAddress", mock.Anything).Return(resultid, nil)
	shippingAddresService.CreateShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)
	mockShippingAddressRepo.AssertNumberOfCalls(t, "InsertShippingAddress", 1)
}

func TestShouldDeleteShippingAddressById(t *testing.T) {
	shippinAddressId := "1"
	mockShippingAddressRepo.On("DeleteShippingAddressById", shippinAddressId).Return(true, nil)
	res, err := shippingAddresService.DeleteShippingAddressById(shippinAddressId)
	assert.Equal(t, res, true)
	assert.Nil(t, err)
}

func TestShouldGetShippingAddressById(t *testing.T) {
	shippinAddressId := "1"
	firstname := "Naveen"
	lastname := "Kumar"
	city := "Banglore"
	address1 := "Address1"
	address2 := "Address2"
	countryid := 81
	postcode := 560063

	newShippingAddress := domain.NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)

	mockShippingAddressRepo.On("FindShippingAddressById", shippinAddressId).Return(newShippingAddress, nil)
	resShippingAddress, _ := shippingAddresService.GetShippingAddressById(shippinAddressId)

	assert.Equal(t, firstname, resShippingAddress.FirstName)
	assert.Equal(t, lastname, resShippingAddress.LastName)
	assert.Equal(t, city, resShippingAddress.City)
	assert.Equal(t, address1, resShippingAddress.Address1)
	assert.Equal(t, address2, resShippingAddress.Address2)
	assert.Equal(t, countryid, resShippingAddress.CountryID)
	assert.Equal(t, postcode, resShippingAddress.PostCode)
}

func TestShouldNotDeleteShippingAddressByShippingAddressIdUponInvalidShippingAddressId(t *testing.T) {
	shippinAddressId := "-99"
	errMessage := "some error"
	mockShippingAddressRepo.On("DeleteShippingAddressById", shippinAddressId).Return(false, errs.NewUnexpectedError(errMessage))

	res, err := shippingAddresService.DeleteShippingAddressById(shippinAddressId)
	assert.Equal(t, res, false)
	assert.Error(t, err.Error(), errMessage)
}

func TestShouldUpdateShippingAddressById(t *testing.T) {
	firstname := "Naveen"
	lastname := "Kumar"
	city := "Banglore"
	address1 := "Address1"
	address2 := "Address2"
	countryid := 99
	postcode := 560063

	newShippingAddress := domain.NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)
	mockShippingAddressRepo.On("UpdateShippingAddressById", "112324", *newShippingAddress).Return(true, nil)
	res, err := shippingAddresService.UpdateShippingAddressById("112324", *newShippingAddress)

	assert.Equal(t, res, true)
	assert.Nil(t, err)
}
