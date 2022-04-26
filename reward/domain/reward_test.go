package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToReturnNewReward(t *testing.T) {

	userId := "12345678"
	rewardPoints:=10
	newCart := NewReward(userId, rewardPoints)

	assert.Equal(t, userId, newCart.UserID)
	assert.Equal(t, rewardPoints, newCart.RewardPoints)
}
