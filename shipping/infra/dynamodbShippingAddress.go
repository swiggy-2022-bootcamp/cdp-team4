package infra

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"
)

type ShippingAddressDynamoRepository struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}

func connect() *dynamodb.DynamoDB {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))
	sess, _ := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8042"),
	})

	// create dynamo client
	svc := dynamodb.New(sess)

	return svc

}

func (sdr ShippingAddressDynamoRepository) InsertShippingAddress(p domain.ShippingAddress) (string, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ShippingAddressRecord := toPersistedDynamodbEntitySA(p)
	av, err := dynamodbattribute.MarshalMap(ShippingAddressRecord)
	if err != nil {
		return "", &errs.AppError{Message: fmt.Sprintf("unable to marshal - %s", err.Error())}
	}
	fmt.Println(p)
	fmt.Println(ShippingAddressRecord)
	fmt.Println(av)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("ShippingAddress"),
	}

	_, err = sdr.Session.PutItemWithContext(ctx, input)

	if err != nil {
		return "", &errs.AppError{Message: fmt.Sprintf("unable to put the item - %s", err.Error())}
	}

	return ShippingAddressRecord.ID, nil
}

func (sdr ShippingAddressDynamoRepository) FindShippingAddressById(ShippingAddressID string) (*domain.ShippingAddress, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("ShippingAddress"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(ShippingAddressID),
			},
		},
	}

	result, err := sdr.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, &errs.AppError{Message: fmt.Sprintf("unable to get the item - %s", err.Error())}
	}

	if result.Item == nil {
		return nil, &errs.AppError{Message: fmt.Sprintf("unable to get the item - %s", err.Error())}
	}

	ShippingAddressModel := domain.ShippingAddress{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &ShippingAddressModel)

	if err != nil {
		return nil, &errs.AppError{Message: fmt.Sprintf("unmarshal map - %s", err.Error())}
	}

	return &ShippingAddressModel, nil
}

func (sdr ShippingAddressDynamoRepository) UpdateShippingAddressById(id string, sh domain.ShippingAddress) (bool, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(sh.FirstName),
			}, ":s1": {
				S: aws.String(sh.LastName),
			}, ":s2": {
				S: aws.String(sh.City),
			}, ":s3": {
				S: aws.String(sh.Address1),
			}, ":s4": {
				S: aws.String(sh.Address2),
			}, ":s5": {
				N: aws.String(strconv.Itoa(sh.CountryID)),
			}, ":s6": {
				N: aws.String(strconv.Itoa(sh.PostCode)),
			}, ":s7": {
				S: aws.String(time.Now().String()),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set firstname =:s, lastname = :s1, city = :s2, address_1 = :s3, address_2 = :s4, country_id = :s5, postcode =:s6, updated_at =:s7"),
		TableName:        aws.String("ShippingAddress"),
	}

	_, err := sdr.Session.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, &errs.AppError{Message: fmt.Sprintf("unable to update - %s", err.Error())}
	}
	return true, nil
}

func (sdr ShippingAddressDynamoRepository) DeleteShippingAddressById(id string) (bool, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String("ShippingAddress"),
	}

	_, err := sdr.Session.DeleteItemWithContext(ctx, input)
	if err != nil {
		return false, &errs.AppError{Message: fmt.Sprintf("unable to delete- %s", err.Error())}
	}
	return true, nil
}

func toPersistedDynamodbEntitySA(o domain.ShippingAddress) *ShippingAddressModel {
	return &ShippingAddressModel{
		ID:        uuid.New().String(),
		FirstName: o.FirstName,
		LastName:  o.LastName,
		City:      o.City,
		Address1:  o.Address1,
		Address2:  o.Address2,
		CountryID: o.CountryID,
		PostCode:  o.PostCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewDynamoShippingAddressRepository() ShippingAddressDynamoRepository {
	svc := connect()
	return ShippingAddressDynamoRepository{Session: svc, Tablename: "ShippingAddress"}
}
