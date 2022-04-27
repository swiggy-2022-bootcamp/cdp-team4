package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/domain"
)

type CategoryDynamoRepository struct {
	Session *dynamodb.DynamoDB
}

func connect() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// create dynamo client
	svc := dynamodb.New(sess)
	return svc
}

func (categoryRepo CategoryDynamoRepository) InsertCategory(category domain.Category) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	categoryRecord := _toDynamoCategoryModel(&category)
	av, err := dynamodbattribute.MarshalMap(categoryRecord)
	if err != nil {
		return false, fmt.Errorf("unable to marshal - %s", err.Error())
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Category"),
	}

	_, err = categoryRepo.Session.PutItemWithContext(ctx, input)

	if err != nil {
		return false, fmt.Errorf("unable to insert the item - %s", err.Error())
	}
	return true, nil
}

func (categoryRepo CategoryDynamoRepository) FindAllCategories() ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.ScanInput{
		TableName: aws.String("Category"),
	}
	result, err := categoryRepo.Session.ScanWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	// Make the DynamoDB Query API call
	var categories = []domain.Category{}
	for _, item := range result.Items {
		category := domain.Category{}
		if err := dynamodbattribute.UnmarshalMap(item, &category); err != nil {
			return []domain.Category{}, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (categoryRepo CategoryDynamoRepository) FindCategoryByID(categoryId string) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Category"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(categoryId),
			},
		},
	}

	result, err := categoryRepo.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return domain.Category{}, fmt.Errorf("unable to get the item - %s", err.Error())
	}

	if result.Item == nil {
		return domain.Category{}, fmt.Errorf("item not found")
	}

	categoryModel := domain.Category{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &categoryModel)
	if err != nil {
		return domain.Category{}, fmt.Errorf("unmarshal map - %s", err.Error())
	}
	return categoryModel, nil
}

func (categoryRepo CategoryDynamoRepository) UpdateCategoryById(category domain.Category) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	prevCategoryInput := &dynamodb.GetItemInput{
		TableName: aws.String("Category"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(category.Id),
			},
		},
	}
	prevCategoryResult, err := categoryRepo.Session.GetItemWithContext(ctx, prevCategoryInput)
	if err != nil {
		return false, err
	}
	oldCategory := domain.Category{}
	err = dynamodbattribute.UnmarshalMap(prevCategoryResult.Item, &oldCategory)

	if err != nil {
		return false, fmt.Errorf("unmarshal map - %s", err.Error())
	}

	if category.CategoryDescription[0].Description == "" {
		category.CategoryDescription[0].Description = oldCategory.CategoryDescription[0].Description
	}
	if category.CategoryDescription[0].Name == "" {
		category.CategoryDescription[0].Name = oldCategory.CategoryDescription[0].Name
	}
	if category.CategoryDescription[0].MetaDescription == "" {
		category.CategoryDescription[0].MetaDescription = oldCategory.CategoryDescription[0].MetaDescription
	}
	if category.CategoryDescription[0].MetaKeyword == "" {
		category.CategoryDescription[0].MetaKeyword = oldCategory.CategoryDescription[0].MetaKeyword
	}
	if category.CategoryDescription[0].MetaTitle == "" {
		category.CategoryDescription[0].MetaTitle = oldCategory.CategoryDescription[0].MetaTitle
	}
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(category.CategoryDescription[0].Name),
			}, ":s1": {
				S: aws.String(category.CategoryDescription[0].Description),
			}, ":s2": {
				S: aws.String(category.CategoryDescription[0].MetaDescription),
			}, ":s3": {
				S: aws.String(category.CategoryDescription[0].MetaKeyword),
			}, ":s4": {
				S: aws.String(category.CategoryDescription[0].MetaTitle),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(category.Id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set category_description.name =:s, category_description.description = :s1, category_description.mata_description = :s2,  category_description.mata_keyword= :s3, category_description.meta_title = :s4"),
		TableName:        aws.String("Category"),
	}
	_, err = categoryRepo.Session.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, err
	}
	return true, err
}

func (categoryRepo CategoryDynamoRepository) DeleteCategories(catregoryIds []string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, categoryId := range catregoryIds {
		var queryInput = &dynamodb.QueryInput{
			TableName: aws.String("ProductCategoryRelation"),
			KeyConditions: map[string]*dynamodb.Condition{
				"category_id": {
					ComparisonOperator: aws.String("EQ"),
					AttributeValueList: []*dynamodb.AttributeValue{
						{
							S: aws.String(categoryId),
						},
					},
				},
			},
		}
		var resp, err = categoryRepo.Session.Query(queryInput)
		if err != nil {
			return false, err
		}
		if resp != nil {
			return false, fmt.Errorf("unable to delete - %s", err.Error())
		}
		//delete the category
		input := &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					N: aws.String(categoryId),
				},
			},
			TableName: aws.String("Category"),
		}

		_, err = categoryRepo.Session.DeleteItemWithContext(ctx, input)
		if err != nil {
			return false, fmt.Errorf("unable to delete - %s", err.Error())
		}
	}
	return true, nil
}

func (categoryRepo CategoryDynamoRepository) DeleteCategoryById(categoryId string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//check if category is associated with products
	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String("ProductCategoryRelation"),
		KeyConditions: map[string]*dynamodb.Condition{
			"category_id": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(categoryId),
					},
				},
			},
		},
	}
	var resp, err = categoryRepo.Session.Query(queryInput)
	if err != nil {
		return false, err
	}
	if resp != nil {
		return false, fmt.Errorf("unable to delete - %s", err.Error())
	}
	//delete the category
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(categoryId),
			},
		},
		TableName: aws.String("Category"),
	}

	_, err = categoryRepo.Session.DeleteItemWithContext(ctx, input)
	if err != nil {
		return false, fmt.Errorf("unable to delete - %s", err.Error())
	}
	return true, nil
}

func NewDynamoRepository() CategoryDynamoRepository {
	svc := connect()
	return CategoryDynamoRepository{Session: svc}
}

func _toDynamoCategoryModel(category *domain.Category) CategoryModel {
	var categoryDescriptionModel []CategoryDescriptionModel
	for _, item := range category.CategoryDescription {
		categoryDescriptionModel = append(categoryDescriptionModel, CategoryDescriptionModel{Name: item.Name, Description: item.Description,
			MetaDescription: item.MetaDescription, MetaKeyword: item.MetaKeyword, MetaTitle: item.MetaTitle})
	}
	return CategoryModel{
		Id:                  category.Id,
		CategoryDescription: categoryDescriptionModel,
	}
}
