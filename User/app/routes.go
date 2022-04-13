package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())
}

func UserRouter(gin *gin.Engine) {
	gin.POST("/user", userHandler.HandleUserCreation())
}