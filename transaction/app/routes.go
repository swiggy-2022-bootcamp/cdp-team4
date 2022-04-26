package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(router *gin.Engine) {
	router.GET("/", HealthCheck())
}

func transactionRouter(router *gin.Engine, transactionHandler TransactionHandler) {
	transactionApiGroup := router.Group("/transaction")
	transactionApiGroup.GET("/:userId", transactionHandler.HandleGetTransactionRecordByUserID())
	transactionApiGroup.PUT("/:userId", transactionHandler.HandleUpdateTransactionByUserId())
}
