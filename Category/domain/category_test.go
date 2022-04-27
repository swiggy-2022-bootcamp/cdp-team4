package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/domain"
)

func TestShouldReturnNewCategory(t *testing.T) {
	categoryDescription := domain.CategoryDescription{Name: "Earphones",
		Description:     "earphones",
		MetaDescription: "meta description for earphones",
		MetaKeyword:     "earphones, earbuds",
		MetaTitle:       "gadgets"}
	categoryDescriptions := []domain.CategoryDescription{categoryDescription}
	newCategory := domain.NewCategoryObject(categoryDescriptions)
	assert.NotEmpty(t, newCategory.Id)
	assert.Equal(t, categoryDescription.Name, newCategory.CategoryDescription[0].Name)
	assert.Equal(t, categoryDescription.Description, newCategory.CategoryDescription[0].Description)
	assert.Equal(t, categoryDescription.MetaDescription, newCategory.CategoryDescription[0].MetaDescription)
	assert.Equal(t, categoryDescription.MetaKeyword, newCategory.CategoryDescription[0].MetaKeyword)
	assert.Equal(t, categoryDescription.MetaTitle, newCategory.CategoryDescription[0].MetaTitle)
}
