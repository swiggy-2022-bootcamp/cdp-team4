package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
)

type OrderDynamoRepository struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}

func connect() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8042"),
	})

	if err != nil {
		panic(err.Error())
	}

	// create dynamo client
	svc := dynamodb.New(sess)

	return svc
}

func (odr OrderDynamoRepository) InsertOrder(p domain.Order) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	orderRecord := toPersistedDynamodbEntity(p)
	// type no struct {
	// 	ID     int
	// 	UserID int
	// 	Status string
	// }
	// var n no = no{
	// 	ID:     121323,
	// 	UserID: 141244,
	// 	Status: "confirmed",
	// }
	av, err := dynamodbattribute.MarshalMap(orderRecord)
	if err != nil {
		return "", fmt.Errorf("unable to marshal - %s", err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Order"),
	}

	_, err = odr.Session.PutItemWithContext(ctx, input)

	if err != nil {
		return "", fmt.Errorf("unable to put the item - %s", err.Error())
	}

	return orderRecord.ID, nil
}

func (pdr OrderDynamoRepository) FindOrderById(orderID string) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Order"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(orderID),
			},
		},
	}

	result, err := pdr.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("unable to get the item - %s", err.Error())
	}

	if result.Item == nil {
		return nil, fmt.Errorf("item not found")
	}

	orderModel := domain.Order{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &orderModel)

	if err != nil {
		return nil, fmt.Errorf("unmarshal map - %s", err.Error())
	}

	return &orderModel, nil
}

func (pdr OrderDynamoRepository) FindOrderByUserId(userId string) ([]domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filt := expression.Name("user_id").Equal(expression.Value(userId))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		return nil, fmt.Errorf("expression new builder - %s", err.Error())
	}

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String("Order"),
	}

	result, err := pdr.Session.ScanWithContext(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("scan with filter - %s", err.Error())
	}

	orderRecords := make([]domain.Order, 1)

	for _, item := range result.Items {
		record := domain.Order{}
		err := dynamodbattribute.UnmarshalMap(item, &record)
		if err != nil {
			return nil, fmt.Errorf("expression new builder - %s", err.Error())
		}
		orderRecords = append(orderRecords, record)
	}

	return orderRecords, nil
}

func (pdr OrderDynamoRepository) UpdateOrderStatus(id, attributeValue string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(attributeValue),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set order_status = :s"),
		TableName:        aws.String("Order"),
	}

	_, err := pdr.Session.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to update - %s", err.Error())
	}
	return true, nil
}

func (pdr OrderDynamoRepository) DeleteByID(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String("Order"),
	}

	_, err := pdr.Session.DeleteItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to delete - %s", err.Error())
	}
	return true, nil
}

// func NewDynamoRepository() OrderDynamoRepository {
// 	svc := connect()
// 	return OrderDynamoRepository{Session: svc}
// }

func toPersistedDynamodbEntity(o domain.Order) *OrderModel {
	return &OrderModel{
		ID:               uuid.New().String(),
		UserID:           o.UserID,
		Status:           o.Status,
		DateTime:         o.DateTime,
		ProductsQuantity: o.ProductsQuantity,
		ProductsCost:     o.ProductsCost,
		TotalCost:        o.TotalCost,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func NewDynamoRepository() OrderDynamoRepository {
	svc := connect()
	return OrderDynamoRepository{Session: svc, Tablename: "Order"}
}

// func (dyr OrderDynamoRepository) ListTables() ([]string, error) {
// 	input := &dynamodb.ListTablesInput{}

// 	// if docker container of dynamoDB is not running then code
// 	// should be blocked indefinitely, that's why using context with time out.
// 	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	var tableNames []string

// 	result, err := dyr.Session.ListTablesWithContext(context, input)
// 	if err != nil {
// 		log.Fatal(err)
// 		return tableNames, err
// 	}

// 	for _, n := range result.TableNames {
// 		tableNames = append(tableNames, *n)
// 	}

// 	return tableNames, nil
// }

// func (dyr OrderDynamoRepository) CreateTable(tableName string) bool {
// 	// input := &dynamodb.CreateTableInput{}

// 	// context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancel()
// 	return true
// }
