package infra

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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

func (padr ProductAdminDynamoRepository) GetProductAvailability(productId string, QuantityNeeded int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Product"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(productId),
			},
		},
	}

	result, err := padr.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to get the item - %s", err.Error())
	}

	if result.Item == nil {
		return false, fmt.Errorf("item not found")
	}

	productModel := domain.Product{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &productModel)
	if err != nil {
		return false, fmt.Errorf("unmarshal map - %s", err.Error())
	}
	if productModel.Quantity < QuantityNeeded {
		return false, fmt.Errorf("product quantity is not enough for this order")
	}
	return true, nil
}

func (padr ProductAdminDynamoRepository) FindByCategoryID(categoryId string) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := expression.Name("category_id").Equal(expression.Value(categoryId))
	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return nil, fmt.Errorf("expression new builder - %s", err.Error())
	}
	fmt.Println("categoryid", categoryId)
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String("ProductCategoryRelation"),
	}
	result, err := padr.Session.ScanWithContext(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("scan with filter - %s", err.Error())
	}

	var products = []domain.Product{}
	for _, item := range result.Items {
		record := ProductCategory{}
		err := dynamodbattribute.UnmarshalMap(item, &record)
		if err != nil {
			return nil, fmt.Errorf("expression new builder - %s", err.Error())
		}
		product, err := padr.FindByID(record.ProductId)
		if err != nil {
			return nil, fmt.Errorf("error in fetching product details -%s", err.Error())
		}
		products = append(products, product)
	}
	return products, nil
}

func (padr ProductAdminDynamoRepository) FindByManufacturerID(manufacturerId string) ([]domain.Product, error) {
	filter := expression.Name("manufacturer_id").Equal(expression.Value(manufacturerId))

	// Build condition from above filter
	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return []domain.Product{}, err
	}
	_products, err2 := padr.GetProductsByCondition(condition)
	if err2 != nil {
		return []domain.Product{}, err2
	}
	return _products, nil
}

func (padr ProductAdminDynamoRepository) FindByKeyword(keyword string) ([]domain.Product, error) {
	// Define the filter expression for searching product by keyword
	filter1 := expression.Contains(expression.Name("model"), keyword)
	filter2 := expression.Contains(expression.Name("model"), strings.ToUpper(keyword))
	filter3 := expression.Contains(expression.Name("sku"), keyword)
	filter4 := expression.Contains(expression.Name("sku"), strings.ToUpper(keyword))
	filter := filter1.Or(filter2).Or(filter3).Or(filter4)
	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return []domain.Product{}, err
	}
	_products, err2 := padr.GetProductsByCondition(condition)
	if err2 != nil {
		return []domain.Product{}, err2
	}
	return _products, nil
}

func (padr ProductAdminDynamoRepository) GetProductsByCondition(condition expression.Expression) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String("Product"),
	}
	response, err := padr.Session.ScanWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	// Unmarshal dynamodb map to domain.Product
	products := []domain.Product{}
	for _, _dbProduct := range response.Items {
		var _product domain.Product
		err = dynamodbattribute.UnmarshalMap(_dbProduct, &_product)
		if err != nil {
			return nil, err
		}
		products = append(products, _product)
	}
	return products, nil
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
