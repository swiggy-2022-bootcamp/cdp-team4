package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryDescription struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	MetaDescription string `json:"meta_description"`
	MetaKeyword     string `json:"meta_keyword"`
	MetaTitle       string `json:"meta_title"`
}

type Category struct {
	Id                  string                `json:"id"`
	CategoryDescription []CategoryDescription `json:"category_description"`
}

type CategoryDynamoRepository interface {
	InsertCategory(Category) (bool, error)
	FindAllCategories() ([]Category, error)
	FindCategoryByID(string) (Category, error)
	UpdateCategoryById(Category) (bool, error)
	DeleteCategories([]string) (bool, error)
	DeleteCategoryById(string) (bool, error)
}

func GenerateUniqueId() string {
	return primitive.NewObjectID().Hex()
}

func NewCategoryObject(categoryDescription []CategoryDescription) *Category {
	return &Category{
		Id:                  GenerateUniqueId(),
		CategoryDescription: categoryDescription,
	}
}
