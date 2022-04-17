package domain

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
