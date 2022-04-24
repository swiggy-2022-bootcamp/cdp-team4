package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())
}

// function that is used to register all the routes and their
// respective handler methods
func ProductAdminRouter(gin *gin.Engine) {
	productApiGroup := gin.Group("/products")
	productApiGroup.POST("/", productAdminHandler.HandleAddProduct())
	productApiGroup.GET("/", productAdminHandler.HandleGetAllProducts())
	productApiGroup.GET("/:id", productAdminHandler.HandleGetProductByID())
	productApiGroup.PUT("/:id", productAdminHandler.HandleUpdateProduct())
	productApiGroup.DELETE("/:id", productAdminHandler.HandleDeleteProductByID())

	searchRoutes := productApiGroup.Group("/search")
	searchRoutes.GET("/category/:categoryid", productAdminHandler.HandleSearchByCategoryID())
	searchRoutes.GET("/manufacturer/:id", productAdminHandler.HandleSearchByManufacturerID())
	searchRoutes.GET("/keyword/:keyword", productAdminHandler.HandleSearchByKeyword())

}
