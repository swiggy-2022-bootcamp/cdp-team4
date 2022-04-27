package infra

import (
	"time"
)

type RewardModel struct {
	Id           string    `json:"id"`
	UserID       string    `json:"user_id"`
	RewardPoints int       `json:"reward_points"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
