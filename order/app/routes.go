package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())
}

func OrderRouter(router *gin.Engine, orderHandler OrderHandler) {

	router.POST("/order", orderHandler.handleOrder())
	router.GET("/orders", orderHandler.HandleGetAllRecords())
	router.GET("/order/:id", orderHandler.HandleGetOrderRecordByID())
	router.GET("/order/user/:user_id", orderHandler.HandleGetOrderRecordsByUserID())
	router.GET("/order/status/:status", orderHandler.HandleGetOrderRecordsByStatus())
	router.PUT("/order/status", orderHandler.handleUpdateOrderStatus())
	router.DELETE("/order/:id", orderHandler.HandleDeleteOrderById())
	router.POST("/confirm/:user_id", orderHandler.HandleAddOrderFromCheckout())
	router.GET("/order/invoice/:order_id", orderHandler.HandleGetOrderInvoice())
}
