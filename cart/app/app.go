package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/infra"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var cartHandler CartHandler

func setupRouter(cartHandler CartHandler) *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	CartRouter(router, cartHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Payment API"
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	cartHandler = CartHandler{CartService: domain.NewCartService(dynamoRepository)}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("PORT")

	configureSwaggerDoc()
	router := setupRouter(cartHandler)

	
	if err := router.Run(":" + PORT); err != nil {
		log.Fatal(err)
	}
}
