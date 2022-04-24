package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/domain"
)

type RewardHandler struct {
	RewardService domain.RewardService
}

type RewardRecordDTO struct {
	UserID       string `json:"user_id"`
	RewardPoints int    `json:"reward_points"`
}

// Get Reward by userID
// @Summary      Get Reward by userId
// @Description  This Handle returns Reward given userid
// @Tags         Reward
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /reward/:userId    [get]
func (rh RewardHandler) HandleGetRewardRecordByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("userId")
		if id == "" {
			log.WithFields(logrus.Fields{"message": "Invalid User ID", "status": http.StatusBadRequest}).
				Error("Error while Getting reward by reward id")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid User ID"})
			return
		}
		res, err := rh.RewardService.GetRewardByUserId(id)

		if err != nil {
			log.WithFields(logrus.Fields{"message": err.Message, "status": http.StatusBadRequest}).
				Error("Record not found")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found: "+err.Message})
			return
		}
		rewarddto := convertRewardModeltoRewardDTO(*res)
		ctx.JSON(http.StatusAccepted, gin.H{"record": rewarddto})
	}
}

// Update reward for a userId
// @Summary      Update reward points for a userId
// @Description  This Handle Update reward given user id
// @Tags         Reward
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /reward/:userId   [put]
func (oh RewardHandler) HandleUpdateRewardByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestDTO struct {
			UserID       string `json:"user_id"`
			RewardPoints int    `json:"reward_points"`
		}
		if err := ctx.BindJSON(&requestDTO); err != nil {
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ok, err := oh.RewardService.UpdateRewardByUserId(requestDTO.UserID, requestDTO.RewardPoints)
		if !ok {
			log.WithFields(logrus.Fields{"message": err.Message, "status": http.StatusBadRequest}).
				Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Message})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "reward record updated"})
	}
}

func convertRewardModeltoRewardDTO(reward domain.Reward) RewardRecordDTO {
	return RewardRecordDTO{
		UserID:       reward.UserID,
		RewardPoints: reward.RewardPoints,
	}
}
func NewRewardHandler(rewardService domain.RewardService) RewardHandler {
	return RewardHandler{
		RewardService: rewardService,
	}
}
