package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/domain"
)

type CategoryHandler struct {
	CategoryService domain.CategoryService
}

// Add Category
// @Summary      Add Category
// @Description  This Handle allows admin to create a new category
// @Tags         Category
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /categories/    [post]
func (categoryHandler CategoryHandler) HandleAddCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var category domain.CategoryDescription
		if err := ctx.BindJSON(&category); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		result, err := categoryHandler.CategoryService.AddCategory(category)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"category": result})
	}
}

// Get all categories
// @Summary      Get all Categories
// @Description  This Handle allows admin to fetch all the categories in the datastore
// @Tags         Categories
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /categories/    [get]
func (categoryHandler CategoryHandler) HandleGetAllCategories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categories, err := categoryHandler.CategoryService.GetCategories()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"categories": categories})
	}
}

// Get category by ID
// @Summary      Get Category details by Id
// @Description  This Handle allows admin to get a category, given Id
// @Tags         Category
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /categories/:id    [get]
func (categoryHandler CategoryHandler) HandleGetCategoryByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryId := ctx.Param("id")
		category, err := categoryHandler.CategoryService.GetCategoryById(categoryId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"category": category})
	}
}

// Update category by ID
// @Summary      Update Category details by Id
// @Description  This Handle allows admin to update a category, given Id
// @Tags         Category
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /categories/:id    [put]
func (categoryHandler CategoryHandler) HandleUpdateCategoryByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryId := ctx.Param("id")
		var categoryDetails []domain.CategoryDescription
		if err := ctx.BindJSON(&categoryDetails[0]); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		category := domain.Category{Id: categoryId, CategoryDescription: categoryDetails}
		_, err := categoryHandler.CategoryService.UpdateCategoryById(category)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "category updated successfully"})
	}
}

// Delete category by ID
// @Summary      Delete Category details by Id
// @Description  This Handle allows admin to delete a category, given Id
// @Tags         Category
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /categories/:id    [delete]
func (categoryHandler CategoryHandler) HandleDeleteCategoryByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryId := ctx.Param("id")
		_, err := categoryHandler.CategoryService.DeleteCategoryById(categoryId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Category deleted successfully"})
	}
}

// Delete categories
// @Summary      Delete Categories details
// @Description  This Handle allows admin to delete categories, given Ids
// @Tags         Category
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /categories/    [delete]
func (categoryHandler CategoryHandler) HandleDeleteCategories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var categoryList []string
		if err := ctx.BindJSON(&categoryList); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		_, err := categoryHandler.CategoryService.DeleteCategories(categoryList)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "categories deleted successfully"})
	}
}
