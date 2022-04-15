package domain_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

	"github.com/stretchr/testify/assert"
)

var mockShippingCostRepo = mocks.ShippingCostRepository{}
var shippingCostService = domain.NewShippingCostService(&mockShippingCostRepo)

func TestShouldReturnNewShippingCostService(t *testing.T) {
	userService := domain.NewShippingCostService(nil)
	assert.NotNil(t, userService)
}

func TestShouldCreateNewShippingCost(t *testing.T) {

	city := "Banglore"
	cost := 199

	_ = domain.NewShippingCost(city, cost)

	mockShippingCostRepo.On("InsertShippingCost", mock.Anything).Return(true, nil)
	shippingCostService.CreateShippingCost(city, cost)
	mockShippingCostRepo.AssertNumberOfCalls(t, "InsertShippingCost", 1)
}

func TestShouldDeleteShippingCostByCity(t *testing.T) {
	city := "Banglore"
	mockShippingCostRepo.On("DeleteShippingCostByCity", city).Return(true, nil)
	res, err := shippingCostService.DeleteShippingCostByCity(city)
	assert.Equal(t, res, true)
	assert.Nil(t, err)
}

func TestShouldGetShippingCostByCity(t *testing.T) {
	city := "Banglore"
	cost := 199

	newShippingCost := domain.NewShippingCost(city, cost)

	mockShippingCostRepo.On("FindShippingCostByCity", city).Return(newShippingCost, nil)
	resShippingCost, _ := shippingCostService.GetShippingCostByCity(city)

	assert.Equal(t, city, resShippingCost.City)
	assert.Equal(t, cost, resShippingCost.ShippingCost)
}

func TestShouldNotDeleteShippingCostByCityUponInvalidCity(t *testing.T) {
	non_existing_city := "fsdjf"
	errMessage := "some error"
	mockShippingCostRepo.On("DeleteShippingCostByCity", non_existing_city).Return(false, errs.NewUnexpectedError(errMessage))

	res, err := shippingCostService.DeleteShippingCostByCity(non_existing_city)
	assert.Equal(t, res, false)
	assert.Error(t, err.Error(), errMessage)
}

func TestShouldUpdateShippingCost(t *testing.T) {
	city := "Banglore"
	cost := 290

	newShippingCost := domain.NewShippingCost(city, cost)
	mockShippingCostRepo.On("UpdateShippingCost", *newShippingCost).Return(true, nil)
	res, err := shippingCostService.UpdateShippingCost(*newShippingCost)

	assert.Equal(t, res, true)
	assert.Nil(t, err)
}
