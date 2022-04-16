package infra_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/infra"
)

var testShippingCostService = infra.NewShippingCostDynamoRepository()

func TestShouldCreateNewShippingCostinDynamoDB(t *testing.T) {

	city := "Banglore"
	cost := 10

	newShippingCost := domain.NewShippingCost(city, cost)
	res, err := testShippingCostService.InsertShippingCost(*newShippingCost)
	t.Logf("Inserted Id is %s\n", insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldCreateNewShippingCost2inDynamoDB(t *testing.T) {

	city := "Chennai"
	cost := 109

	newShippingCost := domain.NewShippingCost(city, cost)
	res, err := testShippingCostService.InsertShippingCost(*newShippingCost)
	t.Logf("Inserted Id is %s\n", insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldGetShippingCostsByShippingCostinDynamoDB(t *testing.T) {
	t.Logf("Inserted Id is %s Reading\n", insertedid)
	res, err := testShippingCostService.FindShippingCostByCity("Banglore")
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldUpdateShippingCostinDynamoDB(t *testing.T) {

	city := "Banglore"
	cost := 19

	newShippingCost := domain.NewShippingCost(city, cost)
	res, err := testShippingCostService.UpdateShippingCost(*newShippingCost)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldDeleteShippingCostsByShippingCostsIdDynamoDB(t *testing.T) {
	res, err := testShippingCostService.DeleteShippingCostByCity("Banglore")
	assert.NotNil(t, res)
	assert.Nil(t, err)
}
