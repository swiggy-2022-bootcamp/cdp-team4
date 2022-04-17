package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/domain"
)

type ProductFrontStoreHandler struct {
	ProductFrontStoreService domain.ProductFrontStoreService
}

func (prodHandler ProductFrontStoreHandler) HandleGetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := prodHandler.ProductFrontStoreService.GetProducts()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"products": products})
	}
}

func (prodHandler ProductFrontStoreHandler) HandleGetProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productId := ctx.Param("id")
		productDetails, err := prodHandler.ProductFrontStoreService.GetProductById(productId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"product": productDetails})
	}
}

func (prodHandler ProductFrontStoreHandler) HandleProductsByCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryId := ctx.Param("id")
		products, err := prodHandler.ProductFrontStoreService.GetProductsByCategoryId(categoryId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"products": products})
	}
}
