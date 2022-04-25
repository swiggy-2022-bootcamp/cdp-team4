package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/infra/logger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var log logrus.Logger = *logger.GetLogger()

func SetupRouter(categoryHandler CategoryHandler) *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	CategoryRouter(router, categoryHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func ConfigureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Category API"
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	categoryService := domain.NewCategoryService(dynamoRepository)
	categoryHandler := NewCategoryHandler(categoryService)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	ConfigureSwaggerDoc()

	PORT := os.Getenv("PORT")
	router := SetupRouter(categoryHandler)
	router.Run(":" + PORT)
	log.WithFields(logrus.Fields{"PORT": PORT}).Info("Running on PORT")
}
