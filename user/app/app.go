package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/swiggy-2022-bootcamp/cdp-team4/user/infra"	
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"		
)

type Routes struct {
	router *gin.Engine
}

var userHandler UserHandler

func setupRouter() *gin.Engine {
	router := gin.Default()
	
	// health check route
	HealthCheckRouter(router)

	// user route
	UserRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger User API"
}

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("PORT")

	userDynamodbRepository := infra.NewDynamoRepository()

	userHandler = UserHandler{
		userService: domain.NewUserService(userDynamodbRepository),
	}

	configureSwaggerDoc()
	router := setupRouter()

	router.Run(":" + PORT)
}