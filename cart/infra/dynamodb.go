package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	// "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/utils/errs"
)

type CartDynamoRepository struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}

func createTable(svc *dynamodb.DynamoDB) {
	tableInput := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("Cart"),
	}

	_, tableErr := svc.CreateTable(tableInput)

	if tableErr != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(tableErr.Error())
	}
}

func connect() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Endpoint:    aws.String("http://localhost:8000"),
		Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
	})

	if err != nil {
		panic(err.Error())
	}

	// create dynamo client
	svc := dynamodb.New(sess)

	//createTable(svc)

	return svc
}

func (crt CartDynamoRepository) InsertCart(p domain.Cart) (string, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cartRecord := toPersistedDynamodbEntity(p)

	av, err := dynamodbattribute.MarshalMap(cartRecord)
	if err != nil {
		errstring := fmt.Sprintf("unable to marshal - %s", err.Error())
		return "", &errs.AppError{Message: errstring}
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Cart"),
	}

	_, err = crt.Session.PutItemWithContext(ctx, input)

	if err != nil {
		errstring := fmt.Sprintf("unable to put item - %s", err.Error())
		return "", &errs.AppError{Message: errstring}
	}

	return cartRecord.Id, nil
}

func (crt CartDynamoRepository) FindAllCarts() ([]domain.Cart, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	input := &dynamodb.ScanInput{
		TableName: aws.String("Cart"),
	}
	out, err := crt.Session.ScanWithContext(ctx, input)

	if err != nil {
		return nil, &errs.AppError{Message: err.Error()}
	}

	cartRecords := make([]domain.Cart, 0)

	for _, item := range out.Items {
		record := CartModel{}
		err := dynamodbattribute.UnmarshalMap(item, &record)
		if err != nil {
			errstring := fmt.Sprintf("expression new builder - %s", err.Error())
			return nil, &errs.AppError{Message: errstring}
		}
		recordm := toModelfromDynamodbEntity(record)
		cartRecords = append(cartRecords, *recordm)
	}

	return cartRecords, nil
}

func (odr CartDynamoRepository) FindCartById(cartID string) (*domain.Cart, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Cart"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(cartID),
			},
		},
	}

	result, err := odr.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, &errs.AppError{Message: "Failed to read"}
	}

	if result.Item == nil {
		return nil, &errs.AppError{Message: "Item not found"}
	}

	cartModel := CartModel{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &cartModel)

	if err != nil {
		errstring := fmt.Sprintf("unmarshal map - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}
	ordModel := toModelfromDynamodbEntity(cartModel)
	return ordModel, nil
}

func (crt CartDynamoRepository) DeleteCartById(id string) (bool, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String("Cart"),
	}

	_, err := crt.Session.DeleteItemWithContext(ctx, input)
	if err != nil {
		errstring := fmt.Sprintf("unable to delete - %s", err.Error())
		return false, &errs.AppError{Message: errstring}
	}
	return true, nil
}

func toPersistedDynamodbEntity(o domain.Cart) *CartModel {
	return &CartModel{
		Id:               uuid.New().String(),
		UserID:           o.UserID,
		ProductsQuantity: o.ProductsQuantity,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func toModelfromDynamodbEntity(c CartModel) *domain.Cart {
	return &domain.Cart{
		Id:               c.Id,
		UserID:           c.UserID,
		ProductsQuantity: c.ProductsQuantity,
	}
}

func NewDynamoRepository() CartDynamoRepository {
	svc := connect()
	return CartDynamoRepository{Session: svc, Tablename: "Cart"}
}
