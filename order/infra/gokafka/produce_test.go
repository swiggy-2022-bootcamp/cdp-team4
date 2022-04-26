package gokafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduceKafkaPaymentStatus(t *testing.T) {
	newpayment := PaymentRecord{
		Amount:   15,
		Currency: "INR",
		UserID:   "b533ab38-97f4-4c43-a891-c43872d9f15a",
		OrderID:  "fe476221-054d-4caa-805b-27b32a3fc82b",
		Status:   "Confirmed",
		Notes:    []string{"string", "string2"},
	}
	res, err := WriteMsgToKafka("payment", newpayment)
	assert.Equal(t, res, true)
	assert.Nil(t, err)
}
