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

func RegisterProductAdminRoutes() {
	productAdminHandler := productAdminHandler{}
	products := v1.Group("/products")

	v1.GET("/products", ValidateAuthToken(), productAdminHandler.GetAllProducts)

	products.POST("/", ValidateAuthToken(), productAdminHandler.AddProduct)
	products.GET("/:id", ValidateAuthToken(), productAdminHandler.GetProductByID)
	products.PUT("/:id", ValidateAuthToken(), productAdminHandler.UpdateProductByID)
	products.DELETE("/:id", ValidateAuthToken(), productAdminHandler.DeleteProductByID)

	search := products.Group("/search")
	search.GET("/category/:categoryid", ValidateAuthToken(), productAdminHandler.SearchByCategoryID)
	search.GET("/manufacturer/:id", ValidateAuthToken(), productAdminHandler.SearchByManufacturerID)
	search.GET("/keyword/:keyword", ValidateAuthToken(), productAdminHandler.SearchByKeyword)
}

func RegisterPaymentRoutes() {
	paymentHandler := PaymentHandler{}
	payment := v1.Group("/pay")

	payment.POST("/", ValidateAuthToken(), paymentHandler.InitiatePayment)
	payment.POST("/paymentMethod", ValidateAuthToken(), paymentHandler.AddPaymentMethod)
	payment.GET("/paymentMethod/:id", ValidateAuthToken(), paymentHandler.GetPaymentMethod)
}

func RegisterProductFrontStoreRoutes() {
	// productFrontStoreHandler := productFrontStoreHandler{}
	// products := v1.Group("/products")

	// v1.GET("/products", ValidateAuthToken(), productFrontStoreHandler.HandleGetAllProducts)

	// products.GET("/:id", ValidateAuthToken(), productFrontStoreHandler.HandleGetProductByID)
	// products.GET("/category/:id", ValidateAuthToken(), productFrontStoreHandler.HandleGetProductsByCategory)
}

func RegisterCategoryRoutes() {
	// categoryHandler := categoryHandler{}
	// categories := v1.Group("/categories")

	// categories.POST("/", ValidateAuthToken(), categoryHandler.HandleAddCategory)
	// categories.GET("/", ValidateAuthToken(), categoryHandler.HandleGetAllCategories)
	// categories.GET("/:id", ValidateAuthToken(), categoryHandler.HandleGetCategoryByID)
	// categories.PUT("/:id", ValidateAuthToken(), categoryHandler.HandleUpdateCategoryByID)
	// categories.DELETE("/", ValidateAuthToken(), categoryHandler.HandleDeleteCategories)
	// categories.DELETE("/:id", ValidateAuthToken(), categoryHandler.HandleDeleteCategoryByID)
}

func RegisterRewardRouter() {
	rewardHandler := RewardHandler{}
	rewardApiGroup := v1.Group("/reward")
	rewardApiGroup.GET("/:userId", ValidateAuthToken(), rewardHandler.GetRewardByUserID)
	rewardApiGroup.PUT("/:userId", ValidateAuthToken(), rewardHandler.UpdateReward)
}

func RegisterCartRouter() {
	cartHandler := CartHandler{}
	cartApiGroup := v1.Group("/cart")
	cartApiGroup.GET("/:userId", ValidateAuthToken(), cartHandler.GetCartByUserID)
	cartApiGroup.POST("/", ValidateAuthToken(), cartHandler.CreateCart)
	cartApiGroup.PUT("/:userId", ValidateAuthToken(), cartHandler.UpdateCartByUserID)
	cartApiGroup.DELETE("/empty/:userId", ValidateAuthToken(), cartHandler.DeleteCartByUserID)
	cartApiGroup.DELETE("/:userId", ValidateAuthToken(), cartHandler.DeleteCartItemByUserId)
	// cartApiGroup.GET("/", cartHandler.HandleGetAllRecords())
}

func RegisterTransactionRouter() {
	transactionHandler := TransactionHandler{}
	transactionApiGroup := v1.Group("/transaction")
	transactionApiGroup.GET("/:userId", ValidateAuthToken(), transactionHandler.GetTransactionByUserID)
	transactionApiGroup.PUT("/:userId", ValidateAuthToken(), transactionHandler.UpdateTransaction)
}
