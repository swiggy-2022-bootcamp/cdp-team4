package app

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(router *gin.Engine) {
	router.GET("/", HealthCheck())
}

func rewardRouter(router *gin.Engine,rewardHandler RewardHandler) {
	rewardApiGroup := router.Group("/reward")
	rewardApiGroup.GET("/:userId", rewardHandler.HandleGetRewardRecordByUserID())
	rewardApiGroup.PUT("/:userId", rewardHandler.HandleUpdateRewardByUserId())
}
