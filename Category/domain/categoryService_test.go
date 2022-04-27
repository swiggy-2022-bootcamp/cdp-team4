package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/domain"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/Category/mocks"
)

var mockCategoryRepo = mocks.CategoryDynamoRepository{}
var categoryService = domain.NewCategoryService(&mockCategoryRepo)

func TestShouldReturnNewCategoryService(t *testing.T) {
	categoryService := domain.NewCategoryService(nil)
	assert.NotNil(t, categoryService)
}

func TestShouldCreateNewCategory(t *testing.T) {
	categoryDescription := domain.CategoryDescription{Name: "Earphones",
		Description:     "earphones",
		MetaDescription: "meta description for earphones",
		MetaKeyword:     "earphones, earbuds",
		MetaTitle:       "gadgets"}
	categoryDescriptions := []domain.CategoryDescription{categoryDescription}
	_ = domain.NewCategoryObject(categoryDescriptions)

	mockCategoryRepo.On("InsertCategory", mock.Anything).Return(true, nil)
	categoryService.AddCategory(categoryDescription)
	mockCategoryRepo.AssertNumberOfCalls(t, "InsertCategory", 1)
}

func TestGetCategoryById(t *testing.T) {
	categoryDescription := domain.CategoryDescription{Name: "Earphones",
		Description:     "earphones",
		MetaDescription: "meta description for earphones",
		MetaKeyword:     "earphones, earbuds",
		MetaTitle:       "gadgets"}
	categoryDescriptions := []domain.CategoryDescription{categoryDescription}
	newCategory := domain.NewCategoryObject(categoryDescriptions)
	mockCategoryRepo.On("FindCategoryByID", newCategory.Id).Return(*newCategory, nil)
	resCategory, _ := categoryService.GetCategoryById(newCategory.Id)
	assert.Equal(t, categoryDescription.Name, resCategory.CategoryDescription[0].Name)
	assert.Equal(t, categoryDescription.Description, resCategory.CategoryDescription[0].Description)
	assert.Equal(t, categoryDescription.MetaDescription, resCategory.CategoryDescription[0].MetaDescription)
	assert.Equal(t, categoryDescription.MetaKeyword, resCategory.CategoryDescription[0].MetaKeyword)
	assert.Equal(t, categoryDescription.MetaTitle, resCategory.CategoryDescription[0].MetaTitle)
}

func TestGetAllCategories(t *testing.T) {
	categoryDescription := domain.CategoryDescription{Name: "Earphones",
		Description:     "earphones",
		MetaDescription: "meta description for earphones",
		MetaKeyword:     "earphones, earbuds",
		MetaTitle:       "gadgets"}
	categoryDescriptions := []domain.CategoryDescription{categoryDescription}
	newCategory := domain.NewCategoryObject(categoryDescriptions)
	mockCategoryRepo.On("FindAllCategories").Return([]domain.Category{*newCategory}, nil)
	resCategory, _ := categoryService.GetCategories()
	assert.Equal(t, categoryDescription.Name, resCategory[0].CategoryDescription[0].Name)
	assert.Equal(t, categoryDescription.Description, resCategory[0].CategoryDescription[0].Description)
	assert.Equal(t, categoryDescription.MetaDescription, resCategory[0].CategoryDescription[0].MetaDescription)
	assert.Equal(t, categoryDescription.MetaKeyword, resCategory[0].CategoryDescription[0].MetaKeyword)
	assert.Equal(t, categoryDescription.MetaTitle, resCategory[0].CategoryDescription[0].MetaTitle)
}

/*
DeleteCategories([]string) (bool, error)*/
func TestUpdateCategoryById(t *testing.T) {
	categoryDescription := domain.CategoryDescription{Name: "Earphones",
		Description:     "earphones",
		MetaDescription: "meta description for earphones",
		MetaKeyword:     "earphones, earbuds",
		MetaTitle:       "gadgets"}
	categoryDescriptions := []domain.CategoryDescription{categoryDescription}
	newCategory := domain.NewCategoryObject(categoryDescriptions)
	mockCategoryRepo.On("UpdateCategoryById", *newCategory).Return(true, nil)
	res, err := categoryService.UpdateCategoryById(*newCategory)
	assert.Nil(t, err)
	assert.Equal(t, res, true)
}
func TestDeleteCategories(t *testing.T) {
	categories := []string{"10293194182", "10293194183"}
	mockCategoryRepo.On("DeleteCategories", categories).Return(true, nil)
	res, err := categoryService.DeleteCategories(categories)
	assert.Nil(t, err)
	assert.Equal(t, res, true)
}
func TestDeleteCategoryByCategoryID(t *testing.T) {
	categoryId := "10293194182"
	mockCategoryRepo.On("DeleteCategoryById", categoryId).Return(true, nil)
	res, err := categoryService.DeleteCategoryById(categoryId)
	assert.Nil(t, err)
	assert.Equal(t, res, true)
}
