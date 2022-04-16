package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/infra"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var productAdminHandler ProductAdminHandler

func setupRouter() *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	ProductAdminRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Product Admin API"
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	productAdminHandler = ProductAdminHandler{ProductAdminService: domain.NewProductAdminService(dynamoRepository)}
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
