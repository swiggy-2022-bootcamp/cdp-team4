package infra

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
)

type PaymentDynamoRepository struct {
	Session *dynamodb.DynamoDB
}

func connect() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
	})

	if err != nil {
		panic(err.Error())
	}

	// create dynamo client
	svc := dynamodb.New(sess)

	return svc
}

func (pdr PaymentDynamoRepository) Insert(p domain.Payment) (bool, error) {
	return true, errors.New("dummy error!")
}

func (pdr PaymentDynamoRepository) FindById(string) (*domain.Payment, error) {
	return &domain.Payment{}, errors.New("dummy error!")

}

func (pdr PaymentDynamoRepository) FindByUserID(string) ([]*domain.Payment, error) {
	return []*domain.Payment{}, errors.New("dummy error!")

}
func (pdr PaymentDynamoRepository) UpdateStatus(string, string) (*domain.Payment, error) {
	return &domain.Payment{}, errors.New("dummy error!")

}
func (pdr PaymentDynamoRepository) DeleteByID(string) (bool, error) {
	return true, errors.New("dummy error!")

}

func NewDynamoRepository() PaymentDynamoRepository {
	svc := connect()
	return PaymentDynamoRepository{Session: svc}
}

// func (dyr PaymentDynamoRepository) ListTables() ([]string, error) {
// 	input := &dynamodb.ListTablesInput{}

// 	// if docker container of dynamoDB is not running then code
// 	// should be blocked indefinitely, that's why using context with time out.
// 	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	var tableNames []string

// 	result, err := dyr.Session.ListTablesWithContext(context, input)
// 	if err != nil {
// 		log.Fatal(err)
// 		return tableNames, err
// 	}

// 	for _, n := range result.TableNames {
// 		tableNames = append(tableNames, *n)
// 	}

// 	return tableNames, nil
// }

// func (dyr PaymentDynamoRepository) CreateTable(tableName string) bool {
// 	// input := &dynamodb.CreateTableInput{}

// 	// context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancel()
// 	return true
// }
