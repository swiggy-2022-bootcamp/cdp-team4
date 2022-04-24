package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/", HealthCheck())

}

func rewardRouter(gin *gin.Engine,rewardHandler RewardHandler) {
	productApiGroup := gin.Group("/reward")
	productApiGroup.GET("/:userId", rewardHandler.HandleGetRewardRecordByUserID())
	productApiGroup.PUT("/:userId", rewardHandler.HandleUpdateRewardByUserId())
}
