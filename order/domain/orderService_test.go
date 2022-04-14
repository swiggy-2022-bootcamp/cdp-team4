package domain_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/utils/errs"

	"github.com/stretchr/testify/assert"
)

var mockOrderRepo = mocks.OrderRepository{}
var orderService = domain.NewOrderService(&mockOrderRepo)

func TestShouldReturnNewOrderService(t *testing.T) {
	orderService := domain.NewOrderService(nil)
	assert.NotNil(t, orderService)
}

func TestShouldCreateNewOrder(t *testing.T) {

	userid := "12321324"
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

	newOrder := domain.NewOrder(userid, status, prodquant, prodcost, totalcost)

	mockOrderRepo.On("InsertOrder", mock.Anything).Return(*newOrder, nil)
	orderService.CreateOrder(userid, status, prodquant, prodcost, totalcost)
	mockOrderRepo.AssertNumberOfCalls(t, "InsertOrder", 1)
}

func TestShouldDeleteOrderByOrderId(t *testing.T) {
	orderId := "10293194182"
	mockOrderRepo.On("DeleteOrderById", orderId).Return(nil)
	var err = orderService.DeleteOrderById(orderId)
	assert.Nil(t, err)
}

func TestShouldGetOrderByOrderId(t *testing.T) {
	orderId := "128132121"
	userid := "12321322"
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

	newOrder := domain.NewOrder(userid, status, prodquant, prodcost, totalcost)

	mockOrderRepo.On("FindOrderById", orderId).Return(newOrder, nil)
	resOrder, _ := orderService.GetOrderById(orderId)

	assert.Equal(t, userid, resOrder.UserID)
	assert.Equal(t, status, resOrder.Status)
	assert.Equal(t, prodquant, resOrder.ProductsQuantity)
	assert.Equal(t, prodcost, resOrder.ProductsCost)
	assert.Equal(t, totalcost, resOrder.TotalCost)
}

func TestShouldNotDeleteOrderByOrderIdUponInvalidOrderId(t *testing.T) {
	orderId := "-99"
	errMessage := "some error"
	mockOrderRepo.On("DeleteOrderById", orderId).Return(errs.NewUnexpectedError(errMessage))

	err := orderService.DeleteOrderById(orderId)
	assert.Error(t, err.Error(), errMessage)
}

func TestShouldUpdateOrder(t *testing.T) {
	orderid := "31490934"
	userid := "12321324"
	status := "confirmed"
	prodquant := map[string]int{
		"Origin of life":  1,
		"Reynolds trimax": 10,
	}
	prodcost := map[string]int{
		"Origin of life":  1021,
		"Reynolds trimax": 60,
	}
	totalcost := 1800

	newOrder := domain.NewOrder(userid, status, prodquant, prodcost, totalcost)
	mockOrderRepo.On("UpdateOrderStatus", orderid, status).Return(newOrder, nil)
	resOrder, _ := orderService.UpdateOrderStatus(orderid, status)

	assert.Equal(t, userid, resOrder.UserID)
	assert.Equal(t, status, resOrder.Status)
	assert.Equal(t, prodquant, resOrder.ProductsQuantity)
	assert.Equal(t, prodcost, resOrder.ProductsCost)
	assert.Equal(t, totalcost, resOrder.TotalCost)
}

func TestShouldGetOrderByUserId(t *testing.T) {
	userId := "12321324"
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

	newOrder := domain.NewOrder(userId, status, prodquant, prodcost, totalcost)

	mockOrderRepo.On("FindOrderByUserId", userId).Return([]domain.Order{*newOrder}, nil)
	resOrders, _ := orderService.GetOrderByUserId(userId)

	assert.Equal(t, userId, resOrders[0].UserID)
	assert.Equal(t, status, resOrders[0].Status)
	assert.Equal(t, prodquant, resOrders[0].ProductsQuantity)
	assert.Equal(t, prodcost, resOrders[0].ProductsCost)
	assert.Equal(t, totalcost, resOrders[0].TotalCost)
}
