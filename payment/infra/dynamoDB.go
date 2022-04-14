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
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
)

type PaymentDynamoRepository struct {
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

func (pdr PaymentDynamoRepository) Insert(p domain.Payment) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	payRecord := _toDynamoPayModel(&p, "XYZ", "0")

	av, err := dynamodbattribute.MarshalMap(payRecord)
	if err != nil {
		return false, fmt.Errorf("unable to marshal - %s", err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("payment"),
	}

	_, err = pdr.Session.PutItemWithContext(ctx, input)

	if err != nil {
		return false, fmt.Errorf("unable to put the item - %s", err.Error())
	}

	return true, nil
}

func (pdr PaymentDynamoRepository) FindById(paymentID string) (*domain.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("payment"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(paymentID),
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

	payModel := domain.Payment{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &payModel)

	if err != nil {
		return nil, fmt.Errorf("unmarshal map - %s", err.Error())
	}

	return &payModel, nil
}

func (pdr PaymentDynamoRepository) FindByUserID(userId string) ([]domain.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filt := expression.Name("UserID").Equal(expression.Value(userId))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		return nil, fmt.Errorf("expression new builder - %s", err.Error())
	}

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String("payment"),
	}

	result, err := pdr.Session.ScanWithContext(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("scan with filter - %s", err.Error())
	}

	paymentRecords := make([]domain.Payment, 1)

	for _, item := range result.Items {
		record := domain.Payment{}
		err := dynamodbattribute.UnmarshalMap(item, &record)
		if err != nil {
			return nil, fmt.Errorf("expression new builder - %s", err.Error())
		}
		paymentRecords = append(paymentRecords, record)
	}

	return paymentRecords, nil
}

func (pdr PaymentDynamoRepository) UpdateItem(id, attributeValue string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				N: aws.String(attributeValue),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
		ReturnValues:     aws.String("UPDATE_NEW"),
		UpdateExpression: aws.String("set Status = :s"),
		TableName:        aws.String("payment"),
	}

	_, err := pdr.Session.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to update - %s", err.Error())
	}
	return true, nil
}

func (pdr PaymentDynamoRepository) DeleteByID(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(id),
			},
		},
		TableName: aws.String("payment"),
	}

	_, err := pdr.Session.DeleteItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to delete - %s", err.Error())
	}
	return true, nil
}

func NewDynamoRepository() PaymentDynamoRepository {
	svc := connect()
	return PaymentDynamoRepository{Session: svc}
}

func _toDynamoPayModel(p *domain.Payment, bank, wallet string) PayModel {
	return PayModel{
		Id:          p.Id,
		Amount:      p.Amount,
		Currency:    p.Currency,
		Status:      p.Status,
		OrderID:     p.OrderID,
		Method:      p.Method,
		Description: p.Description,
		Bank:        bank,
		Wallet:      wallet,
		VPA:         p.VPA,
		UserID:      p.UserID,
		Notes:       p.Notes,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
