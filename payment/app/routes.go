package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())
}

func PayRouter(gin *gin.Engine) {
	p := gin.Group("/pay")
	{
		p.POST("/", paymentHandler.handlePay())
		p.PUT("/", paymentHandler.handleUpdatePayStatus())

		p.GET("/user/{user_id}", paymentHandler.handleGetPayRecordsByUserID())
		p.GET("/{id}", paymentHandler.handleGetPayRecordByID())

		p.POST("/paymentMethods", paymentHandler.handleAddPaymentMethods())
		p.GET("/paymentMethods", paymentHandler.handleGetPaymentMethods())
	}
}
