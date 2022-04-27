package infra

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/utils/errs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
)

type RewardDynamoRepository struct {
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
		TableName: aws.String("Reward"),
	}

	_, tableErr := svc.CreateTable(tableInput)

	if tableErr != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(tableErr.Error())
	}
}

func connect() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// create dynamo client
	svc := dynamodb.New(sess)

	createTable(svc)
	
	return svc
}

func (rwd RewardDynamoRepository) InsertReward(p domain.Reward) (string, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rewardRecord := toPersistedDynamodbEntity(p)

	av, err := dynamodbattribute.MarshalMap(rewardRecord)
	if err != nil {
		errstring := fmt.Sprintf("unable to marshal - %s", err.Error())
		return "", &errs.AppError{Message: errstring}
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Reward"),
	}

	_, err = rwd.Session.PutItemWithContext(ctx, input)

	if err != nil {
		errstring := fmt.Sprintf("unable to put item - %s", err.Error())
		return "", &errs.AppError{Message: errstring}
	}

	return rewardRecord.Id, nil
}

func (rwd RewardDynamoRepository) FindRewardById(rewardID string) (*domain.Reward, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Reward"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(rewardID),
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

	rewardModel := RewardModel{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &rewardModel)

	if err != nil {
		errstring := fmt.Sprintf("unmarshal map - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}
	rwdModel := toModelfromDynamodbEntity(rewardModel)
	return rwdModel, nil
}

func (rwd RewardDynamoRepository) FindRewardByUserId(userId string) (*domain.Reward, *errs.AppError) {
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
		TableName:                 aws.String("Reward"),
	}

	result, err := rwd.Session.ScanWithContext(ctx, input)
	if err != nil {
		errstring := fmt.Sprintf("scan with filter - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}

	if len(result.Items) == 0 {
		return nil, &errs.AppError{Message: "Item not found"}
	}

	record := RewardModel{}
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &record)
	if err != nil {
		fmt.Println("Stray records inside db")
	}

	rwdModel := toModelfromDynamodbEntity(record)

	return rwdModel, nil
}

func (rwd RewardDynamoRepository) UpdateRewardByUserId(userId string, points int) (bool, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//find the current record
	currentReward, err := rwd.FindRewardByUserId(userId)
	if err!=nil{
		newDomainRecord := domain.NewReward(userId,points)
		_,err=rwd.InsertReward(*newDomainRecord)
		if err != nil {
			return false,&errs.AppError{Message: "Unable to update"}
		}
		return true,nil
	}

	updatedRewardPoints := currentReward.RewardPoints + points
	if updatedRewardPoints<0 {
		return false, &errs.AppError{Message:"Reward point shortage"}
	}
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				N: aws.String(strconv.Itoa(updatedRewardPoints)),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(currentReward.Id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set reward_points = :s"),
		TableName:        aws.String("Reward"),
	}

	_, updError := rwd.Session.UpdateItemWithContext(ctx, input)
	if updError  != nil {
		return false, &errs.AppError{Message: updError.Error()+"Unable to update"}
	}
	fmt.Println("Updated exisitng user with points",updatedRewardPoints)
	return true, nil
}

// func (odr RewardDynamoRepository) FindAllRewardUserId(userId string) {
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
// 		TableName:                 aws.String("Reward"),
// 	}

// 	result, err := odr.Session.ScanWithContext(ctx, input)
// 	if err != nil {
// 		errstring := fmt.Sprintf("scan with filter - %s", err.Error())
// 		fmt.Println(&errs.AppError{Message: errstring})
// 	}

	

// 	for _, item := range result.Items {
// 		record := RewardModel{}
// 		err := dynamodbattribute.UnmarshalMap(item, &record)
// 		if err != nil {
// 			fmt.Println("Stray records inside db")
// 		} else {
// 			recordm := toModelfromDynamodbEntity(record)
// 			fmt.Println(recordm)
// 		}
// 	}
// }

func toPersistedDynamodbEntity(r domain.Reward) *RewardModel {
	return &RewardModel{
		Id:           uuid.New().String(),
		UserID:       r.UserID,
		RewardPoints: r.RewardPoints,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func toModelfromDynamodbEntity(r RewardModel) *domain.Reward {
	return &domain.Reward{
		Id:           r.Id,
		UserID:       r.UserID,
		RewardPoints: r.RewardPoints,
	}
}

func NewDynamoRepository() RewardDynamoRepository {
	svc := connect()
	return RewardDynamoRepository{Session: svc, Tablename: "Reward"}
}
