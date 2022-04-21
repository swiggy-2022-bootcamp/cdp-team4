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

func TestHandleShippingAddress(t *testing.T) {
	// router := gin.Default()

	requestData := app.ShippingAddressRecordDTO{
		FirstName: "Naveen",
		LastName:  "Kumar",
		City:      "Banglore",
		Address1:  "address1",
		Address2:  "address2",
		CountryID: 56,
		PostCode:  454645,
	}

	testCases := []struct {
		name       string
		createStub func(*mocks.MockShippingAddressService)
		expected   int
	}{
		{
			name: "SucessHandleShippingAddress",
			createStub: func(mps *mocks.MockShippingAddressService) {
				response := "9123793213232312"

				mps.EXPECT().
					CreateShippingAddress(gomock.Any(), gomock.Any(), gomock.Any(),
						gomock.Any(), gomock.Any(), requestData.CountryID, requestData.PostCode).
					Return(response, nil)
			},
			expected: 202,
		},
		{
			name: "ErrorCreateShippingAddressRecord",
			createStub: func(mps *mocks.MockShippingAddressService) {
				mps.EXPECT().
					CreateShippingAddress(gomock.Any(), gomock.Any(), gomock.Any(),
						gomock.Any(), gomock.Any(), requestData.CountryID, requestData.PostCode).
					Return("", &errs.AppError{Message: "Unable to insert record"})
			},
			expected: 400,
		},
		{
			name: "ErrorNothingReturnFromCreateShippingAddressRecord",
			createStub: func(mps *mocks.MockShippingAddressService) {
				mps.EXPECT().
					CreateShippingAddress(gomock.Any(), gomock.Any(), gomock.Any(),
						gomock.Any(), gomock.Any(), requestData.CountryID, requestData.PostCode).
					Return("", &errs.AppError{Message: "Unable to insert record"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockShippingAddressService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.ShippingHandler{
				ShippingAddressService: mockService,
			})

			data, _ := json.Marshal(requestData)

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/shippingaddress", bytes.NewReader(data))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)

		})
	}

	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockShippingAddressService(mockCtrl)

	router := app.SetupRouter(app.ShippingHandler{
		ShippingAddressService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/shippingaddress", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestHandleGetShippingAddressRecordByID(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockShippingAddressService)
		expected   int
	}{
		{
			name: "SuccessGetShippingAddressRecordByID",
			createStub: func(mps *mocks.MockShippingAddressService) {
				mps.EXPECT().GetShippingAddressById("xyx" /* id */).Return(&domain.ShippingAddress{}, nil)
			},
			expected: 202,
		},
		{
			name: "FailGetShippingAddressRecordByID",
			createStub: func(mps *mocks.MockShippingAddressService) {
				mps.EXPECT().GetShippingAddressById("xyx" /* id */).Return(nil, &errs.AppError{Message: "errstring"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockShippingAddressService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.ShippingHandler{
				ShippingAddressService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/shippingaddress/xyx", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}

func TestHandleUpdateShippingAddressStatus(t *testing.T) {

	requestData := app.ShippingAddressRecordDTO{
		FirstName: "Naveen",
		LastName:  "Kumar",
		City:      "Banglore",
		Address1:  "address1",
		Address2:  "address2",
		CountryID: 56,
		PostCode:  454645,
	}
	reqdomain := domain.ShippingAddress{
		FirstName: requestData.FirstName,
		LastName:  requestData.LastName,
		City:      requestData.City,
		Address1:  requestData.Address1,
		Address2:  requestData.Address2,
		PostCode:  requestData.PostCode,
		CountryID: requestData.CountryID,
	}
	testcases := []struct {
		name       string
		createStub func(mocks.MockShippingAddressService)
		expected   int
	}{
		{
			name: "SuccesshandleUpdateShippingAddressStatus",
			createStub: func(mps mocks.MockShippingAddressService) {
				mps.EXPECT().UpdateShippingAddressById("xyz",
					reqdomain).Return(true, nil)
			},
			expected: 202,
		},
		{
			name: "FailurehandleUpdateShippingAddressStatus",
			createStub: func(mps mocks.MockShippingAddressService) {
				mps.EXPECT().UpdateShippingAddressById("xyz", reqdomain).Return(false, &errs.AppError{Message: "error"})
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockShippingAddressService(mockCtrl)
			testcase.createStub(*mockService)

			router := app.SetupRouter(app.ShippingHandler{
				ShippingAddressService: mockService,
			})

			recorder := httptest.NewRecorder()

			data, _ := json.Marshal(reqdomain)

			req := httptest.NewRequest(http.MethodPut, "/shippingaddress/xyz", bytes.NewReader(data))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}

	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockShippingAddressService(mockCtrl)

	router := app.SetupRouter(app.ShippingHandler{
		ShippingAddressService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/shippingaddress/xyz", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestHandleUpdateDeleteShippingAddressByID(t *testing.T) {
	testcases := []struct {
		name       string
		createStub func(mocks.MockShippingAddressService)
		expected   int
	}{
		{
			name: "SuccesshandleUpdateShippingAddressStatus",
			createStub: func(mps mocks.MockShippingAddressService) {
				mps.EXPECT().DeleteShippingAddressById("xyz" /* Id */).Return(true, nil)
			},
			expected: 202,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockShippingAddressService(mockCtrl)
			testcase.createStub(*mockService)

			router := app.SetupRouter(app.ShippingHandler{
				ShippingAddressService: mockService,
			})

			recshipping := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodDelete, "/shippingaddress/xyz", nil)
			router.ServeHTTP(recshipping, req)

			assert.Equal(t, testcase.expected, recshipping.Code)
		})
	}

}
