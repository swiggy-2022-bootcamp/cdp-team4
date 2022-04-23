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
var mockOrderOverviewRepo = mocks.OrderOverviewRepository{}
var orderService = domain.NewOrderService(&mockOrderRepo)
var orderOverviewService = domain.NewOrderOverviewService(&mockOrderOverviewRepo)

func TestShouldReturnNewOrderService(t *testing.T) {
	orderService := domain.NewOrderService(nil)
	assert.NotNil(t, orderService)
}

func TestShouldCreateNewOrder(t *testing.T) {
	orderid := "1203712"
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

	_ = domain.NewOrder(userid, status, prodquant, prodcost, totalcost)

	mockOrderRepo.On("InsertOrder", mock.Anything).Return(orderid, nil)
	orderService.CreateOrder(userid, status, prodquant, prodcost, totalcost)
	mockOrderRepo.AssertNumberOfCalls(t, "InsertOrder", 1)
}

func TestShouldDeleteOrderByOrderId(t *testing.T) {
	orderId := "10293194182"
	mockOrderRepo.On("DeleteOrderById", orderId).Return(true, nil)
	res, err := orderService.DeleteOrderById(orderId)
	assert.Nil(t, err)
	assert.Equal(t, res, true)
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
	mockOrderRepo.On("DeleteOrderById", orderId).Return(false, errs.NewUnexpectedError(errMessage))

	res, err := orderService.DeleteOrderById(orderId)
	assert.Equal(t, res, false)
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

	_ = domain.NewOrder(userid, status, prodquant, prodcost, totalcost)
	mockOrderRepo.On("UpdateOrderStatus", orderid, status).Return(true, nil)
	res, err := orderService.UpdateOrderStatus(orderid, status)

	assert.Nil(t, err)
	assert.Equal(t, res, true)
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

func TestShouldGetOrderByStatus(t *testing.T) {
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

	mockOrderRepo.On("FindOrderByStatus", status).Return([]domain.Order{*newOrder}, nil)
	resOrders, _ := orderService.GetOrderByStatus(status)

	assert.Equal(t, userId, resOrders[0].UserID)
	assert.Equal(t, status, resOrders[0].Status)
	assert.Equal(t, prodquant, resOrders[0].ProductsQuantity)
	assert.Equal(t, prodcost, resOrders[0].ProductsCost)
	assert.Equal(t, totalcost, resOrders[0].TotalCost)
}

func TestShouldGetAllOrders(t *testing.T) {
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

	mockOrderRepo.On("FindAllOrders").Return([]domain.Order{*newOrder}, nil)
	resOrders, _ := orderService.GetAllOrders()

	assert.Equal(t, userId, resOrders[0].UserID)
	assert.Equal(t, status, resOrders[0].Status)
	assert.Equal(t, prodquant, resOrders[0].ProductsQuantity)
	assert.Equal(t, prodcost, resOrders[0].ProductsCost)
	assert.Equal(t, totalcost, resOrders[0].TotalCost)
}

func TestShouldCreateOrderOverview(t *testing.T) {
	orderId := "12321324"
	products_map := map[string]int{
		"2134123": 10,
		"2314343": 11,
	}
	neworderoverview := domain.OrderOverview{
		OrderID:            orderId,
		ProductsIdQuantity: products_map,
	}
	mockOrderOverviewRepo.On("InsertOrderOverview", mock.Anything).Return(true, nil)

	resOrder, resnil := orderOverviewService.CreateOrderOverview(neworderoverview)

	assert.Equal(t, resOrder, true)
	assert.Nil(t, resnil)

}

func TestShouldGetOrderOverview(t *testing.T) {
	orderId := "12321324"
	products_map := map[string]int{
		"2134123": 10,
		"2314343": 11,
	}

	neworderoverview := domain.OrderOverview{
		OrderID:            orderId,
		ProductsIdQuantity: products_map,
	}

	mockOrderOverviewRepo.On("GetOrderOverview", orderId).Return(&neworderoverview, nil)
	_, resnil := orderOverviewService.GetOrderOverviewByOrderID(orderId)
	assert.Nil(t, resnil)

}
