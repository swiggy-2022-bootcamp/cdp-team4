package infra

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/domain"
)

type ProductAdminDynamoRepository struct {
	Session *dynamodb.DynamoDB
}

func connect() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// create dynamo client
	svc := dynamodb.New(sess)

	return svc
	// sess, err := session.NewSession(&aws.Config{
	// 	Region:   aws.String("us-east-1"),
	// 	Endpoint: aws.String("http://localhost:8000"),
	// })
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // create dynamo client
	// svc := dynamodb.New(sess)
	// return svc
}

type ProductCategory struct {
	CategoryId string `json:"category_id"`
	ProductId  string `json:"product_id"`
}

func (padr ProductAdminDynamoRepository) Insert(product domain.Product) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	productRecord := _toDynamoProductModel(&product)
	av, err := dynamodbattribute.MarshalMap(productRecord)
	if err != nil {
		return false, fmt.Errorf("unable to marshal - %s", err.Error())
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Product"),
	}

	_, err = padr.Session.PutItemWithContext(ctx, input)

	if err != nil {
		return false, fmt.Errorf("unable to insert the item - %s", err.Error())
	}

	//add product categories in relation table
	if product.ProductCategories != nil {
		for _, categoryid := range product.ProductCategories {
			productCategory := ProductCategory{CategoryId: categoryid, ProductId: product.Id}
			av, err := dynamodbattribute.MarshalMap(productCategory)
			if err != nil {
				return false, fmt.Errorf("unable to marshal - %s", err.Error())
			}
			input := &dynamodb.PutItemInput{
				Item:      av,
				TableName: aws.String("ProductCategoryRelation"),
			}
			_, err = padr.Session.PutItemWithContext(ctx, input)

			if err != nil {
				return false, fmt.Errorf("unable to insert the item - %s", err.Error())
			}
		}
	}

	return true, nil
}

func (padr ProductAdminDynamoRepository) Find() ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.ScanInput{
		TableName: aws.String("Product"),
	}
	result, err := padr.Session.ScanWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	// Make the DynamoDB Query API call
	var products = []domain.Product{}
	for _, item := range result.Items {
		product := domain.Product{}
		if err := dynamodbattribute.UnmarshalMap(item, &product); err != nil {
			return []domain.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (padr ProductAdminDynamoRepository) FindByID(productID string) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Product"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(productID),
			},
		},
	}

	result, err := padr.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return domain.Product{}, fmt.Errorf("unable to get the item - %s", err.Error())
	}

	if result.Item == nil {
		return domain.Product{}, fmt.Errorf("item not found")
	}

	productModel := domain.Product{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &productModel)
	if err != nil {
		return domain.Product{}, fmt.Errorf("unmarshal map - %s", err.Error())
	}
	return productModel, nil

}

func (padr ProductAdminDynamoRepository) UpdateItem(productID string, quantiyReduction int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//find the total quantity
	currentProduct, err := padr.FindByID(productID)
	if currentProduct.Quantity < quantiyReduction {
		return false, fmt.Errorf("Product shortage - %s", err.Error())
	}
	updatedQuantity := currentProduct.Quantity - quantiyReduction
	//update the quantity
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				N: aws.String(strconv.Itoa(int(updatedQuantity))),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(productID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set quantity = :s"),
		TableName:        aws.String("Product"),
	}

	_, err = padr.Session.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to update - %s", err.Error())
	}
	return true, nil
}

func (padr ProductAdminDynamoRepository) DeleteByID(productID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(productID),
			},
		},
		TableName: aws.String("Product"),
	}

	_, err := padr.Session.DeleteItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to delete - %s", err.Error())
	}
	return true, nil
}

func NewDynamoRepository() ProductAdminDynamoRepository {
	svc := connect()
	return ProductAdminDynamoRepository{Session: svc}
}

func _toDynamoProductModel(p *domain.Product) ProductModel {
	var productSEOURLModel []ProductSEOURLModel
	for _, item := range p.ProductSEOURLs {
		productSEOURLModel = append(productSEOURLModel, ProductSEOURLModel{Keyword: item.Keyword, LanguageID: item.LanguageID, StoreID: item.StoreID})
	}
	var productDescriptionModel []ProductDescriptionModel
	for _, item := range p.ProductDescriptions {
		productDescriptionModel = append(productDescriptionModel, ProductDescriptionModel{LanguageID: item.LanguageID, Name: item.Name,
			Description: item.Description, MetaTitle: item.MetaTitle, MetaDescription: item.MetaDescription, MetaKeyword: item.MetaKeyword,
			Tag: item.Tag})
	}
	var productCategoriesModel []string
	for _, item := range p.ProductCategories {
		productCategoriesModel = append(productCategoriesModel, item)
	}

	return ProductModel{
		Id:                  p.Id,
		Model:               p.Model,
		Quantity:            p.Quantity,
		Price:               p.Price,
		ManufacturerID:      p.ManufacturerID,
		SKU:                 p.SKU,
		ProductSEOURLs:      productSEOURLModel,
		Points:              p.Points,
		Reward:              p.Reward,
		ImageURL:            p.ImageURL,
		IsShippable:         p.IsShippable,
		Weight:              p.Weight,
		Length:              p.Length,
		Width:               p.Width,
		Height:              p.Height,
		MinimumQuantity:     p.MinimumQuantity,
		RelatedProducts:     p.RelatedProducts,
		ProductDescriptions: productDescriptionModel,
		ProductCategories:   productCategoriesModel,
	}
}
