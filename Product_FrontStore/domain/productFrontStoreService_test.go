package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/domain"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/mocks"
)

var mockProductFrontStoreRepo = mocks.ProductFrontStoreDynamoRepository{}
var productFrontStoreService = domain.NewProductFrontStoreService(&mockProductFrontStoreRepo)

func TestShouldReturnNewCategoryService(t *testing.T) {
	productFrontStoreService := domain.NewProductFrontStoreService(nil)
	assert.NotNil(t, productFrontStoreService)
}

func TestGetProductByProductId(t *testing.T) {
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

	mockProductFrontStoreRepo.On("FindByProductID", productId).Return(*newProduct, nil)
	resProduct, _ := productFrontStoreService.GetProductById(productId)
	mockProductFrontStoreRepo.AssertNumberOfCalls(t, "FindByProductID", 1)

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

func TestGetAllPeoducts(t *testing.T) {
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

	mockProductFrontStoreRepo.On("Find").Return([]domain.Product{*newProduct}, nil)
	resProducts, _ := productFrontStoreService.GetProducts()

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

func TestGetProductsByCategoryId(t *testing.T) {
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
	productCategories := []string{"ZG011AQA", "ZG011AXZ"}

	newProduct := domain.NewProductObject(model, quantity, price, manufacturerID, sku, productSEOURLs, points, rewards,
		imageURL, isShippable, weight, length, width, height, minimumQuantity, relatedProducts, productDescription,
		productCategories)
	mockProductFrontStoreRepo.On("FindByCategoryID", productCategories[0]).Return([]domain.Product{*newProduct}, nil)
	resProducts, err := productFrontStoreService.GetProductsByCategoryId(productCategories[0])
	assert.Nil(t, err)
	assert.NotEmpty(t, resProducts)
}
