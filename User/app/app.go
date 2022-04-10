package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/User/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)

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

	configureSwaggerDoc()
	router := setupRouter()

	router.Run(":" + PORT)
}