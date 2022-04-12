package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNewShippingCost(t *testing.T) {
	city := "Banglore"
	cost := 199

	newShippingCost := NewShippingCost(city, cost)

	assert.Equal(t, city, newShippingCost.City)
	assert.Equal(t, cost, newShippingCost.ShippingCost)
}
