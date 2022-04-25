package app

import (
	"github.com/gin-gonic/gin"
)

var v1 *gin.RouterGroup
var Router *gin.Engine

func init() {
	Router = gin.Default()
	api := Router.Group("/api")
	v1 = api.Group("/v1")

	//docs.SwaggerInfo.BasePath = "/api/v1"
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RegisterUserRoutes() {
	userHandler := userHandler{}
	users := v1.Group("/user")

	v1.GET("/users", ValidateAuthToken(), userHandler.GetAllUsers)

	users.POST("/", userHandler.CreateUser)
	users.GET("/", ValidateAuthToken(), userHandler.GetUser)
	users.PATCH("/", ValidateAuthToken(), userHandler.UpdateUser)
	users.DELETE("/", ValidateAuthToken(), userHandler.DeleteUser)
}

func RegisterOrderRoutes() {
	orderhandler := orderHandler{}
	orders := v1.Group("/order")

	v1.GET("/orders", ValidateAuthToken(), orderhandler.GetAllOrders)

	orders.POST("/", ValidateAuthToken(), orderhandler.CreateOrder)
	orders.GET("/:id", ValidateAuthToken(), orderhandler.GetOrderByID)
	orders.GET("/user/:user_id", ValidateAuthToken(), orderhandler.GetOrderByUserID)
	orders.GET("/status/:status", ValidateAuthToken(), orderhandler.GetOrderByStatus)
	orders.PUT("/:id", ValidateAuthToken(), orderhandler.UpdateOrder)
	orders.DELETE("/:id", ValidateAuthToken(), orderhandler.DeleteOrder)
	orders.POST("/confirm/:user_id", ValidateAuthToken(), orderhandler.ConfirmOrder)
	orders.GET("/order/invoice/:order_id", ValidateAuthToken(), orderhandler.GetOrderInvoice)
}

func RegisterShippingRoutes() {
	shippingHandler := shippingHandler{}
	shippingAddress := v1.Group("/shippingaddress")

	shippingAddress.POST("/", ValidateAuthToken(), shippingHandler.CreateShippingAddress)
	shippingAddress.GET("/:id", ValidateAuthToken(), shippingHandler.GetShippingAddressByID)
	shippingAddress.PUT("/:id", ValidateAuthToken(), shippingHandler.UpdateShippingAddress)
	shippingAddress.DELETE("/:id", ValidateAuthToken(), shippingHandler.DeleteShippingAddress)

	shippingCost := v1.Group("/shippingcost")

	shippingCost.POST("/", ValidateAuthToken(), shippingHandler.CreateShippingCost)
	shippingCost.GET("/:city", ValidateAuthToken(), shippingHandler.GetShippingCostByCity)
	shippingCost.PUT("/", ValidateAuthToken(), shippingHandler.UpdateShippingCost)
	shippingCost.DELETE("/:city", ValidateAuthToken(), shippingHandler.DeleteShippingCost)

}
