package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNewShippingAddress(t *testing.T) {

	userid := 12321324
	status := "pending"
	prodquant := map[string]int{
		"Origin of life":  1,
		"Reynolds trimax": 10,
	}
	prodcost := map[string]int{
		"Origin of life":  999,
		"Reynolds trimax": 60,
	}
	totalcost := 1700

	newOrder := NewOrder(userid, status, prodquant, prodcost, totalcost)

	assert.Equal(t, userid, newOrder.UserID)
	assert.Equal(t, status, newOrder.Status)
	assert.Equal(t, prodquant, newOrder.ProductsQuantity)
	assert.Equal(t, prodcost, newOrder.ProductsCost)
	assert.Equal(t, totalcost, newOrder.TotalCost)
}
