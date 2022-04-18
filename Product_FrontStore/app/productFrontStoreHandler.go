package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/domain"
)

type ProductFrontStoreHandler struct {
	ProductFrontStoreService domain.ProductFrontStoreService
}

// Get all products
// @Summary      Get all products
// @Description  This Handle allows front store to fetch all the products in the datastore
// @Tags         Product Front Store
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/    [get]
func (prodHandler ProductFrontStoreHandler) HandleGetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := prodHandler.ProductFrontStoreService.GetProducts()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to fetch categories")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"products": products})
		log.WithFields(logrus.Fields{"products": products, "status": http.StatusOK}).
			Info("fetch product records")
	}
}
// Get product details by category
// @Summary      Get product by id 
// @Description  This Handle allows front store to fetch the product by id in the datastore
// @Tags         Product Front Store
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/:id    [get]
func (prodHandler ProductFrontStoreHandler) HandleGetProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productId := ctx.Param("id")
		productDetails, err := prodHandler.ProductFrontStoreService.GetProductById(productId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to fetch product by id")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"product": productDetails})
		log.WithFields(logrus.Fields{"product": productDetails, "status": http.StatusOK}).
			Info("fetch product record")
	}
}
// Get all products
// @Summary      Get products by category id
// @Description  This Handle allows front store to fetch all the products in the datastore based category id 
// @Tags         Product Front Store
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/    [get]
func (prodHandler ProductFrontStoreHandler) HandleProductsByCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryId := ctx.Param("id")
		products, err := prodHandler.ProductFrontStoreService.GetProductsByCategoryId(categoryId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to fetch products by category id")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"products": products})
		log.WithFields(logrus.Fields{"product": products, "status": http.StatusOK}).
			Info("fetch product by category id")
	}
}
