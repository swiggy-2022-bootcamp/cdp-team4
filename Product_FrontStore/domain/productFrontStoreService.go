package domain

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=ProductFrontStoreService
type ProductFrontStoreService interface {
	GetProducts() ([]Product, error)
	GetProductById(string) (Product, error)
	GetProductsByCategoryId(string) ([]Product, error)
}

type productFrontStoreService struct {
	ProductFrontStoreDynamoRepository ProductFrontStoreDynamoRepository
}

func _generateUniqueId() string {
	return primitive.NewObjectID().Hex()
}

func (service productFrontStoreService) GetProducts() ([]Product, error) {
	productRecords, err := service.ProductFrontStoreDynamoRepository.Find()
	fmt.Println("getproduct ", productRecords, err)
	if err != nil {
		return nil, err
	}
	return productRecords, nil
}

func (service productFrontStoreService) GetProductById(productId string) (Product, error) {
	productRecord, err := service.ProductFrontStoreDynamoRepository.FindByProductID(productId)
	if err != nil {
		return Product{}, err
	}
	return productRecord, nil
}

func (service productFrontStoreService) GetProductsByCategoryId(categoryId string) ([]Product, error) {
	productRecords, err := service.ProductFrontStoreDynamoRepository.FindByCategoryID(categoryId)
	fmt.Println("getproduct ", productRecords, err)
	if err != nil {
		return nil, err
	}
	return productRecords, nil
}

func NewProductFrontStoreService(productFrontStoreDynamoRepository ProductFrontStoreDynamoRepository) ProductFrontStoreService {
	return &productFrontStoreService{
		ProductFrontStoreDynamoRepository: productFrontStoreDynamoRepository,
	}
}

func NewProductObject(model string, quantity int64, price float64, manufacturerID string, sku string,
	productSEOURLs []ProductSEOURL, points int64, reward int64, imageURL string, isShippable bool,
	weight float64, length float64, width float64, height float64, minimumQuantity int64,
	relatedProducts []string, productDescription []ProductDescription, productCategories []string) *Product {
	return &Product{
		Id:                  _generateUniqueId(),
		Model:               model,
		Quantity:            quantity,
		Price:               price,
		ManufacturerID:      manufacturerID,
		SKU:                 sku,
		ProductSEOURLs:      productSEOURLs,
		Points:              points,
		Reward:              reward,
		ImageURL:            imageURL,
		IsShippable:         isShippable,
		Weight:              weight,
		Length:              length,
		Width:               width,
		Height:              height,
		MinimumQuantity:     minimumQuantity,
		RelatedProducts:     relatedProducts,
		ProductDescriptions: productDescription,
		ProductCategories:   productCategories,
	}
}
