package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/docs"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Shipping API"
}

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("SHIPPING_SERVICE_PORT")

	configureSwaggerDoc()
	router := setupRouter()

	router.Run(":" + PORT)
}
