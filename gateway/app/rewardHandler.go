package app

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	// "log"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)


type RewardRecordDTO struct {
	UserID       string `json:"user_id"`
	RewardPoints int    `json:"reward_points"`
}

type RewardHandler struct {
}

func (rh RewardHandler) GetRewardByUserID(c *gin.Context) {

	rewardServiceURI := os.Getenv("REWARD_SERVICE_URI") + "/reward/" + c.Param("userId")

	req, err := http.NewRequest("GET", rewardServiceURI, nil)
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var rewardResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&rewardResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(200, res.Body)
	//c.JSON(rewardResDTO.Status, rewardResDTO)
}

func (rh RewardHandler) UpdateReward(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	rewardServiceURI := os.Getenv("REWARD_SERVICE_URI") + "/reward/" + c.Param("userId")

	var updateReward RewardRecordDTO

	if err := c.BindJSON(&updateReward); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	updateReward.UserID=c.Param("userId")

	fmt.Printf("here1: %v", updateReward)

	rewardJSON, err1 := json.Marshal(updateReward)

	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("PUT", rewardServiceURI, bytes.NewBuffer(rewardJSON))

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var rewardResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&rewardResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	c.JSON(200, res.Body)
	//c.JSON(rewardResDTO.Status, rewardResDTO)
}
