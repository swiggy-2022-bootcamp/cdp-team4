package infra

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/utils/errs"
	"time"
)

type userRepository struct {
	Session   *dynamodb.DynamoDB
	TableName string
}

func NewUserRepository() domain.UserRepository {
	svc := connect()
	return userRepository{Session: svc, TableName: "users"}
}

func connect() *dynamodb.DynamoDB {
	// Create AWS Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Return DynamoDB client
	return dynamodb.New(sess)
}

func (repo userRepository) FindByUsername(username string) (*domain.UserModel, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filt := expression.Name("username").Equal(expression.Value(username))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String("users"),
	}

	result, err := repo.Session.ScanWithContext(ctx, input)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	//paymentRecords := make([]domain.Payment, 1)

	item := result.Items[0]
	record := domain.UserModel{}
	err = dynamodbattribute.UnmarshalMap(item, &record)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	//paymentRecords = append(paymentRecords, record)

	return &record, nil
}
