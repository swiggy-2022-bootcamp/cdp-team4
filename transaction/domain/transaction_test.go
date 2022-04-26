package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToReturnNewTransaction(t *testing.T) {

	userId := "12345678"
	transactionPoints := 10
	newCart := NewTransaction(userId, transactionPoints)

	assert.Equal(t, userId, newCart.UserID)
	assert.Equal(t, transactionPoints, newCart.TransactionPoints)
}
