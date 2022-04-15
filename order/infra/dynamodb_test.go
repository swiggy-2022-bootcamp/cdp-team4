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
		"Reynolds trimax": 70,
	}
	totalcost := 1999

	newOrder := domain.NewOrder(userid, status, prodquant, prodcost, totalcost)
	res, err := testOrderService.InsertOrder(*newOrder)
	insertedid = res
	t.Logf("Inserted Id is %s\n", insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldCreateNewOrder2inDynamoDB(t *testing.T) {
	userid := uuid.New().String()
	status := "confirmed"
	prodquant := map[string]int{
		"Origin of life":  1,
		"Reynolds trimax": 10,
	}
	prodcost := map[string]int{
		"Origin of life":  999,
		"Reynolds trimax": 70,
	}
	totalcost := 1999

	newOrder := domain.NewOrder(userid, status, prodquant, prodcost, totalcost)
	res, err := testOrderService.InsertOrder(*newOrder)
	t.Logf("Inserted Id is %s\n", insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}
func TestShouldGetOrderByOrderIdDynamoDB(t *testing.T) {
	status := "confirmed"
	prodquant := map[string]int{
		"Origin of life":  1,
		"Reynolds trimax": 10,
	}
	prodcost := map[string]int{
		"Origin of life":  999,
		"Reynolds trimax": 70,
	}
	totalcost := 1999
	t.Logf("Inserted Id is %s Reading\n", insertedid)
	res, err := testOrderService.FindOrderById(insertedid)
	t.Logf("Read %v", res)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, res.Status, status)
	assert.Equal(t, res.ProductsQuantity, prodquant)
	assert.Equal(t, res.ProductsCost, prodcost)
	assert.Equal(t, res.TotalCost, totalcost)
}

func TestShouldGetOrderByUserIdDynamoDB(t *testing.T) {
	status := "confirmed"
	prodquant := map[string]int{
		"Origin of life":  1,
		"Reynolds trimax": 10,
	}
	prodcost := map[string]int{
		"Origin of life":  999,
		"Reynolds trimax": 70,
	}
	totalcost := 1999
	t.Logf("Inserted User Id is %s Reading\n", inserteduserid)
	res, err := testOrderService.FindOrderByUserId(inserteduserid)
	t.Logf("Read %v", res)
	t.Logf("Length %d", len(res))
	t.Logf("Read item %v", res[len(res)-1])
	t.Logf("Read item status %v", res[len(res)-1].Status)
	t.Logf("Read item cost %v", res[len(res)-1].ProductsCost)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, res[0].Status, status)
	assert.Equal(t, res[0].ProductsQuantity, prodquant)
	assert.Equal(t, res[0].ProductsCost, prodcost)
	assert.Equal(t, res[0].TotalCost, totalcost)
}

func TestShouldUpdateOrderStatusDynamoDB(t *testing.T) {
	res, err := testOrderService.UpdateOrderStatus(insertedid, "declined")
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestShouldDeleteOrderByOrderIdDynamoDB(t *testing.T) {
	res, err := testOrderService.DeleteOrderById(insertedid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}
