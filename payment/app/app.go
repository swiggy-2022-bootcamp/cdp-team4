package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/infra"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var paymentHandler PayHandler

func setupRouter() *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	// payment router
	PayRouter(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Payment API"
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	paymentHandler = PayHandler{PaymentService: domain.NewPaymentService(dynamoRepository)}

	configureSwaggerDoc()
	router := setupRouter()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("PORT")
	router.Run(":" + PORT)
}
