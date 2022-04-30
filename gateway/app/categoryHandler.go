package app

import (
	"net/http"
	"os"

	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
}

func (ch categoryHandler) GetAllCategories(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	categoryServiceURI := os.Getenv("CATEGORY_SERVICE_URI") + "/categories/"

	req, err := http.NewRequest("GET", categoryServiceURI, nil)
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	var categoryResponseDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&categoryResponseDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(categoryResponseDTO.Status, categoryResponseDTO)
}

func (ch categoryHandler) AddCategory(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	productAdminServiceURI := os.Getenv("CATEGORY_SERVICE_URI") + "/categories/"

	var newCategory CategoryDTO
	if err := c.BindJSON(&newCategory); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	productJSON, err1 := json.Marshal(newCategory)
	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("POST", productAdminServiceURI, bytes.NewBuffer(productJSON))
	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var categoryResponseDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&categoryResponseDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(categoryResponseDTO.Status, categoryResponseDTO)
}

func (ch categoryHandler) GetCategoryByID(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	categoryServiceURI := os.Getenv("CATEGORY_SERVICE_URI") + "/categories/:id" + c.Param("id")

	req, err := http.NewRequest("GET", categoryServiceURI, nil)
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()
	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var categoryResponseDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&categoryResponseDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(categoryResponseDTO.Status, categoryResponseDTO)
}

func (ch categoryHandler) UpdateCategoryByID(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	categoryServiceURI := os.Getenv("CATEGORY_SERVICE_URI") + "/categories/:id" + c.Param("id")
	var newProduct ProductAdminDTO
	if err := c.BindJSON(&newProduct); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("here1: %v", newProduct)

	productJSON, err1 := json.Marshal(newProduct)
	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}
	req, err := http.NewRequest("PUT", categoryServiceURI, bytes.NewBuffer(productJSON))
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var productResponseDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&productResponseDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(productResponseDTO.Status, productResponseDTO)
}

func (ch categoryHandler) DeleteCategoryByID(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	categoryServiceURI := os.Getenv("CATEGORY_SERVICE_URI") + "/categories/:id" + c.Param("id")

	req, err := http.NewRequest("DELETE", categoryServiceURI, nil)
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()
	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var categoryResponseDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&categoryResponseDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(categoryResponseDTO.Status, categoryResponseDTO)
	c.Abort()
}

func (ch categoryHandler) DeleteCategories(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	categoryServiceURI := os.Getenv("CATEGORY_SERVICE_URI") + "/categories"

	var categories []string
	if err := c.BindJSON(&categories); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	productJSON, err1 := json.Marshal(categories)
	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err := http.NewRequest("DELETE", categoryServiceURI, bytes.NewBuffer(productJSON))
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()
	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var productResponseDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&productResponseDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(productResponseDTO.Status, productResponseDTO)
}

//================DTO==================
type CategoryDescriptionDTO struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	MetaDescription string `json:"meta_description"`
	MetaKeyword     string `json:"meta_keyword"`
	MetaTitle       string `json:"meta_title"`
}

type CategoryDTO struct {
	Id                  string                   `json:"id"`
	CategoryDescription []CategoryDescriptionDTO `json:"category_description"`
}
