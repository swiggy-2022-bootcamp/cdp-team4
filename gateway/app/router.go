package app

import (
	"github.com/gin-gonic/gin"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swiggy-2022-bootcamp/cdp-team4/gateway/docs"
)

var v1 *gin.RouterGroup
var Router *gin.Engine

func init() {
	Router = gin.Default()
	api := Router.Group("/api")
	v1 = api.Group("/v1")

	//docs.SwaggerInfo.BasePath = "/api/v1"
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RegisterUserRoutes() {
	userHandler := userHandler{}
	users := v1.Group("/user")

	v1.GET("/users", ValidateAuthToken(), userHandler.GetAllUsers)

	users.POST("/", userHandler.CreateUser)
	users.GET("/", ValidateAuthToken(), userHandler.GetUser)
	users.PATCH("/", ValidateAuthToken(), userHandler.UpdateUser)
	users.DELETE("/", ValidateAuthToken(), userHandler.DeleteUser)
}
