package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/infra"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var categoryHandler CategoryHandler

func setupRouter() *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	CategoryRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Category API"
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	categoryHandler = CategoryHandler{CategoryService: domain.NewCategoryService(dynamoRepository)}

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
