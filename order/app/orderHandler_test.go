package app_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/app"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/order/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/utils/errs"
)

func TestHandleOrder(t *testing.T) {
	// router := gin.Default()

	requestData := app.OrderRecordDTO{
		UserID:  "9123790217093210",
		OrderID: "9123793213232312",
		Status:  "Confirmed",
		Products: []app.ProductRecordDTO{
			{
				Product:  "Trimax",
				Quantity: 10,
				Cost:     90,
			},
		},
		TotalCost: 1000,
	}

	prod_qt, prod_ct := app.ConvertProductsDTOtoMaps(requestData.Products)

	testCases := []struct {
		name       string
		createStub func(*mocks.MockOrderService)
		expected   int
	}{
		{
			name: "SucessHandleOrder",
			createStub: func(mps *mocks.MockOrderService) {
				response := "9123793213232312"

				mps.EXPECT().
					CreateOrder(gomock.Any(), requestData.Status, prod_qt, prod_ct, int(requestData.TotalCost)).
					Return(response, nil)
			},
			expected: 202,
		},
		{
			name: "ErrorCreateOrderRecord",
			createStub: func(mps *mocks.MockOrderService) {
				mps.EXPECT().
					CreateOrder(gomock.Any(), requestData.Status, prod_qt, prod_ct, int(requestData.TotalCost)).
					Return("", &errs.AppError{Message: "Unable to insert record"})
			},
			expected: 400,
		},
		{
			name: "ErrorNothingReturnFromCreateOrderRecord",
			createStub: func(mps *mocks.MockOrderService) {
				mps.EXPECT().
					CreateOrder(gomock.Any(), requestData.Status, prod_qt, prod_ct, int(requestData.TotalCost)).
					Return("", &errs.AppError{Message: "Unable to insert record"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockOrderService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.OrderHandler{
				OrderService: mockService,
			})

			data, _ := json.Marshal(requestData)

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewReader(data))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)

		})
	}

	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockOrderService(mockCtrl)

	router := app.SetupRouter(app.OrderHandler{
		OrderService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/order", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestHandleGetOrderRecordByID(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockOrderService)
		expected   int
	}{
		{
			name: "SuccessGetOrderRecordByID",
			createStub: func(mps *mocks.MockOrderService) {
				mps.EXPECT().GetOrderById("xyx" /* id */).Return(&domain.Order{}, nil)
			},
			expected: 202,
		},
		{
			name: "FailGetOrderRecordByID",
			createStub: func(mps *mocks.MockOrderService) {
				mps.EXPECT().GetOrderById("xyx" /* id */).Return(nil, &errs.AppError{Message: "errstring"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockOrderService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.OrderHandler{
				OrderService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/order/xyx", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}

func TestHandleGetOrderRecordByUserID(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockOrderService)
		expected   int
	}{
		{
			name: "SuccessGetOrderRecordByUserID",
			createStub: func(mps *mocks.MockOrderService) {
				mps.EXPECT().GetOrderByUserId("xyx" /* id */).Return([]domain.Order{}, nil)
			},
			expected: 202,
		},
		{
			name: "FailGetOrderRecordByUserID",
			createStub: func(mps *mocks.MockOrderService) {
				mps.EXPECT().GetOrderByUserId("xyx" /* id */).Return(nil, &errs.AppError{Message: "errstring"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockOrderService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.OrderHandler{
				OrderService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/order/user/xyx", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}

func TestHandleGetOrderRecordByStatus(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockOrderService)
		expected   int
	}{
		{
			name: "SuccessGetOrderRecordByUserID",
			createStub: func(mps *mocks.MockOrderService) {
				mps.EXPECT().GetOrderByStatus("xyx" /* id */).Return([]domain.Order{}, nil)
			},
			expected: 202,
		},
		{
			name: "FailGetOrderRecordByUserID",
			createStub: func(mps *mocks.MockOrderService) {
				mps.EXPECT().GetOrderByStatus("xyx" /* id */).Return(nil, &errs.AppError{Message: "errstring"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockOrderService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.OrderHandler{
				OrderService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/order/status/xyx", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}

func TestHandleGetAllRecords(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockOrderService)
		expected   int
	}{
		{
			name: "SuccessGetOrderRecordByUserID",
			createStub: func(mps *mocks.MockOrderService) {
				mps.EXPECT().GetAllOrders().Return([]domain.Order{}, nil)
			},
			expected: 202,
		},
		{
			name: "FailGetOrderRecordByUserID",
			createStub: func(mps *mocks.MockOrderService) {
				mps.EXPECT().GetAllOrders().Return(nil, &errs.AppError{Message: "errstring"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockOrderService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.OrderHandler{
				OrderService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/orders", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}

func TestHandleUpdateOrderStatus(t *testing.T) {
	testcases := []struct {
		name       string
		createStub func(mocks.MockOrderService)
		expected   int
	}{
		{
			name: "SuccesshandleUpdateOrderStatus",
			createStub: func(mps mocks.MockOrderService) {
				mps.EXPECT().UpdateOrderStatus("xyx" /* Id */, "confirmed" /* Status */).Return(true, nil)
			},
			expected: 202,
		},
		{
			name: "FailurehandleUpdateOrderStatus",
			createStub: func(mps mocks.MockOrderService) {
				mps.EXPECT().UpdateOrderStatus("xyx" /* Id */, "confirmed" /* Status */).Return(false, &errs.AppError{Message: "error"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockOrderService(mockCtrl)
			testcase.createStub(*mockService)

			router := app.SetupRouter(app.OrderHandler{
				OrderService: mockService,
			})

			recorder := httptest.NewRecorder()
			type requestDTO struct {
				Id     string `json:"id"`
				Status string `json:"status"`
			}

			requestData, _ := json.Marshal(requestDTO{Id: "xyx", Status: "confirmed"})
			req := httptest.NewRequest(http.MethodPut, "/order/status", bytes.NewReader(requestData))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}

	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockOrderService(mockCtrl)

	router := app.SetupRouter(app.OrderHandler{
		OrderService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/order/status", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestHandleUpdateDeleteOrderByID(t *testing.T) {
	testcases := []struct {
		name       string
		createStub func(mocks.MockOrderService)
		expected   int
	}{
		{
			name: "SuccesshandleDeleteOrderByID",
			createStub: func(mps mocks.MockOrderService) {
				mps.EXPECT().DeleteOrderById("xyx" /* Id */).Return(true, nil)
			},
			expected: 202,
		},
		{
			name: "FailurehandleDeleteOrderByID",
			createStub: func(mps mocks.MockOrderService) {
				mps.EXPECT().DeleteOrderById("xyx" /* Id */).Return(false, &errs.AppError{Message: "errstring"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockOrderService(mockCtrl)
			testcase.createStub(*mockService)

			router := app.SetupRouter(app.OrderHandler{
				OrderService: mockService,
			})

			recorder := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodDelete, "/order/xyx", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}

}

func TestShouldReturnNewOrderHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockOrderService(mockCtrl)
	orderHandler := app.NewOrderHandler(mockService, nil)
	assert.NotNil(t, orderHandler)
}

func TestConvertProductsDTOtoMaps(t *testing.T) {
	testproductdto := []app.ProductRecordDTO{
		{
			Product:  "Pen",
			Cost:     10,
			Quantity: 4,
		}, {
			Product:  "Pen",
			Cost:     12,
			Quantity: 5,
		},
	}
	res, res1 := app.ConvertProductsDTOtoMaps(testproductdto)
	assert.NotNil(t, res)
	assert.NotNil(t, res1)
}
