package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"
	"time"
)

type UserDynamoDBRepository struct {
	tableName string
}

func NewDynamoDBRepository() UserDynamoDBRepository {
	return UserDynamoDBRepository{
		tableName: "users",
	}
}

// To use local dynamodb
// func createDynamoDBClient() *dynamodb.DynamoDB {
// 	// Create AWS Session
// 	sess, err := session.NewSession(&aws.Config{
// 		Region:   aws.String("ap-southeast-2"),
// 		Endpoint: aws.String("http://localhost:8000")})

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// Return DynamoDB client
// 	return dynamodb.New(sess)
// }

func createDynamoDBClient() *dynamodb.DynamoDB {
	// Create AWS Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Return DynamoDB client
	return dynamodb.New(sess)
}

func (repo UserDynamoDBRepository) Save(user domain.User) (domain.User, error) {
	dynamoDBClient := createDynamoDBClient()
	dynamodbUser := toPersistedDynamodbEntity(user)

	attributeValue, err := dynamodbattribute.MarshalMap(dynamodbUser)
	if err != nil {
		return domain.User{}, err
	}

	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(repo.tableName),
	}

	_, err = dynamoDBClient.PutItem(item)
	if err != nil {
		return domain.User{}, err
	}

	return user, err
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
// 	dynamoDBClient := createDynamoDBClient()

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

// func (repo *dynamoDBRepo) FindByID(id string) (*entity.Post, error) {
// 	// Get a new DynamoDB client
// 	dynamoDBClient := createDynamoDBClient()

// 	result, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
// 		TableName: aws.String(repo.tableName),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"id": {
// 				N: aws.String(id),
// 			},
// 		},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	post := entity.Post{}
// 	err = dynamodbattribute.UnmarshalMap(result.Item, &post)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &post, nil
// }

// // Delete: TODO
// func (repo *dynamoDBRepo) Delete(post *entity.Post) error {
// 	return nil
// }

