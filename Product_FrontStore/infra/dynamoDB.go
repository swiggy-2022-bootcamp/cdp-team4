package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/domain"
)

type ProductFrontStoreDynamoRepository struct {
	Session *dynamodb.DynamoDB
}

func connect() *dynamodb.DynamoDB {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	// // create dynamo client
	// svc := dynamodb.New(sess)

	// return svc
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
	})
	if err != nil {
		panic(err.Error())
	}
	// create dynamo client
	svc := dynamodb.New(sess)
	return svc
}

func (prodDynamoRepository ProductFrontStoreDynamoRepository) Find() ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Print("find function in dynamodb.go")
	input := &dynamodb.ScanInput{
		TableName: aws.String("Product"),
	}
	result, err := prodDynamoRepository.Session.ScanWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	// Make the DynamoDB Query API call
	var products = []domain.Product{}
	for _, item := range result.Items {
		product := domain.Product{}
		if err := dynamodbattribute.UnmarshalMap(item, product); err != nil {
			return []domain.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (padr ProductFrontStoreDynamoRepository) FindByProductID(productID string) (domain.Product, error) {
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

//Todo - create seconday index on category id in the product table
func (padr ProductFrontStoreDynamoRepository) FindByCategoryID(productID string) ([]domain.Product, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// input := &dynamodb.GetItemInput{
	// 	TableName: aws.String("Product"),
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"id": {
	// 			S: aws.String(productID),
	// 		},
	// 	},
	// }
	var products = []domain.Product{}
	return products, nil
}

func NewDynamoRepository() ProductFrontStoreDynamoRepository {
	svc := connect()
	return ProductFrontStoreDynamoRepository{Session: svc}
}
