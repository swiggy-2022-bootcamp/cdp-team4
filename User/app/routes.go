package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())
}

func UserRouter(gin *gin.Engine) {
	gin.POST("/user", userHandler.HandleUserCreation())
	gin.GET("/users", userHandler.HandleGetAllUsers())
	gin.GET("/user/:id", userHandler.HandleGetUserByID())
	gin.DELETE("/user/:id", userHandler.HandleDeleteUserByID())
}