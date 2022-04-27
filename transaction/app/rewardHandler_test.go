package app_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/app"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/domain"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/transaction/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/utils/errs"
)

func TestHandleGetTransactionRecordByUserID(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockTransactionService)
		expected   int
	}{
		{
			name: "SuccessGetTransactionRecordByID",
			createStub: func(mps *mocks.MockTransactionService) {
				mps.EXPECT().
					GetTransactionByUserId("xyz").
					Return(&domain.Transaction{}, nil)
			},
			expected: 202,
		},
		{
			name: "FailGetTransactionRecordByID",
			createStub: func(mps *mocks.MockTransactionService) {
				errstring := "record not found"
				mps.EXPECT().
					GetTransactionByUserId("xyz").
					Return(nil, &errs.AppError{Message: errstring})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockTransactionService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.TransactionHandler{
				TransactionService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/transaction/xyz", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}

func TestHandleUpdateTransactionRecordByUserID(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockTransactionService)
		expected   int
	}{
		{
			name: "SuccessUpdateTransactionRecordByID",
			createStub: func(mps *mocks.MockTransactionService) {
				mps.EXPECT().
					UpdateTransactionByUserId("xyz", 10).
					Return(true, nil)
			},
			expected: 202,
		},
		{
			name: "FailUpdateTransactionRecordByID",
			createStub: func(mps *mocks.MockTransactionService) {
				errstring := "record not found"
				mps.EXPECT().
					UpdateTransactionByUserId("xyz", 10).
					Return(false, &errs.AppError{Message: errstring})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockTransactionService(mockCtrl)
			testcase.createStub(mockService)

			type requestDTO struct {
				UserID            string `json:"user_id"`
				TransactionPoints int    `json:"transaction_points"`
			}

			requestData, _ := json.Marshal(
				requestDTO{UserID: "xyz", TransactionPoints: 10},
			)

			router := app.SetupRouter(app.TransactionHandler{
				TransactionService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/transaction/xyz", bytes.NewReader(requestData))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockTransactionService(mockCtrl)

	router := app.SetupRouter(app.TransactionHandler{
		TransactionService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/transaction/xyz", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}
