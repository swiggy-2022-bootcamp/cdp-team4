package infra

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/utils/errs"
	"time"
)

type authRepository struct {
	Session   *dynamodb.DynamoDB
	TableName string
}

func NewAuthRepository() domain.AuthRepository {
	svc := connect()
	return authRepository{Session: svc, TableName: "auth"}
}

func (repo authRepository) FindByAuthToken(authToken string) (*domain.AuthModel, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filt := expression.Name("auth_token").Equal(expression.Value(authToken)).And(
		expression.Name("is_expired").Equal(expression.Value(false)))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String("auth"),
	}

	result, err := repo.Session.ScanWithContext(ctx, input)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if len(result.Items) == 0 {
		return nil, errs.NewNotFoundError("Auth token is expired or not found")
	}

	item := result.Items[0]
	record := domain.AuthModel{}
	err = dynamodbattribute.UnmarshalMap(item, &record)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return &record, nil
}

func (repo authRepository) Save(model domain.AuthModel) *errs.AppError {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	payRecord := toPersistedEntity(model)

	av, err := dynamodbattribute.MarshalMap(payRecord)
	if err != nil {
		fmt.Println(err.Error())
		return errs.NewUnexpectedError(err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("auth"),
	}

	_, err = repo.Session.PutItemWithContext(ctx, input)

	if err != nil {
		fmt.Println(err.Error())
		return errs.NewUnexpectedError(err.Error())
	}

	return nil
}
func toPersistedEntity(model domain.AuthModel) AuthModel {
	return AuthModel{
		UserId:    model.UserId,
		AuthToken: model.AuthToken,
		IsExpired: model.IsExpired,
		ExpiresOn: model.ExpiresOn,
		Role:      model.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
