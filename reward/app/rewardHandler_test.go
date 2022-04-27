package app_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/app"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/domain"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/reward/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/utils/errs"
)

func TestHandleGetRewardRecordByUserID(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockRewardService)
		expected   int
	}{
		{
			name: "SuccessGetRewardRecordByID",
			createStub: func(mps *mocks.MockRewardService) {
				mps.EXPECT().
					GetRewardByUserId("xyz").
					Return(&domain.Reward{}, nil)
			},
			expected: 202,
		},
		{
			name: "FailGetRewardRecordByID",
			createStub: func(mps *mocks.MockRewardService) {
				errstring := "record not found"
				mps.EXPECT().
					GetRewardByUserId("xyz").
					Return(nil, &errs.AppError{Message: errstring})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockRewardService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.RewardHandler{
				RewardService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/reward/xyz", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}


func TestHandleUpdateRewardRecordByUserID(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockRewardService)
		expected   int
	}{
		{
			name: "SuccessUpdateRewardRecordByID",
			createStub: func(mps *mocks.MockRewardService) {
				mps.EXPECT().
					UpdateRewardByUserId("xyz", 10).
					Return(true, nil)
			},
			expected: 202,
		},
		{
			name: "FailUpdateRewardRecordByID",
			createStub: func(mps *mocks.MockRewardService) {
				errstring := "record not found"
				mps.EXPECT().
					UpdateRewardByUserId("xyz", 10).
					Return(false, &errs.AppError{Message: errstring})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockRewardService(mockCtrl)
			testcase.createStub(mockService)

			type requestDTO struct {
				UserID       string `json:"user_id"`
				RewardPoints int    `json:"reward_points"`
			}

			requestData, _ := json.Marshal(
				requestDTO{UserID: "xyz", RewardPoints: 10},
			)

			router := app.SetupRouter(app.RewardHandler{
				RewardService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/reward/xyz", bytes.NewReader(requestData))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockRewardService(mockCtrl)

	router := app.SetupRouter(app.RewardHandler{
		RewardService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/reward/xyz", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}
