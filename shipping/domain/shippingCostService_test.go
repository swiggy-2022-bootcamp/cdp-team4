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

	newShippingCost := domain.NewShippingCost(city, cost)

	mockShippingCostRepo.On("InsertShippingCost", mock.Anything).Return(*newShippingCost, nil)
	shippingCostService.CreateShippingCost(city, cost)
	mockShippingCostRepo.AssertNumberOfCalls(t, "InsertShippingCost", 1)
}

func TestShouldDeleteShippingCostByShippingCostId(t *testing.T) {
	shippingCostId := 1
	mockShippingCostRepo.On("DeleteShippingCostById", shippingCostId).Return(nil)
	var err = shippingCostService.DeleteShippingCostById(shippingCostId)
	assert.Nil(t, err)
}

func TestShouldGetShippingCostByShippingCostId(t *testing.T) {
	shippingCostId := 1
	city := "Banglore"
	cost := 199

	newShippingCost := domain.NewShippingCost(city, cost)

	mockShippingCostRepo.On("FindShippingCostById", shippingCostId).Return(newShippingCost, nil)
	resShippingCost, _ := shippingCostService.GetShippingCostById(shippingCostId)

	assert.Equal(t, city, resShippingCost.City)
	assert.Equal(t, cost, resShippingCost.ShippingCost)
}

func TestShouldNotDeleteShippingCostByShippingCostIdUponInvalidShippingCostId(t *testing.T) {
	shippingCostId := -99
	errMessage := "some error"
	mockShippingCostRepo.On("DeleteShippingCostById", shippingCostId).Return(errs.NewUnexpectedError(errMessage))

	err := shippingCostService.DeleteShippingCostById(shippingCostId)
	assert.Error(t, err.Error(), errMessage)
}

func TestShouldUpdateShippingCost(t *testing.T) {
	city := "Banglore"
	cost := 290

	newShippingCost := domain.NewShippingCost(city, cost)
	mockShippingCostRepo.On("UpdateShippingCost", *newShippingCost).Return(newShippingCost, nil)
	updatedShippingCost, _ := shippingCostService.UpdateShippingCost(*newShippingCost)

	assert.Equal(t, newShippingCost.City, updatedShippingCost.City)
	assert.Equal(t, newShippingCost.ShippingCost, updatedShippingCost.ShippingCost)
}
