package domain

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/mocks"
// )

// var mockProductAdminRepo = mocks.ProductAdminRepository{}
// var orderService = NewProductAdminService(&mockProductAdminRepo)

// func TestShouldReturnNewProductAdminService(t *testing.T) {
// 	productAdminService := NewProductAdminService(nil)
// 	assert.NotNil(t, productAdminService)
// }

// func TestShouldCreateNewProduct(t *testing.T) {
// 	productid := "1203712"
// 	model := "boat bassheads 100"
// 	quantity := int64(100)
// 	price := float64(299)
// 	manufacturerID := "boat_company_id"
// 	sku := "ZG011AQA"
// 	productSEOURLs := []ProductSEOURL{}
// 	points := int64(15)
// 	rewards := int64(20)
// 	imageURL := "http://usnews.com/vivamus/metus/arcu/adipiscing.json?luctus=placerat&ultricies=praesent"
// 	isShippable := true
// 	weight := 120.25
// 	length := 0.0
// 	width := 0.0
// 	height := 30.5
// 	minimumQuantity := int64(2)
// 	relatedProducts := []string{"boat bassheads 102", "boat bassheads 152"}
// 	productDescription := []ProductDescription{}
// 	productCategories := []ProductCategory{}

// 	_ = NewProductObject(model, quantity, price, manufacturerID, sku, productSEOURLs, points, rewards,
// 		imageURL, isShippable, weight, length, width, height, minimumQuantity, relatedProducts, productDescription,
// 		productCategories)

// 	mockProductAdminRepo.On("InsertProduct", mock.Anything).Return(productid, nil)
// 	// productAdminService.CreateDynamoProductAdminRecord(model, quantity, price, manufacturerID, sku, productSEOURLs,
// 	// 	points, rewards, imageURL, isShippable, weight, length, width, height, minimumQuantity,
// 	// 	relatedProducts, productDescription, productCategories)

// 	mockProductAdminRepo.AssertNumberOfCalls(t, "InsertProduct", 1)
// }

// // func TestShouldDeleteOrderByOrderId(t *testing.T) {
// // 	orderId := "10293194182"
// // 	mockOrderRepo.On("DeleteOrderById", orderId).Return(true, nil)
// // 	res, err := orderService.DeleteOrderById(orderId)
// // 	assert.Nil(t, err)
// // 	assert.Equal(t, res, true)
// // }

// // func TestShouldGetOrderByOrderId(t *testing.T) {
// // 	orderId := "128132121"
// // 	userid := "12321322"
// // 	status := "pending"
// // 	prodquant := map[string]int{
// // 		"Origin of life":  1,
// // 		"Reynolds trimax": 10,
// // 	}
// // 	prodcost := map[string]int{
// // 		"Origin of life":  999,
// // 		"Reynolds trimax": 60,
// // 	}
// // 	totalcost := 1700

// // 	newOrder := domain.NewOrder(userid, status, prodquant, prodcost, totalcost)

// // 	mockOrderRepo.On("FindOrderById", orderId).Return(newOrder, nil)
// // 	resOrder, _ := orderService.GetOrderById(orderId)

// // 	assert.Equal(t, userid, resOrder.UserID)
// // 	assert.Equal(t, status, resOrder.Status)
// // 	assert.Equal(t, prodquant, resOrder.ProductsQuantity)
// // 	assert.Equal(t, prodcost, resOrder.ProductsCost)
// // 	assert.Equal(t, totalcost, resOrder.TotalCost)
// // }

// // func TestShouldNotDeleteOrderByOrderIdUponInvalidOrderId(t *testing.T) {
// // 	orderId := "-99"
// // 	errMessage := "some error"
// // 	mockOrderRepo.On("DeleteOrderById", orderId).Return(false, errs.NewUnexpectedError(errMessage))

// // 	res, err := orderService.DeleteOrderById(orderId)
// // 	assert.Equal(t, res, false)
// // 	assert.Error(t, err.Error(), errMessage)
// // }

// // func TestShouldUpdateOrder(t *testing.T) {
// // 	orderid := "31490934"
// // 	userid := "12321324"
// // 	status := "confirmed"
// // 	prodquant := map[string]int{
// // 		"Origin of life":  1,
// // 		"Reynolds trimax": 10,
// // 	}
// // 	prodcost := map[string]int{
// // 		"Origin of life":  1021,
// // 		"Reynolds trimax": 60,
// // 	}
// // 	totalcost := 1800

// // 	_ = domain.NewOrder(userid, status, prodquant, prodcost, totalcost)
// // 	mockOrderRepo.On("UpdateOrderStatus", orderid, status).Return(true, nil)
// // 	res, err := orderService.UpdateOrderStatus(orderid, status)

// // 	assert.Nil(t, err)
// // 	assert.Equal(t, res, true)
// // }

// // func TestShouldGetAllOrders(t *testing.T) {
// // 	userId := "12321324"
// // 	status := "pending"
// // 	prodquant := map[string]int{
// // 		"Origin of life":  1,
// // 		"Reynolds trimax": 10,
// // 	}
// // 	prodcost := map[string]int{
// // 		"Origin of life":  999,
// // 		"Reynolds trimax": 60,
// // 	}
// // 	totalcost := 1700

// // 	newOrder := domain.NewOrder(userId, status, prodquant, prodcost, totalcost)

// // 	mockOrderRepo.On("FindAllOrders").Return([]domain.Order{*newOrder}, nil)
// // 	resOrders, _ := orderService.GetAllOrders()

// // 	assert.Equal(t, userId, resOrders[0].UserID)
// // 	assert.Equal(t, status, resOrders[0].Status)
// // 	assert.Equal(t, prodquant, resOrders[0].ProductsQuantity)
// // 	assert.Equal(t, prodcost, resOrders[0].ProductsCost)
// // 	assert.Equal(t, totalcost, resOrders[0].TotalCost)
// // }
