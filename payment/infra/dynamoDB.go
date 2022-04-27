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

// function to connect with dynamoDB with the credentials stored in
// the local system
func connect() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// create dynamoDB client
	svc := dynamodb.New(sess)

	return svc
}

// inserting the record in dynamoDB
// Tablename = payment
//
// ref: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-create-table-item.html
func (pdr PaymentDynamoRepository) InsertPaymentRecord(p domain.Payment) (bool, error) {
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

// get the record using unique identifier
//
// ref: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-read-table-item.html
func (pdr PaymentDynamoRepository) FindPaymentRecordById(
	paymentID string,
) (*domain.Payment, error) {
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

// get the payment records using user id
// it is going to return all the payments done by the user so far
//
// ref: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-scan-table-item.html
func (pdr PaymentDynamoRepository) FindPaymentRecordByUserID(
	userId string,
) ([]domain.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create the Expression to fill the input struct with.
	// Get all the payment records of given user id
	filt := expression.Name("UserID").Equal(expression.Value(userId))

	// Get back every fields of the record
	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		return nil, fmt.Errorf("expression new builder - %s", err.Error())
	}

	// Build the query input parameters
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String("payment"),
	}

	// Make the DynamoDB Query API call with context
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

// upate payment status, it can be done, pending or failed
//
// ref: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-update-table-item.html
func (pdr PaymentDynamoRepository) UpdatePaymentRecord(
	id, attributeValue string,
) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// create update input with expression
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				N: aws.String(attributeValue),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Status = :s"),
		TableName:        aws.String("payment"),
	}

	_, err := pdr.Session.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to update - %s", err.Error())
	}
	return true, nil
}

// delete payment record by id and if it is not present then
// handle it gracefully
//
// ref: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-delete-table-item.html
func (pdr PaymentDynamoRepository) DeletePaymentRecordByID(id string) (bool, error) {
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

// function is used to insert new record of supported payment method
// for the particular user
func (pdr PaymentDynamoRepository) InsertPaymentMethod(
	pm domain.PaymentMethod,
) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	payMethodModel := _toDynamoPayMethodModel(&pm)
	av, err := dynamodbattribute.MarshalMap(payMethodModel)
	if err != nil {
		return false, fmt.Errorf("unable to marshal - %s", err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("paymentmethod"),
	}

	// make dynamo API cal
	_, err = pdr.Session.PutItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to put the item - %s", err.Error())
	}

	return true, nil
}

// get the supported payment methods which is an array of strings
// eg. ["debit-card","credit-card","upi"]
func (pdr PaymentDynamoRepository) GetPaymentMethods(id string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("paymentmethod"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
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

	payMethodModel := PaymentMethodModel{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &payMethodModel)

	if err != nil {
		return nil, fmt.Errorf("unmarshal map - %s", err.Error())
	}

	return payMethodModel.Methods, nil
}

// update the current existing list of supported payment methods
// for the particular user
// before updating the array of methods, it first checks whether the element is
// already present or not. if it is not there then only insert/push the element in the
// array otherwise do nothing.
func (pdr PaymentDynamoRepository) UpdatePaymentMethods(
	id, paymentMethod string,
) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	av := &dynamodb.AttributeValue{
		S: aws.String(paymentMethod),
	}

	var methodList []*dynamodb.AttributeValue
	methodList = append(methodList, av)

	// create expression to push element in array datatype
	// and check method exists or not before pushing it
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":method": {
				L: methodList,
			},
			":methodStr": av,
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ReturnValues: aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(
			"set methods = list_append (methods, :method)",
		),
		ConditionExpression: aws.String("not contains (methods, :methodStr)"),
		TableName:           aws.String("paymentmethod"),
	}

	_, err := pdr.Session.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to update - %s", err.Error())
	}
	return true, nil
}

// constructor method to get dynamo Repository after connecting with it
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

func _toDynamoPayMethodModel(p *domain.PaymentMethod) PaymentMethodModel {
	return PaymentMethodModel{
		Id:      p.Id,
		Methods: p.Method,
		Agree:   p.Agree,
		Comment: p.Comment,
	}
}
