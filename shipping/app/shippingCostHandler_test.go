package app_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/app"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"
)

func TestHandleShippingCost(t *testing.T) {
	// router := gin.Default()

	requestData := app.ShippingCostRecordDTO{
		City: "Banglore",
		Cost: 100,
	}

	testCases := []struct {
		name       string
		createStub func(*mocks.MockShippingCostService)
		expected   int
	}{
		{
			name: "SucessHandleShippingCost",
			createStub: func(mps *mocks.MockShippingCostService) {

				mps.EXPECT().
					CreateShippingCost(gomock.Any(), gomock.Any()).
					Return(true, nil)
			},
			expected: 202,
		},
		{
			name: "ErrorCreateShippingCostRecord",
			createStub: func(mps *mocks.MockShippingCostService) {
				mps.EXPECT().
					CreateShippingCost(gomock.Any(), gomock.Any()).
					Return(false, &errs.AppError{Message: "Unable to insert record"})
			},
			expected: 400,
		},
		{
			name: "ErrorNothingReturnFromCreateShippingCostRecord",
			createStub: func(mps *mocks.MockShippingCostService) {
				mps.EXPECT().
					CreateShippingCost(gomock.Any(), gomock.Any()).
					Return(false, &errs.AppError{Message: "Unable to insert record"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockShippingCostService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.ShippingHandler{
				ShippingCostService: mockService,
			})

			data, _ := json.Marshal(requestData)

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/shippingcost", bytes.NewReader(data))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)

		})
	}

	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockShippingCostService(mockCtrl)

	router := app.SetupRouter(app.ShippingHandler{
		ShippingCostService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/shippingcost", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestHandleGetShippingCostRecordByID(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockShippingCostService)
		expected   int
	}{
		{
			name: "SuccessGetShippingCostRecordByID",
			createStub: func(mps *mocks.MockShippingCostService) {
				mps.EXPECT().GetShippingCostByCity("xyx" /* id */).Return(&domain.ShippingCost{}, nil)
			},
			expected: 202,
		},
		{
			name: "FailGetShippingCostRecordByID",
			createStub: func(mps *mocks.MockShippingCostService) {
				mps.EXPECT().GetShippingCostByCity("xyx" /* id */).Return(nil, &errs.AppError{Message: "errstring"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockShippingCostService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.ShippingHandler{
				ShippingCostService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/shippingcost/xyx", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}

func TestHandleUpdateShippingCostStatus(t *testing.T) {

	requestData := domain.ShippingCost{
		City:         "Banglore",
		ShippingCost: 0,
	}
	testcases := []struct {
		name       string
		createStub func(mocks.MockShippingCostService)
		expected   int
	}{
		{
			name: "SuccesshandleUpdateShippingCostStatus",
			createStub: func(mps mocks.MockShippingCostService) {
				mps.EXPECT().UpdateShippingCost(requestData).Return(true, nil)
			},
			expected: 202,
		},
		{
			name: "FailurehandleUpdateShippingCostStatus",
			createStub: func(mps mocks.MockShippingCostService) {
				mps.EXPECT().UpdateShippingCost(requestData).Return(false, &errs.AppError{Message: "error"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockShippingCostService(mockCtrl)
			testcase.createStub(*mockService)

			router := app.SetupRouter(app.ShippingHandler{
				ShippingCostService: mockService,
			})

			recorder := httptest.NewRecorder()

			data, _ := json.Marshal(requestData)

			req := httptest.NewRequest(http.MethodPut, "/shippingcost", bytes.NewReader(data))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}

	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockShippingCostService(mockCtrl)

	router := app.SetupRouter(app.ShippingHandler{
		ShippingCostService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/shippingcost", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestHandleUpdateDeleteShippingCostByID(t *testing.T) {
	testcases := []struct {
		name       string
		createStub func(mocks.MockShippingCostService)
		expected   int
	}{
		{
			name: "SuccesshandleUpdateShippingCostStatus",
			createStub: func(mps mocks.MockShippingCostService) {
				mps.EXPECT().DeleteShippingCostByCity("xyz" /* Id */).Return(true, nil)
			},
			expected: 202,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockShippingCostService(mockCtrl)
			testcase.createStub(*mockService)

			router := app.SetupRouter(app.ShippingHandler{
				ShippingCostService: mockService,
			})

			recshipping := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodDelete, "/shippingcost/xyz", nil)
			router.ServeHTTP(recshipping, req)

			assert.Equal(t, testcase.expected, recshipping.Code)
		})
	}

}
