package app

import (
	"net/http"
	"os"

	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

type productFrontStoreHandler struct {
}

func (pah productFrontStoreHandler) GetAllProducts(c *gin.Context) {
	productFrontStoreServiceURI := os.Getenv("PRODUCT_FRONT_STORE_SERVICE_URI") + "/products"

	req, err := http.NewRequest("GET", productFrontStoreServiceURI, nil)
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

func (pah productFrontStoreHandler) GetProductByID(c *gin.Context) {
	productFrontStoreServiceURI := os.Getenv("PRODUCT_FRONT_STORE_SERVICE_URI") + "/products/" + c.Param("id")

	req, err := http.NewRequest("GET", productFrontStoreServiceURI, nil)
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

func (pah productFrontStoreHandler) GetProductsByCategory(c *gin.Context) {
	productFrontStoreServiceURI := os.Getenv("PRODUCT_FRONT_STORE_SERVICE_URI") + "/products/category/:id" + c.Param("id")

	req, err := http.NewRequest("GET", productFrontStoreServiceURI, nil)
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
