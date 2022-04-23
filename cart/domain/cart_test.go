package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToReturnNewCart(t *testing.T) {

	userid := "12321324"
	prodquant := map[string]int{
		"Origin of life":  1,
		"Reynolds trimax": 10,
	}

	newCart := NewCart(userid, prodquant)

	assert.Equal(t, userid, newCart.UserID)
	assert.Equal(t, prodquant, newCart.ProductsQuantity)
}
