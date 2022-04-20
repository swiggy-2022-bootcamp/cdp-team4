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
		p.PUT("/", paymentHandler.HandleUpdatePayStatus())

		p.GET("/user/:user_id", paymentHandler.handleGetPayRecordsByUserID())
		p.GET("/:id", paymentHandler.HandleGetPayRecordByID())

		p.POST("/paymentMethods", paymentHandler.HandleAddPaymentMethods())
		p.GET("/paymentMethods/:id", paymentHandler.HandleGetPaymentMethods())
	}
}
