package domain_test

import (
	// "errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/domain"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/mocks"
)

var mockProductAdminRepo = mocks.ProductAdminDynamoRepository{}
var productService = domain.NewProductAdminService(&mockProductAdminRepo)

func TestShouldReturnNewProductAdminService(t *testing.T) {
	productAdminService := domain.NewProductAdminService(nil)
	assert.NotNil(t, productAdminService)
}

func TestShouldCreateNewProduct(t *testing.T) {
	model := "boat bassheads 100"
	quantity := int64(100)
	price := float64(299)
	manufacturerID := "boat_company_id"
	sku := "ZG011AQA"
	productSEOURLs := []domain.ProductSEOURL{}
	points := int64(15)
	rewards := int64(20)
	imageURL := "http://usnews.com/vivamus/metus/arcu/adipiscing.json?luctus=placerat&ultricies=praesent"
	isShippable := true
	weight := 120.25
	length := 0.0
	width := 0.0
	height := 30.5
	minimumQuantity := int64(2)
	relatedProducts := []string{"boat bassheads 102", "boat bassheads 152"}
	productDescription := []domain.ProductDescription{}
	productCategories := []string{}

	_ = domain.NewProductObject(model, quantity, price, manufacturerID, sku, productSEOURLs, points, rewards,
		imageURL, isShippable, weight, length, width, height, minimumQuantity, relatedProducts, productDescription,
		productCategories)

	mockProductAdminRepo.On("Insert", mock.Anything).Return(true, nil)
	productService.CreateDynamoProductAdminRecord(model, quantity, price, manufacturerID, sku, productSEOURLs,
		points, rewards, imageURL, isShippable, weight, length, width, height, minimumQuantity,
		relatedProducts, productDescription, productCategories)

	mockProductAdminRepo.AssertNumberOfCalls(t, "Insert", 1)
}

func TestShouldDeleteProductByProductId(t *testing.T) {
	productId := "10293194182"
	mockProductAdminRepo.On("DeleteByID", productId).Return(true, nil)
	res, err := productService.DeleteProductById(productId)
	assert.Nil(t, err)
	assert.Equal(t, res, true)
}

func TestShouldGetProductByProductId(t *testing.T) {
	productId := "1203712"
	model := "boat bassheads 100"
	quantity := int64(100)
	price := float64(299)
	manufacturerID := "boat_company_id"
	sku := "ZG011AQA"
	productSEOURLs := []domain.ProductSEOURL{}
	points := int64(15)
	rewards := int64(20)
	imageURL := "http://usnews.com/vivamus/metus/arcu/adipiscing.json?luctus=placerat&ultricies=praesent"
	isShippable := true
	weight := 120.25
	length := 0.0
	width := 0.0
	height := 30.5
	minimumQuantity := int64(2)
	relatedProducts := []string{"boat bassheads 102", "boat bassheads 152"}
	productDescription := []domain.ProductDescription{}
	productCategories := []string{}

	newProduct := domain.NewProductObject(model, quantity, price, manufacturerID, sku, productSEOURLs, points, rewards,
		imageURL, isShippable, weight, length, width, height, minimumQuantity, relatedProducts, productDescription,
		productCategories)

	mockProductAdminRepo.On("FindByID", productId).Return(*newProduct, nil)
	resProduct, _ := productService.GetProductById(productId)
	mockProductAdminRepo.AssertNumberOfCalls(t, "FindByID", 1)

	assert.Equal(t, model, resProduct.Model)
	assert.Equal(t, quantity, resProduct.Quantity)
	assert.Equal(t, price, resProduct.Price)
	assert.Equal(t, manufacturerID, resProduct.ManufacturerID)
	assert.Equal(t, sku, resProduct.SKU)
	assert.Equal(t, productSEOURLs, resProduct.ProductSEOURLs)
	assert.Equal(t, points, resProduct.Points)
	assert.Equal(t, rewards, resProduct.Reward)
	assert.Equal(t, imageURL, resProduct.ImageURL)
	assert.Equal(t, isShippable, resProduct.IsShippable)
	assert.Equal(t, weight, resProduct.Weight)
	assert.Equal(t, height, resProduct.Height)
	assert.Equal(t, minimumQuantity, resProduct.MinimumQuantity)
	assert.Equal(t, relatedProducts, resProduct.RelatedProducts)
	assert.Equal(t, productDescription, resProduct.ProductDescriptions)
	assert.Equal(t, productCategories, resProduct.ProductCategories)
}

