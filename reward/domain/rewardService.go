package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/reward/utils/errs"

type RewardService interface {
	CreateReward(string,int) (string, *errs.AppError)
	//GetRewardById(string) (*Reward, *errs.AppError)
	GetRewardByUserId(string) (*Reward, *errs.AppError)
	UpdateRewardByUserId(string,int) (bool, *errs.AppError)
}

type service struct {
	rewardRepository RewardRepository
}

func (s service) CreateReward(userId string, rewardPoints int) (string, *errs.AppError) {
	reward := NewReward(userId, rewardPoints)
	resultId, err := s.rewardRepository.InsertReward(*reward)
	if err != nil {
		return "", err
	}
	return resultId, nil
}

// func (s service) GetRewardById(rewardId string) (*Reward, *errs.AppError) {
// 	res, err := s.rewardRepository.FindRewardById(rewardId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (s service) GetRewardByUserId(userId string) (*Reward, *errs.AppError) {
	res, err := s.rewardRepository.FindRewardByUserId(userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) UpdateRewardByUserId(userId string, points int) (bool, *errs.AppError) {
	_, err := s.rewardRepository.UpdateRewardByUserId(userId,points)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewRewardService(rewardRepository RewardRepository) RewardService {
	return &service{
		rewardRepository: rewardRepository,
	}
}
