package infra

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
)

type ShippingCostDynamoRepository struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}

func (sdr ShippingCostDynamoRepository) InsertShippingCost(p domain.ShippingCost) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ShippingCostRecord := toPersistedDynamodbEntitySC(p)
	av, err := dynamodbattribute.MarshalMap(ShippingCostRecord)
	if err != nil {
		return false, fmt.Errorf("unable to marshal - %s", err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("ShippingCost"),
	}

	_, err = sdr.Session.PutItemWithContext(ctx, input)

	if err != nil {
		return false, fmt.Errorf("unable to put the item - %s", err.Error())
	}

	return true, nil
}

func (sdr ShippingCostDynamoRepository) FindShippingCostByCity(city string) (*domain.ShippingCost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("ShippingCost"),
		Key: map[string]*dynamodb.AttributeValue{
			"city": {
				S: aws.String(city),
			},
		},
	}

	result, err := sdr.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("unable to get the item - %s", err.Error())
	}

	if result.Item == nil {
		return nil, fmt.Errorf("item not found")
	}

	ShippingCostModel := domain.ShippingCost{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &ShippingCostModel)

	if err != nil {
		return nil, fmt.Errorf("unmarshal map - %s", err.Error())
	}

	return &ShippingCostModel, nil
}

func (sdr ShippingCostDynamoRepository) UpdateShippingCost(sh domain.ShippingCost) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(strconv.Itoa(sh.ShippingCost)),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"city": {
				S: aws.String(sh.City),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set cost=:s"),
		TableName:        aws.String("ShippingCost"),
	}

	_, err := sdr.Session.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to update - %s", err.Error())
	}
	return true, nil
}

func (sdr ShippingCostDynamoRepository) DeleteShippingCostByCity(city string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"city": {
				S: aws.String(city),
			},
		},
		TableName: aws.String("ShippingCost"),
	}

	_, err := sdr.Session.DeleteItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to delete - %s", err.Error())
	}
	return true, nil
}

func toPersistedDynamodbEntitySC(o domain.ShippingCost) *ShippingCostModel {
	return &ShippingCostModel{
		City: o.City,
		Cost: o.ShippingCost,
	}
}

func NewShippingCostDynamoRepository() ShippingCostDynamoRepository {
	svc := connect()
	return ShippingCostDynamoRepository{Session: svc, Tablename: "ShippingCost"}
}
