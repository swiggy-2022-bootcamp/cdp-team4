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
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/utils/errs"
)

var OrderTableName string = "order"
var OrderOverviewTableName string = "order-overview"

type OrderDynamoRepository struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}

func connect() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// sess, _ := session.NewSession(&aws.Config{
	// 	Region:   aws.String("us-east-1"),
	// 	Endpoint: aws.String("http://localhost:8042"),
	// })

	// create dynamo client
	svc := dynamodb.New(sess)

	return svc
}

func (odr OrderDynamoRepository) InsertOrder(p domain.Order) (string, *errs.AppError) {
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
		errstring := fmt.Sprintf("unable to marshal - %s", err.Error())
		return "", &errs.AppError{Message: errstring}
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(OrderTableName),
	}

	_, err = odr.Session.PutItemWithContext(ctx, input)

	if err != nil {
		errstring := fmt.Sprintf("unable to put item - %s", err.Error())
		return "", &errs.AppError{Message: errstring}
	}

	return orderRecord.ID, nil
}

func (odr OrderDynamoRepository) FindOrderById(orderID string) (*domain.Order, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String(OrderTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(orderID),
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

	orderModel := OrderModel{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &orderModel)

	if err != nil {
		errstring := fmt.Sprintf("unmarshal map - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}
	ordModel := toModelfromDynamodbEntity(orderModel)
	return ordModel, nil
}

func (odr OrderDynamoRepository) FindOrderByUserId(userId string) ([]domain.Order, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filt := expression.Name("user_id").Equal(expression.Value(userId))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		errstring := fmt.Sprintf("expression new builder - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String(OrderTableName),
	}

	result, err := odr.Session.ScanWithContext(ctx, input)
	if err != nil {
		errstring := fmt.Sprintf("scan with filter - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}

	orderRecords := make([]domain.Order, 0)

	for _, item := range result.Items {
		record := OrderModel{}
		err := dynamodbattribute.UnmarshalMap(item, &record)
		if err != nil {
			fmt.Println("Stray records inside db")
		} else {
			recordm := toModelfromDynamodbEntity(record)
			orderRecords = append(orderRecords, *recordm)
		}
	}

	return orderRecords, nil
}

func (odr OrderDynamoRepository) FindOrderByStatus(status string) ([]domain.Order, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filt := expression.Name("order_status").Equal(expression.Value(status))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		errstring := fmt.Sprintf("expression new builder - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String(OrderTableName),
	}

	result, err := odr.Session.ScanWithContext(ctx, input)
	if err != nil {
		errstring := fmt.Sprintf("scan with filter - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}

	orderRecords := make([]domain.Order, 0)

	for _, item := range result.Items {
		record := OrderModel{}
		err := dynamodbattribute.UnmarshalMap(item, &record)
		if err != nil {
			fmt.Println("Stray records inside db")
		} else {
			recordm := toModelfromDynamodbEntity(record)
			orderRecords = append(orderRecords, *recordm)
		}
	}

	return orderRecords, nil
}

func (odr OrderDynamoRepository) UpdateOrderStatus(id, attributeValue string) (bool, *errs.AppError) {
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
		TableName:        aws.String(OrderTableName),
	}

	_, err := odr.Session.UpdateItemWithContext(ctx, input)
	if err != nil {
		errstring := fmt.Sprintf("unable to update - %s", err.Error())
		return false, &errs.AppError{Message: errstring}
	}
	return true, nil
}

func (odr OrderDynamoRepository) DeleteOrderById(id string) (bool, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(OrderTableName),
	}

	_, err := odr.Session.DeleteItemWithContext(ctx, input)
	if err != nil {
		errstring := fmt.Sprintf("unable to delete - %s", err.Error())
		return false, &errs.AppError{Message: errstring}
	}
	return true, nil
}

func (odr OrderDynamoRepository) FindAllOrders() ([]domain.Order, *errs.AppError) {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	input := &dynamodb.ScanInput{
		TableName: aws.String(OrderTableName),
	}

	result, err := odr.Session.Scan(input)
	if err != nil {
		errstring := fmt.Sprintf("scan with filter - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}

	orderRecords := make([]domain.Order, 0)
	//fmt.Print("items", result)
	for _, item := range result.Items {
		record := OrderModel{}
		//fmt.Print("single item", item)
		err := dynamodbattribute.UnmarshalMap(item, &record)
		if err != nil {
			fmt.Println("Stray records inside db")
		} else {
			recordm := toModelfromDynamodbEntity(record)
			orderRecords = append(orderRecords, *recordm)
		}
	}

	return orderRecords, nil
}

func (odr OrderDynamoRepository) InsertOrderOverview(ov domain.OrderOverview) (bool, *errs.AppError) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	av, err := dynamodbattribute.MarshalMap(ov)
	if err != nil {
		errstring := fmt.Sprintf("unable to marshal - %s", err.Error())
		return false, &errs.AppError{Message: errstring}
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(OrderOverviewTableName),
	}

	_, err = odr.Session.PutItemWithContext(ctx, input)

	if err != nil {
		errstring := fmt.Sprintf("unable to put item - %s", err.Error())
		return false, &errs.AppError{Message: errstring}
	}

	return true, nil

}

func (odr OrderDynamoRepository) GetOrderOverview(orderid string) (*domain.OrderOverview, *errs.AppError) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String(OrderTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"order_id": {
				S: aws.String(orderid),
			},
		},
	}

	result, err := odr.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return &domain.OrderOverview{}, &errs.AppError{Message: "Failed to read"}
	}

	if result.Item == nil {
		return &domain.OrderOverview{}, &errs.AppError{Message: "Item not found"}
	}

	ovModel := OrderOverview{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &ovModel)

	if err != nil {
		errstring := fmt.Sprintf("unmarshal map - %s", err.Error())
		return &domain.OrderOverview{}, &errs.AppError{Message: errstring}
	}
	return &domain.OrderOverview{
		OrderID:            ovModel.OrderID,
		ProductsIdQuantity: ovModel.ProductsIdQuantity,
	}, nil
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

func toModelfromDynamodbEntity(o OrderModel) *domain.Order {
	return &domain.Order{
		ID:               o.ID,
		UserID:           o.UserID,
		Status:           o.Status,
		DateTime:         o.DateTime,
		ProductsQuantity: o.ProductsQuantity,
		ProductsCost:     o.ProductsCost,
		TotalCost:        o.TotalCost,
	}
}

func NewDynamoRepository() OrderDynamoRepository {
	svc := connect()
	return OrderDynamoRepository{Session: svc, Tablename: OrderTableName}
}

func NewDynomoOrderOverviewRepository() OrderDynamoRepository {
	svc := connect()
	return OrderDynamoRepository{Session: svc, Tablename: OrderOverviewTableName}
}
