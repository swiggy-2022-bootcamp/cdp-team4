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

type productAdminHandler struct {
}

func (pah productAdminHandler) GetAllProducts(c *gin.Context) {
	productAdminServiceURI := os.Getenv("PRODUCT_ADMIN_SERVICE_URI") + "/products"

	req, err := http.NewRequest("GET", productAdminServiceURI, nil)
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

func (pah productAdminHandler) AddProduct(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	productAdminServiceURI := os.Getenv("PRODUCT_ADMIN_SERVICE_URI") + "/products/"

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

	var productResponseDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&productResponseDTO)

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

func (pah productAdminHandler) GetProductByID(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	productAdminServiceURI := os.Getenv("PRODUCT_ADMIN_SERVICE_URI") + "/products/" + c.Param("id")

	req, err := http.NewRequest("GET", productAdminServiceURI, nil)
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

func (pah productAdminHandler) UpdateProductByID(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	productAdminServiceURI := os.Getenv("PRODUCT_ADMIN_SERVICE_URI") + "/products/" + c.Param("id")
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
	req, err := http.NewRequest("PUT", productAdminServiceURI, bytes.NewBuffer(productJSON))
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

func (pah productAdminHandler) DeleteProductByID(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	productAdminServiceURI := os.Getenv("PRODUCT_ADMIN_SERVICE_URI") + "/products/" + c.Param("id")

	req, err := http.NewRequest("DELETE", productAdminServiceURI, nil)
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
	c.Abort()
}

func (pah productAdminHandler) SearchByCategoryID(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	productAdminServiceURI := os.Getenv("PRODUCT_ADMIN_SERVICE_URI") + "/products/search/category/" + c.Param("categoryid")

	req, err := http.NewRequest("GET", productAdminServiceURI, nil)
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

func (pah productAdminHandler) SearchByManufacturerID(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	productAdminServiceURI := os.Getenv("PRODUCT_ADMIN_SERVICE_URI") + "/products/search/manufacturer/" + c.Param("id")

	req, err := http.NewRequest("GET", productAdminServiceURI, nil)
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

func (pah productAdminHandler) SearchByKeyword(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	productAdminServiceURI := os.Getenv("PRODUCT_ADMIN_SERVICE_URI") + "/products/search/keyword/" + c.Param("keyword")

	req, err := http.NewRequest("GET", productAdminServiceURI, nil)
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
type ProductSEOURLDTO struct {
	Keyword    string `json:"keyword"`
	LanguageID string `json:"language_id"`
	StoreID    string `json:"store_id"`
}

type ProductDescriptionDTO struct {
	LanguageID      string `json:"language_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeyword     string `json:"meta_keyword"`
	Tag             string `json:"tag"`
}

type ProductAdminDTO struct {
	Model               string                  `json:"model"`
	Quantity            int64                   `json:"quantity"`
	Price               float64                 `json:"price"`
	ManufacturerID      string                  `json:"manufacturer_id"`
	SKU                 string                  `json:"sku"`
	ProductSEOURLs      []ProductSEOURLDTO      `json:"product_seo_url"`
	Points              int64                   `json:"points"`
	Reward              int64                   `json:"reward"`
	ImageURL            string                  `json:"image_url"`
	IsShippable         bool                    `json:"is_shippable"`
	Weight              float64                 `json:"weight"`
	Length              float64                 `json:"length"`
	Width               float64                 `json:"width"`
	Height              float64                 `json:"height"`
	MinimumQuantity     int64                   `json:"minimum_quantity"`
	RelatedProducts     []string                `json:"related_products"`
	ProductDescriptions []ProductDescriptionDTO `json:"product_description"`
	ProductCategories   []string                `json:"product_categories"`
}
