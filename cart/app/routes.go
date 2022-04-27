package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(router *gin.Engine) {
	router.GET("/", HealthCheck())
}

func CartRouter(router *gin.Engine,cartHandler CartHandler) {
	cartApiGroup := router.Group("/cart")
	cartApiGroup.POST("/", cartHandler.HandleCreateCart())
	cartApiGroup.GET("/", cartHandler.HandleGetAllRecords())
	cartApiGroup.GET("/:userId", cartHandler.HandleGetCartRecordByUserID())
	cartApiGroup.PUT("/:userId", cartHandler.HandleUpdateCartItemByUserId())
	cartApiGroup.DELETE("/empty/:userId", cartHandler.HandleDeleteCartByUserId())
	cartApiGroup.DELETE("/:userId", cartHandler.HandleDeleteCartItemByUserId())
}
