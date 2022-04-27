package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToReturnNewCart(t *testing.T) {

	userid := "12345678"
	products := (map[string]Item{})
	products["1"] = Item{Name: "pen", Cost: 10, Quantity: 1}
	products["2"] = Item{Name: "pencil", Cost: 5, Quantity: 2}
	newCart := NewCart(userid, products)

	assert.Equal(t, userid, newCart.UserID)
	assert.Equal(t, products, newCart.Items)
	assert.Equal(t, products["1"], newCart.Items["1"])
}
