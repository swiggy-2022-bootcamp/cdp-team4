package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNewShippingAddress(t *testing.T) {

	model := "boat bassheads 100"
	quantity := int64(100)
	price := float64(299)
	manufacturerID := "boat_company_id"
	sku := "ZG011AQA"
	productSEOURLs := []ProductSEOURL{}
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
	productDescription := []ProductDescription{}
	productCategories := []ProductCategory{}

	newProduct := NewProductObject(model, quantity, price, manufacturerID, sku, productSEOURLs, points, rewards,
		imageURL, isShippable, weight, length, width, height, minimumQuantity, relatedProducts, productDescription,
		productCategories)

	assert.NotEmpty(t, newProduct.Id)
	assert.Equal(t, model, newProduct.Model)
	assert.Equal(t, quantity, newProduct.Quantity)
	assert.Equal(t, price, newProduct.Price)
	assert.Equal(t, manufacturerID, newProduct.ManufacturerID)
	assert.Equal(t, sku, newProduct.SKU)
	assert.Equal(t, productSEOURLs, newProduct.ProductSEOURLs)
	assert.Equal(t, points, newProduct.Points)
	assert.Equal(t, rewards, newProduct.Reward)
	assert.Equal(t, imageURL, newProduct.ImageURL)
	assert.Equal(t, isShippable, newProduct.IsShippable)
	assert.Equal(t, weight, newProduct.Weight)
	assert.Equal(t, length, newProduct.Length)
	assert.Equal(t, width, newProduct.Width)
	assert.Equal(t, height, newProduct.Height)
	assert.Equal(t, minimumQuantity, newProduct.MinimumQuantity)
	assert.Equal(t, relatedProducts, newProduct.RelatedProducts)
	assert.Equal(t, productDescription, newProduct.ProductDescriptions)
	assert.Equal(t, productCategories, newProduct.ProductCategories)
}
