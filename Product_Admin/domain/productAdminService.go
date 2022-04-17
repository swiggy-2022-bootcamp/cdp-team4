package domain

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductAdminService interface {
	CreateDynamoProductAdminRecord(string, int64, float64, string, string,
		[]ProductSEOURL, int64, int64, string, bool, float64, float64, float64,
		float64, int64, []string, []ProductDescription, []ProductCategory) (string, error)
	GetProduct() ([]Product, error)
	GetProductById(string) (Product, error)
	UpdateProduct(string, int64) (bool, error)
	DeleteProductById(string) (bool, error)
}

type productAdminService struct {
	ProductAdminDynamoRepository ProductAdminDynamoRepository
}

func _generateUniqueId() string {
	return primitive.NewObjectID().Hex()
}
func (service productAdminService) CreateDynamoProductAdminRecord(model string, quantity int64,
	price float64, manufacturerID string, sku string, productSEOURLs []ProductSEOURL, points int64,
	reward int64, imageURL string, isShippable bool, weight float64, length float64, width float64,
	height float64, minimumQunatity int64, relatedProducts []string, productDescription []ProductDescription,
	productCategories []ProductCategory) (string, error) {
	id := _generateUniqueId()
	productRecord := Product{
		Id:                  id,
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
		MinimumQuantity:     minimumQunatity,
		RelatedProducts:     relatedProducts,
		ProductDescriptions: productDescription,
		ProductCategories:   productCategories,
	}

	ok, err := service.ProductAdminDynamoRepository.Insert(productRecord)
	if !ok {
		return id, err
	}
	return id, nil
}

func (service productAdminService) GetProduct() ([]Product, error) {
	productRecords, err := service.ProductAdminDynamoRepository.Find()
	fmt.Println("getproduct ", productRecords, err)
	if err != nil {
		return nil, err
	}
	return productRecords, nil
}

func (service productAdminService) GetProductById(id string) (Product, error) {
	productRecord, err := service.ProductAdminDynamoRepository.FindByID(id)
	if err != nil {
		return Product{}, err
	}
	return productRecord, nil
}

func (service productAdminService) UpdateProduct(productId string, quantityReduction int64) (bool, error) {
	_, err := service.ProductAdminDynamoRepository.UpdateItem(productId, quantityReduction)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (service productAdminService) DeleteProductById(productId string) (bool, error) {
	_, err := service.ProductAdminDynamoRepository.DeleteByID(productId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewProductAdminService(productAdminDynamoRepository ProductAdminDynamoRepository) ProductAdminService {
	return &productAdminService{
		ProductAdminDynamoRepository: productAdminDynamoRepository,
	}
}