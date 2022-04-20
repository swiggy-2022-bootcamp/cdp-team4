package app_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/app"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/payment/mocks"
)

func TestHandlePay(t *testing.T) {
	// router := gin.Default()

	requestData := app.PaymentRecordDTO{
		Amount:      45,
		Currency:    "INR",
		Status:      "1",
		OrderID:     "abc",
		UserID:      "xyz",
		Method:      "credit-card",
		Description: "dummy desc",
		VPA:         "zy",
		Notes:       []string{""},
	}

	testCases := []struct {
		name       string
		createStub func(*mocks.MockPaymentService)
		expected   int
	}{
		{
			name: "SucessHandlePay",
			createStub: func(mps *mocks.MockPaymentService) {
				response := map[string]interface{}{"name": "dummy_name"}

				mps.EXPECT().
					CreateDynamoPaymentRecord(gomock.Any(), requestData.Amount, requestData.Currency, requestData.Status, requestData.OrderID, requestData.UserID, requestData.Method, requestData.Description, requestData.VPA, requestData.Notes).
					Return(response, nil)
			},
			expected: 200,
		},
		{
			name: "ErrorCreateDynamoPaymentRecord",
			createStub: func(mps *mocks.MockPaymentService) {
				mps.EXPECT().
					CreateDynamoPaymentRecord(gomock.Any(), requestData.Amount, requestData.Currency, requestData.Status, requestData.OrderID, requestData.UserID, requestData.Method, requestData.Description, requestData.VPA, requestData.Notes).
					Return(nil, fmt.Errorf("unable to insert record"))
			},
			expected: 400,
		},
		{
			name: "ErrorNothingReturnFromCreateDynamoPaymentRecord",
			createStub: func(mps *mocks.MockPaymentService) {
				mps.EXPECT().
					CreateDynamoPaymentRecord(gomock.Any(), requestData.Amount, requestData.Currency, requestData.Status, requestData.OrderID, requestData.UserID, requestData.Method, requestData.Description, requestData.VPA, requestData.Notes).
					Return(nil, nil)
			},
			expected: 400,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockPaymentService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.PayHandler{
				PaymentService: mockService,
			})

			data, _ := json.Marshal(requestData)

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/pay/", bytes.NewReader(data))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)

		})
	}

	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockPaymentService(mockCtrl)

	router := app.SetupRouter(app.PayHandler{
		PaymentService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/pay/", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestHandleGetPayRecordByID(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockPaymentService)
		expected   int
	}{
		{
			name: "SuccessGetPayRecordByID",
			createStub: func(mps *mocks.MockPaymentService) {
				mps.EXPECT().GetPaymentRecordById("xyx" /* id */).Return(&domain.Payment{}, nil)
			},
			expected: 200,
		},
		{
			name: "FailGetPayRecordByID",
			createStub: func(mps *mocks.MockPaymentService) {
				mps.EXPECT().GetPaymentRecordById("xyx" /* id */).Return(nil, fmt.Errorf("record not found"))
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockPaymentService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.PayHandler{
				PaymentService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/pay/xyx", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}

func TestHandleUpdatePayStatus(t *testing.T) {
	testcases := []struct {
		name       string
		createStub func(mocks.MockPaymentService)
		expected   int
	}{
		{
			name: "SuccesshandleUpdatePayStatus",
			createStub: func(mps mocks.MockPaymentService) {
				mps.EXPECT().UpdatePaymentStatus("xyx" /* id */, "confirmed" /* status */).Return(true, nil)
			},
			expected: 200,
		},
		{
			name: "FailurehandleUpdatePayStatus",
			createStub: func(mps mocks.MockPaymentService) {
				mps.EXPECT().UpdatePaymentStatus("xyx" /* id */, "confirmed" /* status */).Return(false, fmt.Errorf("unable to update the status"))
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockPaymentService(mockCtrl)
			testcase.createStub(*mockService)

			router := app.SetupRouter(app.PayHandler{
				PaymentService: mockService,
			})

			requestData, _ := json.Marshal(app.UpdatePayStatusDTO{Id: "xyx", Status: "confirmed"})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/pay/", bytes.NewReader(requestData))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}

	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockPaymentService(mockCtrl)

	router := app.SetupRouter(app.PayHandler{
		PaymentService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/pay/", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestHandleAddPaymentMethods(t *testing.T) {
	testcases := []struct {
		name       string
		createStub func(mocks.MockPaymentService)
		expected   int
	}{
		{
			name: "SuccessHandleAddPaymentMethods",
			createStub: func(mps mocks.MockPaymentService) {
				mps.EXPECT().AddPaymentMethod("xyx", "credit-card", "1", "none").Return(true, nil)
			},
			expected: 200,
		},
		{
			name: "SuccessHandleAddPaymentMethods",
			createStub: func(mps mocks.MockPaymentService) {
				mps.EXPECT().AddPaymentMethod("xyx", "credit-card", "1", "none").Return(false, fmt.Errorf("unable to add"))

			},
			expected: 400,
		},
		{
			name: "SuccessHandleAddPaymentMethods",
			createStub: func(mps mocks.MockPaymentService) {
				mps.EXPECT().AddPaymentMethod("xyx", "credit-card", "1", "none").Return(false, fmt.Errorf("ConditionalCheckFailedException: unable to add"))

			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockPaymentService(mockCtrl)
			testcase.createStub(*mockService)

			router := app.SetupRouter(app.PayHandler{
				PaymentService: mockService,
			})

			requestData, _ := json.Marshal(app.PaymentMethodDTO{
				Id:      "xyx",
				Method:  "credit-card",
				Agree:   "1",
				Comment: "none",
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/pay/paymentMethods", bytes.NewReader(requestData))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}

	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockPaymentService(mockCtrl)

	router := app.SetupRouter(app.PayHandler{
		PaymentService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/pay/paymentMethods", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestHandleGetPaymentMethods(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(mocks.MockPaymentService)
		expected   int
	}{
		{
			name: "SuccessHandleAddPaymentMethods",
			createStub: func(mps mocks.MockPaymentService) {
				mps.EXPECT().GetPaymentMethods("xyx").Return([]string{"debit-card"}, nil)
			},
			expected: 200,
		},
		{
			name: "SuccessHandleAddPaymentMethods",
			createStub: func(mps mocks.MockPaymentService) {
				mps.EXPECT().GetPaymentMethods("xyx").Return(nil, fmt.Errorf("unable to get methods"))
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockPaymentService(mockCtrl)
			testcase.createStub(*mockService)

			router := app.SetupRouter(app.PayHandler{
				PaymentService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/pay/paymentMethods/xyx", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}

}
