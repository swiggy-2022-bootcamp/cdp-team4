package infra

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/utils/errs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
)

type TransactionDynamoRepository struct {
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
		TableName: aws.String("Transaction"),
	}

	_, tableErr := svc.CreateTable(tableInput)

	if tableErr != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(tableErr.Error())
	}
}

func connect() *dynamodb.DynamoDB {

	// sess, err := session.NewSession(&aws.Config{
	// 	Region:      aws.String("us-east-1"),
	// 	Endpoint:    aws.String("http://localhost:8000"),
	// 	Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
	// })

	// if err != nil {
	// 	panic(err.Error())
	// }

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// create dynamo client
	svc := dynamodb.New(sess)

	createTable(svc)

	return svc
}

func (rwd TransactionDynamoRepository) InsertTransaction(p domain.Transaction) (string, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	transactionRecord := toPersistedDynamodbEntity(p)

	av, err := dynamodbattribute.MarshalMap(transactionRecord)
	if err != nil {
		errstring := fmt.Sprintf("unable to marshal - %s", err.Error())
		return "", &errs.AppError{Message: errstring}
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Transaction"),
	}

	_, err = rwd.Session.PutItemWithContext(ctx, input)

	if err != nil {
		errstring := fmt.Sprintf("unable to put item - %s", err.Error())
		return "", &errs.AppError{Message: errstring}
	}

	return transactionRecord.Id, nil
}

func (rwd TransactionDynamoRepository) FindTransactionById(transactionID string) (*domain.Transaction, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Transaction"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(transactionID),
			},
		},
	}

	result, err := rwd.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, &errs.AppError{Message: "Failed to read"}
	}

	if result.Item == nil {
		return nil, &errs.AppError{Message: "Item not found"}
	}

	transactionModel := TransactionModel{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &transactionModel)

	if err != nil {
		errstring := fmt.Sprintf("unmarshal map - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}
	rwdModel := toModelfromDynamodbEntity(transactionModel)
	return rwdModel, nil
}

func (rwd TransactionDynamoRepository) FindTransactionByUserId(userId string) (*domain.Transaction, *errs.AppError) {
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
		TableName:                 aws.String("Transaction"),
	}

	result, err := rwd.Session.ScanWithContext(ctx, input)
	if err != nil {
		errstring := fmt.Sprintf("scan with filter - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}

	if len(result.Items) == 0 {
		return nil, &errs.AppError{Message: "Item not found"}
	}

	record := TransactionModel{}
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &record)
	if err != nil {
		fmt.Println("Stray records inside db")
	}

	rwdModel := toModelfromDynamodbEntity(record)

	return rwdModel, nil
}

func (rwd TransactionDynamoRepository) UpdateTransactionByUserId(userId string, points int) (bool, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//find the current record
	currentTransaction, err := rwd.FindTransactionByUserId(userId)
	if err != nil {
		newDomainRecord := domain.NewTransaction(userId, points)
		_, err = rwd.InsertTransaction(*newDomainRecord)
		if err != nil {
			return false, &errs.AppError{Message: "Unable to update"}
		}
		return true, nil
	}

	updatedTransactionPoints := currentTransaction.TransactionPoints + points
	if updatedTransactionPoints < 0 {
		return false, &errs.AppError{Message: "Transaction point shortage"}
	}
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				N: aws.String(strconv.Itoa(updatedTransactionPoints)),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(currentTransaction.Id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set transaction_points = :s"),
		TableName:        aws.String("Transaction"),
	}

	_, updError := rwd.Session.UpdateItemWithContext(ctx, input)
	if updError != nil {
		return false, &errs.AppError{Message: updError.Error() + "Unable to update"}
	}
	fmt.Println("Updated exisitng user with points", updatedTransactionPoints)
	return true, nil
}

// func (odr TransactionDynamoRepository) FindAllTransactionUserId(userId string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	filt := expression.Name("user_id").Equal(expression.Value(userId))

// 	expr, err := expression.NewBuilder().WithFilter(filt).Build()

// 	if err != nil {
// 		errstring := fmt.Sprintf("expression new builder - %s", err.Error())
// 		fmt.Println(&errs.AppError{Message: errstring})
// 	}

// 	input := &dynamodb.ScanInput{
// 		ExpressionAttributeNames:  expr.Names(),
// 		FilterExpression:          expr.Filter(),
// 		ExpressionAttributeValues: expr.Values(),
// 		TableName:                 aws.String("Transaction"),
// 	}

// 	result, err := odr.Session.ScanWithContext(ctx, input)
// 	if err != nil {
// 		errstring := fmt.Sprintf("scan with filter - %s", err.Error())
// 		fmt.Println(&errs.AppError{Message: errstring})
// 	}

// 	for _, item := range result.Items {
// 		record := TransactionModel{}
// 		err := dynamodbattribute.UnmarshalMap(item, &record)
// 		if err != nil {
// 			fmt.Println("Stray records inside db")
// 		} else {
// 			recordm := toModelfromDynamodbEntity(record)
// 			fmt.Println(recordm)
// 		}
// 	}
// }

func toPersistedDynamodbEntity(r domain.Transaction) *TransactionModel {
	return &TransactionModel{
		Id:                uuid.New().String(),
		UserID:            r.UserID,
		TransactionPoints: r.TransactionPoints,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func toModelfromDynamodbEntity(r TransactionModel) *domain.Transaction {
	return &domain.Transaction{
		Id:                r.Id,
		UserID:            r.UserID,
		TransactionPoints: r.TransactionPoints,
	}
}

func NewDynamoRepository() TransactionDynamoRepository {
	svc := connect()
	return TransactionDynamoRepository{Session: svc, Tablename: "Transaction"}
}
