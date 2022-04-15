package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"
	"time"
	"context"
	"fmt"
)

type UserDynamoDBRepository struct {
	Session   *dynamodb.DynamoDB
	TableName string
}

func NewDynamoRepository() UserDynamoDBRepository {
	svc := connect()
	return UserDynamoDBRepository{Session: svc, TableName: "users"}
}


func connect() *dynamodb.DynamoDB {
	// Create AWS Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Return DynamoDB client
	return dynamodb.New(sess)
}

func (repo UserDynamoDBRepository) Save(user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dynamodbUser := toPersistedDynamodbEntity(user)

	attributeValue, err := dynamodbattribute.MarshalMap(dynamodbUser)
	if err != nil {
		return domain.User{}, err
	}

	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(repo.TableName),
	}

	_, err = repo.Session.PutItemWithContext(ctx, item)

	if err != nil {
		return domain.User{}, err
	}

	return user, err
}

func (repo UserDynamoDBRepository) FindByID(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String(repo.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
	}

	result, err := repo.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, fmt.Errorf("item not found")
	}

	user := domain.User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)

	if err != nil {
		return nil, fmt.Errorf("unmarshal map - %s", err.Error())
	}

	return &user, nil
}

func toPersistedDynamodbEntity(u domain.User) *UserModel {
	return &UserModel{
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Phone:           u.Phone,
		Email:           u.Email,
		Username:        u.Username,
		Password:        u.Password,
		Role:            u.Role,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// func (repo *dynamoDBRepo) FindAll() ([]entity.Post, error) {
// 	// Get a new DynamoDB client
// 	dynamoDBClient := connect()

// 	// Build the query input parameters
// 	params := &dynamodb.ScanInput{
// 		TableName: aws.String(repo.tableName),
// 	}

// 	// Make the DynamoDB Query API call
// 	result, err := dynamoDBClient.Scan(params)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var posts []entity.Post = []entity.Post{}
// 	for _, i := range result.Items {
// 		post := entity.Post{}

// 		err = dynamodbattribute.UnmarshalMap(i, &post)

// 		if err != nil {
// 			panic(err)
// 		}
// 		posts = append(posts, post)
// 	}
// 	return posts, nil
// }



// // Delete: TODO
// func (repo *dynamoDBRepo) Delete(post *entity.Post) error {
// 	return nil
// }

