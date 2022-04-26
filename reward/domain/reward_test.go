package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToReturnNewReward(t *testing.T) {

	userId := "12345678"
	rewardPoints:=10
	newReward := NewReward(userId, rewardPoints)

	assert.Equal(t, userId, newReward.UserID)
	assert.Equal(t, rewardPoints, newReward.RewardPoints)
}
