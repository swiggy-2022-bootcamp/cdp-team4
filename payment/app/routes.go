package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())
}

func PayRouter(gin *gin.Engine, paymentHandler PayHandler) {
	p := gin.Group("/pay")
	{
		p.POST("/", paymentHandler.HandlePay())
		p.PUT("/", paymentHandler.handleUpdatePayStatus())

		p.GET("/user/:user_id", paymentHandler.handleGetPayRecordsByUserID())
		p.GET("/:id", paymentHandler.HandleGetPayRecordByID())

		p.POST("/paymentMethods", paymentHandler.handleAddPaymentMethods())
		p.GET("/paymentMethods/:id", paymentHandler.handleGetPaymentMethods())
	}
}
