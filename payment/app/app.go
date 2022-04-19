package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/infra/logger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var log logrus.Logger = *logger.GetLogger()

func SetupRouter(paymentHandler PayHandler) *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	// payment router
	PayRouter(router, paymentHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func ConfigureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Payment API"
}

func Start(testMode bool) {
	dynamoRepository := infra.NewDynamoRepository()
	paymentService := domain.NewPaymentService(dynamoRepository)
	paymentHandler := NewPaymentHandler(paymentService)

	// PayHandler{PaymentService: domain.NewPaymentService(dynamoRepository)}

	ConfigureSwaggerDoc()
	router := SetupRouter(paymentHandler)

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("PORT")
	if !testMode {
		log.WithFields(logrus.Fields{"PORT": PORT}).Info("Running on PORT")
		router.Run(":" + PORT)
	}
}
