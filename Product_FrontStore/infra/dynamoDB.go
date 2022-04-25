package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/domain"
)

type ProductFrontStoreDynamoRepository struct {
	Session *dynamodb.DynamoDB
}

// function to connect with dynamoDB with the credentials stored in
// the local system
func connect() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// create dynamo client
	svc := dynamodb.New(sess)
	return svc
}

//Function to fetch all the product details from aws datastore
func (prodDynamoRepository ProductFrontStoreDynamoRepository) Find() ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.ScanInput{
		TableName: aws.String(PRODUCT_TABLE),
	}
	result, err := prodDynamoRepository.Session.ScanWithContext(ctx, input)
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
//Function to fetch the product details, given product id from aws datastore
func (padr ProductFrontStoreDynamoRepository) FindByProductID(productID string) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.GetItemInput{
		TableName: aws.String(PRODUCT_TABLE),
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

//Function to fetch the product details, given category id from aws datastore
func (padr ProductFrontStoreDynamoRepository) FindByCategoryID(categoryID string) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := expression.Name("category_id").Equal(expression.Value(categoryID))
	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return nil, fmt.Errorf("expression new builder - %s", err.Error())
	}
	fmt.Println("categoryid", categoryID)
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String(PRODUCT_CATEGORY_TABLE),
	}
	result, err := padr.Session.ScanWithContext(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("scan with filter - %s", err.Error())
	}
	fmt.Println("====================", result.Items)
	var products = []domain.Product{}
	for _, item := range result.Items {
		record := ProductCategoryRelation{}
		err := dynamodbattribute.UnmarshalMap(item, &record)
		if err != nil {
			return nil, fmt.Errorf("expression new builder - %s", err.Error())
		}
		fmt.Println("productid", record.ProductID)
		product, err := padr.FindByProductID(record.ProductID)
		if err != nil {
			return nil, fmt.Errorf("error in fetching product details -%s", err.Error())
		}
		products = append(products, product)
	}
	return products, nil
}

func NewDynamoRepository() ProductFrontStoreDynamoRepository {
	svc := connect()
	return ProductFrontStoreDynamoRepository{Session: svc}
}
