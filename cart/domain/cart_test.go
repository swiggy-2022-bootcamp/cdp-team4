package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToReturnNewCart(t *testing.T) {

	userid := "12345678"
	prodquant := map[string]int{
		"Alchemist":  1,
		"Sapiens": 2,
	}
	prodcost := map[string]int{
		"Alchemist": 399,
		"Sapiens": 500,
	}
	newCart := NewCart(userid, prodquant,prodcost)

	assert.Equal(t, userid, newCart.UserID)
	assert.Equal(t, prodquant, newCart.ProductsQuantity)
}
