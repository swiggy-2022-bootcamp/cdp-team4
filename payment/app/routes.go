package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())
}

func PayRouter(gin *gin.Engine) {
	gin.POST("/pay", paymentHandler.handlePay())
	gin.GET("/pay/{id}", paymentHandler.handleGetPayRecordByID())
	gin.GET("/pay/{user_id}", paymentHandler.handleGetPayRecordsByUserID())
	gin.PUT("/updatePayStatus", paymentHandler.handleUpdatePayStatus())
}