func TestShouldNotDeleteProductByProductIdUponInvalidProductId(t *testing.T) {
	productId := "-99"
	errMessage := "some error"
	mockProductAdminRepo.On("DeleteByID", productId).Return(false, errors.New(errMessage))

	res, err := productService.DeleteProductById(productId)
	assert.Equal(t, res, false)
	assert.Error(t, err, errMessage)
}

func TestShouldUpdateProductDetails(t *testing.T) {
	model := "boat bassheads 100"
	quantity := int64(100)
	price := float64(299)
	manufacturerID := "boat_company_id"
	sku := "ZG011AQA"
	productSEOURLs := []domain.ProductSEOURL{}
	points := int64(15)
	rewards := int64(20)
	imageURL := "http://usnews.com/vivamus/metus/arcu/adipiscing.json?luctus=placerat&ultricies=praesent"
	isShippable := true
	weight := 120.25
	length := 0.0
	width := 0.0
	height := 30.5
	minimumQuantity := int64(2)
	relatedProducts := []string{"boat bassheads 102", "boat bassheads 152"}
	productDescription := []domain.ProductDescription{}
	productCategories := []string{}

	newProduct := domain.NewProductObject(model, quantity, price, manufacturerID, sku, productSEOURLs, points, rewards,
		imageURL, isShippable, weight, length, width, height, minimumQuantity, relatedProducts, productDescription,
		productCategories)

	mockProductAdminRepo.On("UpdateItem", *newProduct).Return(true, nil)
	res, err := productService.UpdateProduct(*newProduct)

	assert.Nil(t, err)
	assert.Equal(t, res, true)
}

func TestShouldGetAllPeoducts(t *testing.T) {
	model := "boat bassheads 100"
	quantity := int64(100)
	price := float64(299)
	manufacturerID := "boat_company_id"
	sku := "ZG011AQA"
	productSEOURLs := []domain.ProductSEOURL{}
	points := int64(15)
	rewards := int64(20)
	imageURL := "http://usnews.com/vivamus/metus/arcu/adipiscing.json?luctus=placerat&ultricies=praesent"
	isShippable := true
	weight := 120.25
	length := 0.0
	width := 0.0
	height := 30.5
	minimumQuantity := int64(2)
	relatedProducts := []string{"boat bassheads 102", "boat bassheads 152"}
	productDescription := []domain.ProductDescription{}
	productCategories := []string{}

	newProduct := domain.NewProductObject(model, quantity, price, manufacturerID, sku, productSEOURLs, points, rewards,
		imageURL, isShippable, weight, length, width, height, minimumQuantity, relatedProducts, productDescription,
		productCategories)

	mockProductAdminRepo.On("Find").Return([]domain.Product{*newProduct}, nil)
	resProducts, _ := productService.GetProduct()

	assert.Equal(t, model, resProducts[0].Model)
	assert.Equal(t, quantity, resProducts[0].Quantity)
	assert.Equal(t, price, resProducts[0].Price)
	assert.Equal(t, manufacturerID, resProducts[0].ManufacturerID)
	assert.Equal(t, sku, resProducts[0].SKU)
	assert.Equal(t, productSEOURLs, resProducts[0].ProductSEOURLs)
	assert.Equal(t, points, resProducts[0].Points)
	assert.Equal(t, rewards, resProducts[0].Reward)
	assert.Equal(t, imageURL, resProducts[0].ImageURL)
	assert.Equal(t, isShippable, resProducts[0].IsShippable)
	assert.Equal(t, weight, resProducts[0].Weight)
	assert.Equal(t, height, resProducts[0].Height)
	assert.Equal(t, minimumQuantity, resProducts[0].MinimumQuantity)
	assert.Equal(t, relatedProducts, resProducts[0].RelatedProducts)
	assert.Equal(t, productDescription, resProducts[0].ProductDescriptions)
	assert.Equal(t, productCategories, resProducts[0].ProductCategories)
}

