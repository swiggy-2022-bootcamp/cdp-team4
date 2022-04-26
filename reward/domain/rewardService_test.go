package domain_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/mocks"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/utils/errs"
)

// func TestGetRazorpayRewardLink(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	mockDynamoRepo := mocks.NewMockRewardRepository(mockCtrl)

// 	service := domain.NewRewardService(mockDynamoRepo)
// }

func TestShouldReturnNewService(t *testing.T) {
	newService := domain.NewRewardService(nil)
	assert.NotNil(t, newService)
}

func TestCreateRewardRecord(t *testing.T) {
	reward := domain.Reward{
		//Id:primitive.NewObjectID().Hex(),
		UserID:       "randomUserId",
		RewardPoints: 10,
	}
	testcases := []struct {
		name       string
		createStub func(mocks.MockRewardRepository)
		assertTest func(*testing.T, string, *errs.AppError)
	}{
		{
			name: "FailCreateRewardRecord",
			createStub: func(mrr mocks.MockRewardRepository) {
				errstring := "unable to insert record"
				mrr.EXPECT().
					InsertReward(reward).
					Return("", &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m string, err *errs.AppError) {
				assert.Equal(t, "", m)
				assert.NotNil(t, err)
			},
		},
		{
			name: "SuccessCreateRewardRecord",
			createStub: func(mrr mocks.MockRewardRepository) {
				mrr.EXPECT().InsertReward(reward).Return(reward.Id, nil)
			},
			assertTest: func(t *testing.T, m string, err *errs.AppError) {
				assert.NotNil(t, m)
				assert.Nil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockRewardRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewRewardService(mockDynamoRepo)

			data, err := service.CreateReward(
				reward.UserID,
				reward.RewardPoints,
			)

			testcase.assertTest(t, data, err)
		})
	}

}

func TestGetRewardRecordByuserId(t *testing.T) {
	reward := domain.Reward{
		//Id:primitive.NewObjectID().Hex(),
		UserID:       "randomUserId",
		RewardPoints: 10,
	}

	testcases := []struct {
		name       string
		createStub func(mocks.MockRewardRepository)
		assertTest func(*testing.T, *domain.Reward, *errs.AppError)
	}{
		{
			name: "SuccessGetRewardRecordByUserId",
			createStub: func(mpdr mocks.MockRewardRepository) {
				mpdr.EXPECT().FindRewardByUserId("randomUserId").Return(&reward, nil)
			},
			assertTest: func(t *testing.T, m *domain.Reward, err *errs.AppError) {
				assert.NotNil(t, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailGetRewardRecordByUserId",
			createStub: func(mpdr mocks.MockRewardRepository) {
				errstring := "unable to find record"
				mpdr.EXPECT().
					FindRewardByUserId("randomUserId").
					Return(nil, &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m *domain.Reward, err *errs.AppError) {
				assert.Nil(t, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockRewardRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewRewardService(mockDynamoRepo)

			data, err := service.GetRewardByUserId("randomUserId")

			testcase.assertTest(t, data, err)
		})
	}

}

func TestUpdateRewardPoints(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(mocks.MockRewardRepository)
		assertTest func(*testing.T, bool, *errs.AppError)
	}{
		{
			name: "SuccessUpdateReward",
			createStub: func(mpdr mocks.MockRewardRepository) {
				mpdr.EXPECT().UpdateRewardByUserId("randomUserId",20).Return(true, nil)
			},
			assertTest: func(t *testing.T, m bool, err *errs.AppError) {
				assert.Equal(t,true, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailUpdateReward",
			createStub: func(mpdr mocks.MockRewardRepository) {
				errstring := "unable to find record"
				mpdr.EXPECT().
					UpdateRewardByUserId("randomUserId",20).
					Return(false, &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m bool, err *errs.AppError) {
				assert.Equal(t,false, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockRewardRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewRewardService(mockDynamoRepo)

			data, err := service.UpdateRewardByUserId("randomUserId", 20)

			testcase.assertTest(t, data, err)
		})
	}

}
