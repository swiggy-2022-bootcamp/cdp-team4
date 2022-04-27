package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	// "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/utils/errs"
)

type CartDynamoRepository struct {
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
		TableName: aws.String("Cart"),
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

func (crt CartDynamoRepository) InsertCart(p domain.Cart) (string, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cartRecord := toPersistedDynamodbEntity(p)

	av, err := dynamodbattribute.MarshalMap(cartRecord)
	if err != nil {
		errstring := fmt.Sprintf("unable to marshal - %s", err.Error())
		return "", &errs.AppError{Message: errstring}
	}
    
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Cart"),
	}

	_ , err = crt.Session.PutItemWithContext(ctx, input)

	if err != nil {
		errstring := fmt.Sprintf("unable to put item - %s", err.Error())
		return "", &errs.AppError{Message: errstring}
	}

	return cartRecord.Id, nil
}

func (crt CartDynamoRepository) UpdateCartByUserId(userId string,productList map[string]domain.Item) (bool, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//find the current record
	currentCart, err := crt.FindCartByUserId(userId)
	if err!=nil{
		newDomainRecord := domain.NewCart(userId,productList)
		_,err=crt.InsertCart(*newDomainRecord)
		if err != nil {
			return false,&errs.AppError{Message: "Unable to update"}
		}
		return true,nil
	}

	for key,value :=range productList{
		currentCart.Items[key]=value
	}

	cartRecord := toPersistedDynamodbEntity(*currentCart)
	cartRecord.Id = currentCart.Id
	av, errMarshal := dynamodbattribute.MarshalMap(cartRecord)
	if errMarshal != nil {
		errstring := fmt.Sprintf("unable to marshal - %s", errMarshal.Error())
		return false, &errs.AppError{Message: errstring}
	}
    
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Cart"),
	}

	_ , putErr := crt.Session.PutItemWithContext(ctx, input)

	if putErr != nil {
		errstring := fmt.Sprintf("unable to put item - %s", putErr.Error())
		return false, &errs.AppError{Message: errstring}
	}

	return true, nil
}

func (crt CartDynamoRepository) DeleteCartItemByUserId(userId string,productIds[]string) (bool, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//find the current record
	currentCart, err := crt.FindCartByUserId(userId)
	if err!=nil{
		return false,&errs.AppError{Message: "Cart for UserID Doesnt exist"}
	}

	for _,key :=range productIds{
		if _, ok := currentCart.Items[key]; ok {
			delete(currentCart.Items, key)
		}
	}

	cartRecord := toPersistedDynamodbEntity(*currentCart)
	cartRecord.Id = currentCart.Id
	av, errMarshal := dynamodbattribute.MarshalMap(cartRecord)
	if errMarshal != nil {
		errstring := fmt.Sprintf("unable to marshal - %s", errMarshal.Error())
		return false, &errs.AppError{Message: errstring}
	}
    
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Cart"),
	}

	_ , putErr := crt.Session.PutItemWithContext(ctx, input)

	if putErr != nil {
		errstring := fmt.Sprintf("unable to put item - %s", putErr.Error())
		return false, &errs.AppError{Message: errstring}
	}

	return true, nil
}

func (crt CartDynamoRepository) FindAllCarts() ([]domain.Cart, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	input := &dynamodb.ScanInput{
		TableName: aws.String("Cart"),
	}
	out, err := crt.Session.ScanWithContext(ctx, input)

	if err != nil {
		return nil, &errs.AppError{Message: err.Error()}
	}

	cartRecords := make([]domain.Cart, 0)

	for _, item := range out.Items {
		record := CartModel{}
		err := dynamodbattribute.UnmarshalMap(item, &record)
		if err != nil {
			errstring := fmt.Sprintf("expression new builder - %s", err.Error())
			return nil, &errs.AppError{Message: errstring}
		}
		recordm := toModelfromDynamodbEntity(record)
		cartRecords = append(cartRecords, *recordm)
	}

	return cartRecords, nil
}

func (crt CartDynamoRepository) FindCartById(cartId string) (*domain.Cart, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Cart"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(cartId),
			},
		},
	}

	result, err := crt.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, &errs.AppError{Message: "Failed to read"}
	}

	if result.Item == nil {
		return nil, &errs.AppError{Message: "Item not found"}
	}

	cartModel := CartModel{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &cartModel)
	if err != nil {
		errstring := fmt.Sprintf("unmarshal map - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}
	crtModel := toModelfromDynamodbEntity(cartModel)
	return crtModel, nil
}

func (crt CartDynamoRepository) FindCartByUserId(userId string) (*domain.Cart, *errs.AppError) {
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
		TableName:                 aws.String("Cart"),
	}

	result, err := crt.Session.ScanWithContext(ctx, input)
	if err != nil {
		errstring := fmt.Sprintf("scan with filter - %s", err.Error())
		return nil, &errs.AppError{Message: errstring}
	}

	if len(result.Items) == 0 {
		return nil, &errs.AppError{Message: "Item not found"}
	}

	record := CartModel{}
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &record)
	if err != nil {
		fmt.Println("Stray records inside db")
	}

	crtModel := toModelfromDynamodbEntity(record)

	return crtModel, nil
}

func (crt CartDynamoRepository) DeleteCartByUserId(userId string) (bool, *errs.AppError) {
	currentCart,err := crt.FindCartByUserId(userId)
	if err != nil {
		return false,&errs.AppError{Message: "Unable to Find cart with given UserId"}
	}
	_ , err = crt.DeleteCartById(currentCart.Id)
	if err != nil {
		return false,&errs.AppError{Message: "Unable to delete cart with Given userId"}
	}
	return true ,nil
}

func (crt CartDynamoRepository) DeleteCartById(cartId string) (bool, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(cartId),
			},
		},
		TableName: aws.String("Cart"),
	}

	_, err := crt.Session.DeleteItemWithContext(ctx, input)
	if err != nil {
		errstring := fmt.Sprintf("unable to delete - %s", err.Error())
		return false, &errs.AppError{Message: errstring}
	}
	return true, nil
}

func toPersistedDynamodbEntity(c domain.Cart) *CartModel {
	ItemMap:=map[string]Item{}
	for key, DomainItem := range c.Items {
		var singleItem Item
		singleItem.Cost = DomainItem.Cost
		singleItem.Name = DomainItem.Name
		singleItem.Quantity = DomainItem.Quantity
		ItemMap[key] = singleItem
	}

	return &CartModel{
		Id:        uuid.New().String(),
		UserID:    c.UserID,
		Items:     ItemMap,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func toModelfromDynamodbEntity(c CartModel) *domain.Cart {
	DomainItemMap:= map[string]domain.Item{}
	for key, ModelItem := range c.Items {
		var singleItem domain.Item
		singleItem.Cost = ModelItem.Cost
		singleItem.Name = ModelItem.Name
		singleItem.Quantity = ModelItem.Quantity
		DomainItemMap[key] = singleItem
	}

	return &domain.Cart{
		Id:     c.Id,
		UserID: c.UserID,
		Items:  DomainItemMap,
	}
}

func NewDynamoRepository() CartDynamoRepository {
	svc := connect()
	return CartDynamoRepository{Session: svc, Tablename: "Cart"}
}
