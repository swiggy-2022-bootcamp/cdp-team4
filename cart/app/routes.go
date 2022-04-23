package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(router *gin.Engine) {
	router.GET("/", HealthCheck())
}

