package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/infra/logger"
)

var log logrus.Logger = *logger.GetLogger()

func setupRouter(rewardHandler RewardHandler) *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	rewardRouter(router,rewardHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Reward API"
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	rewardHandler := RewardHandler{RewardService: domain.NewRewardService(dynamoRepository)}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("PORT")

	configureSwaggerDoc()
	router := setupRouter(rewardHandler)

	router.Run(":" + PORT)
	// if err := r.Run(":3000"); err != nil {
	// 	log.Fatal(err)
	// }
}
