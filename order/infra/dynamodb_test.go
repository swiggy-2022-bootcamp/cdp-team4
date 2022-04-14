package infra_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra"
)

var testOrderService = infra.NewDynamoRepository()
var insertedid string
var inserteduserid string

func TestShouldCreateNewOrderinDynamoDB(t *testing.T) {
	userid := uuid.New().String()
	inserteduserid = userid
	status := "confirmed"
	prodquant := map[string]int{
		"Origin of life":  1,
		"Reynolds trimax": 10,
	}
	prodcost := map[string]int{
		"Origin of life":  999,
		"Reynolds trimax": 60,
	}
	totalcost := 1700

	newOrder := domain.NewOrder(userid, status, prodquant, prodcost, totalcost)
	res, err := testOrderService.InsertOrder(*newOrder)
	insertedid = res
	t.Logf("Inserted Id is %s\n", insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldGetOrderByOrderIdDynamoDB(t *testing.T) {
	t.Logf("Inserted Id is %s Reading\n", insertedid)
	res, err := testOrderService.FindOrderById(insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldGetOrderByUserIdDynamoDB(t *testing.T) {
	t.Logf("Inserted User Id is %s Reading\n", inserteduserid)
	res, err := testOrderService.FindOrderByUserId(inserteduserid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldUpdateOrderStatusDynamoDB(t *testing.T) {
	res, err := testOrderService.UpdateOrderStatus(insertedid, "declined")
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldDeleteOrderByOrderIdDynamoDB(t *testing.T) {
	res, err := testOrderService.DeleteByID(insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}
