package app

import (
	"os"
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/infra/logger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/swiggy-2022-bootcamp/cdp-team4/user/infra"	
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"		
)

var log logrus.Logger = *logger.GetLogger()

type Routes struct {
	router *gin.Engine
}

func SetupRouter(userHandler UserHandler) *gin.Engine {
	router := gin.Default()
	
	// health check route
	HealthCheckRouter(router)

	// user route
	UserRouter(router, userHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger User API"
}

func Start() {
	err1 := godotenv.Load(".env")
	if err1 != nil {
		logrus.Fatal(err1)
		return
	}
	PORT := os.Getenv("PORT")


	userDynamodbRepository := infra.NewDynamoRepository()

	userHandler := UserHandler{
		UserService: domain.NewUserService(userDynamodbRepository),
	}

	configureSwaggerDoc()
	router := SetupRouter(userHandler)

	router.Run(":" + PORT)
}