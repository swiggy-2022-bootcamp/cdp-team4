package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/infra"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var productFrontStoreHandler ProductFrontStoreHandler

func setupRouter() *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	ProductFrontStoreRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Product Front Store API"
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	productFrontStoreHandler = ProductFrontStoreHandler{ProductFrontStoreService: domain.NewProductFrontStoreService(dynamoRepository)}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	configureSwaggerDoc()

	PORT := os.Getenv("PORT")
	router := setupRouter()
	router.Run(":" + PORT)
}
