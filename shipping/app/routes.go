package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())
}

func ShippingRouter(router *gin.Engine, shippingHandler ShippingHandler) {

	router.POST("/shippingaddress", shippingHandler.handleShippingAddress())
	router.GET("/shippingaddress/:id", shippingHandler.HandleGetShippingAddrssByID())
	router.PUT("/shippingaddress/:id", shippingHandler.HandleUpdateShippingAddressByID())
	router.DELETE("/shippingaddress/:id", shippingHandler.HandleDeleteShippingAddressById())

	router.POST("/shippingcost", shippingHandler.handleShippingCost())
	router.GET("/shippingcost/:city", shippingHandler.HandleGetShippingCostByCity())
	router.PUT("/shippingcost", shippingHandler.HandleUpdateShippingCostByCity())
	router.DELETE("/shippingcost/:city", shippingHandler.HandleDeleteShippingCostByCity())

}