func TestShouldGetProductAvailability(t *testing.T) {
	model := "boat bassheads 100"
	quantity := int64(100)
	price := float64(299)
	manufacturerID := "boat_company_id"
	sku := "ZG011AQA"
	productSEOURLs := []domain.ProductSEOURL{}
	points := int64(15)
	rewards := int64(20)
	imageURL := "http://usnews.com/vivamus/metus/arcu/adipiscing.json?luctus=placerat&ultricies=praesent"
	isShippable := true
	weight := 120.25
	length := 0.0
	width := 0.0
	height := 30.5
	minimumQuantity := int64(2)
	relatedProducts := []string{"boat bassheads 102", "boat bassheads 152"}
	productDescription := []domain.ProductDescription{}
	productCategories := []string{}

	newProduct := domain.NewProductObject(model, quantity, price, manufacturerID, sku, productSEOURLs, points, rewards,
		imageURL, isShippable, weight, length, width, height, minimumQuantity, relatedProducts, productDescription,
		productCategories)
	mockProductAdminRepo.On("GetProductAvailability", newProduct.Id, int64(10)).Return(true, nil)
	res, err := productService.GetProductAvailability(newProduct.Id, int64(10))

	assert.Nil(t, err)
	assert.Equal(t, res, true)
}

func TestShouldGetProductByManufacturereId(t *testing.T) {
	model := "boat bassheads 100"
	quantity := int64(100)
	price := float64(299)
	manufacturerID := "boat_company_id"
	sku := "ZG011AQA"
	productSEOURLs := []domain.ProductSEOURL{}
	points := int64(15)
	rewards := int64(20)
	imageURL := "http://usnews.com/vivamus/metus/arcu/adipiscing.json?luctus=placerat&ultricies=praesent"
	isShippable := true
	weight := 120.25
	length := 0.0
	width := 0.0
	height := 30.5
	minimumQuantity := int64(2)
	relatedProducts := []string{"boat bassheads 102", "boat bassheads 152"}
	productDescription := []domain.ProductDescription{}
	productCategories := []string{}

	newProduct := domain.NewProductObject(model, quantity, price, manufacturerID, sku, productSEOURLs, points, rewards,
		imageURL, isShippable, weight, length, width, height, minimumQuantity, relatedProducts, productDescription,
		productCategories)
	products := []domain.Product{*newProduct}
	mockProductAdminRepo.On("FindByManufacturerID", newProduct.ManufacturerID).Return(products, nil)
	res, err := productService.GetProductByManufacturerId(manufacturerID)

	assert.Nil(t, err)
	assert.Equal(t, manufacturerID, res[0].ManufacturerID)
}

func TestShouldGetProductByKeyword(t *testing.T) {
	model := "boat bassheads 100"
	quantity := int64(100)
	price := float64(299)
	manufacturerID := "boat_company_id"
	sku := "ZG011AQA"
	productSEOURLs := []domain.ProductSEOURL{}
	points := int64(15)
	rewards := int64(20)
	imageURL := "http://usnews.com/vivamus/metus/arcu/adipiscing.json?luctus=placerat&ultricies=praesent"
	isShippable := true
	weight := 120.25
	length := 0.0
	width := 0.0
	height := 30.5
	minimumQuantity := int64(2)
	relatedProducts := []string{"boat bassheads 102", "boat bassheads 152"}
	productDescription := []domain.ProductDescription{}
	productCategories := []string{}

	newProduct := domain.NewProductObject(model, quantity, price, manufacturerID, sku, productSEOURLs, points, rewards,
		imageURL, isShippable, weight, length, width, height, minimumQuantity, relatedProducts, productDescription,
		productCategories)
	products := []domain.Product{*newProduct}
	mockProductAdminRepo.On("FindByKeyword", model).Return(products, nil)
	res, err := productService.GetProductByKeyword(model)

	assert.Nil(t, err)
	assert.Equal(t, model, res[0].Model)
}
