package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/infra/logger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var log logrus.Logger = *logger.GetLogger()

//var productFrontStoreHandler ProductFrontStoreHandler

func SetupRouter(productFrontStoreHandler ProductFrontStoreHandler) *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	ProductFrontStoreRouter(router, productFrontStoreHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func ConfigureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Product Front Store API"
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	productFrontStoreService := domain.NewProductFrontStoreService(dynamoRepository)
	productFrontStoreHandler := NewProductFrontStoreHandler(productFrontStoreService)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	ConfigureSwaggerDoc()

	PORT := os.Getenv("PORT")
	router := SetupRouter(productFrontStoreHandler)
	router.Run(":" + PORT)
	log.WithFields(logrus.Fields{"PORT": PORT}).Info("Running on PORT")
}
