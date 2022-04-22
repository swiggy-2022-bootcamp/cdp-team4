package domain

import (
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/utils/errs"
)

type Reward struct {
	Id           string `json:"id"`
	UserID       string `json:"user_id"`
	RewardPoints int    `json:"reward_points"`
}

type RewardRepository interface {
	InsertReward(Reward) (string, *errs.AppError)
	FindAllRewards() ([]Reward, *errs.AppError)
	FindRewardById(string) (*Reward, *errs.AppError)
	DeleteRewardById(string) (bool, *errs.AppError)
}

func NewReward(userId string, rewardpoints int) *Reward {
	return &Reward{
		UserID:           userId,
		RewardPoints: rewardpoints,
	}
}
