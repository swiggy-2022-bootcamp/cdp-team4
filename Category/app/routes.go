package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())
}

func CategoryRouter(gin *gin.Engine, categoryHandler CategoryHandler) {
	categoryApiGroup := gin.Group("/categories")
	categoryApiGroup.POST("/", categoryHandler.HandleAddCategory())
	categoryApiGroup.GET("/", categoryHandler.HandleGetAllCategories())
	categoryApiGroup.GET("/:id", categoryHandler.HandleGetCategoryByID())
	categoryApiGroup.PUT("/:id", categoryHandler.HandleUpdateCategoryByID())
	categoryApiGroup.DELETE("/", categoryHandler.HandleDeleteCategories())
	categoryApiGroup.DELETE("/:id", categoryHandler.HandleDeleteCategoryByID())
}
