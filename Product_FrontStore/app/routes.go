package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())
}

func ProductFrontStoreRouter(gin *gin.Engine) {
	productApiGroup := gin.Group("/products")
	productApiGroup.GET("/", productFrontStoreHandler.HandleGetAllProducts())
	productApiGroup.GET("/:id", productFrontStoreHandler.HandleGetProductByID())
	productApiGroup.GET("/category/:id", productFrontStoreHandler.HandleProductsByCategory())
}
