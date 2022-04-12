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

	firstname := "Naveen"
	lastname := "Kumar"
	city := "Banglore"
	address1 := "Address1"
	address2 := "Address2"
	countryid := 81
	postcode := 560063

	newShippingAddress := domain.NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)

	mockShippingAddressRepo.On("InsertShippingAddress", mock.Anything).Return(*newShippingAddress, nil)
	shippingAddresService.CreateShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)
	mockShippingAddressRepo.AssertNumberOfCalls(t, "InsertShippingAddress", 1)
}

func TestShouldDeleteShippingAddressByShippingAddressId(t *testing.T) {
	shippingAdressId := 1
	mockShippingAddressRepo.On("DeleteShippingAddressById", shippingAdressId).Return(nil)
	var err = shippingAddresService.DeleteShippingAddressById(shippingAdressId)
	assert.Nil(t, err)
}

func TestShouldGetShippingAddressByShippingAddressId(t *testing.T) {
	shippingAdressId := 1
	firstname := "Naveen"
	lastname := "Kumar"
	city := "Banglore"
	address1 := "Address1"
	address2 := "Address2"
	countryid := 81
	postcode := 560063

	newShippingAddress := domain.NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)

	mockShippingAddressRepo.On("FindShippingAddressById", shippingAdressId).Return(newShippingAddress, nil)
	resShippingAddress, _ := shippingAddresService.GetShippingAddressById(shippingAdressId)

	assert.Equal(t, firstname, resShippingAddress.FirstName)
	assert.Equal(t, lastname, resShippingAddress.LastName)
	assert.Equal(t, city, resShippingAddress.City)
	assert.Equal(t, address1, resShippingAddress.Address1)
	assert.Equal(t, address2, resShippingAddress.Address2)
	assert.Equal(t, countryid, resShippingAddress.CountryID)
	assert.Equal(t, postcode, resShippingAddress.PostCode)
}

func TestShouldNotDeleteShippingAddressByShippingAddressIdUponInvalidShippingAddressId(t *testing.T) {
	shippingAddressId := -99
	errMessage := "some error"
	mockShippingAddressRepo.On("DeleteShippingAddressById", shippingAddressId).Return(errs.NewUnexpectedError(errMessage))

	err := shippingAddresService.DeleteShippingAddressById(shippingAddressId)
	assert.Error(t, err.Error(), errMessage)
}

func TestShouldUpdateShippingAddress(t *testing.T) {
	firstname := "Naveen"
	lastname := "Kumar"
	city := "Banglore"
	address1 := "Address1"
	address2 := "Address2"
	countryid := 99
	postcode := 560063

	newShippingAddress := domain.NewShippingAddress(firstname, lastname, city, address1, address2, countryid, postcode)
	mockShippingAddressRepo.On("UpdateShippingAddress", *newShippingAddress).Return(newShippingAddress, nil)
	updatedShippingAddress, _ := shippingAddresService.UpdateShippingAddress(*newShippingAddress)

	assert.Equal(t, newShippingAddress.FirstName, updatedShippingAddress.FirstName)
	assert.Equal(t, newShippingAddress.LastName, updatedShippingAddress.LastName)
	assert.Equal(t, newShippingAddress.City, updatedShippingAddress.City)
	assert.Equal(t, newShippingAddress.Address1, updatedShippingAddress.Address1)
	assert.Equal(t, newShippingAddress.Address2, updatedShippingAddress.Address2)
	assert.Equal(t, newShippingAddress.CountryID, updatedShippingAddress.CountryID)
	assert.Equal(t, newShippingAddress.PostCode, updatedShippingAddress.PostCode)
}
