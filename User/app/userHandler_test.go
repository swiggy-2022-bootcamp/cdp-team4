package app_test

import (
	// "bytes"
	// "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/user/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/app"
)


/*
func TestCreateUser(t *testing.T) {

	shippingAddress := app.ShippingAddressDTO{
		FirstName:		"Swastik",
		LastName:		"Sahoo",
		City:			"swastik15",
		Address1:		"abc",
		Address2:		"95181818181",
		CountryID:		12,
		PostCode:		234,
	}

	requestData := app.UserDTO{
		FirstName:		"Swastik",
		LastName:		"Sahoo",
		Username:		"swastik15",
		Password:		"abc",
		Phone:			"95181818181",
		Email:			"swastiksahoo22@gmail.com",
		Role:			0,
		Fax:			"12-203-9181",
		Address:		shippingAddress,
	}

	testCases := []struct {
		name       string
		createStub func(*mocks.MockUserService, *mocks.MockGrpcHelper)
		expected   int
	}{
		{
			name: "SucessCreateUser",
			createStub: func(mus *mocks.MockUserService, mgs *mocks.MockGrpcHelper) {

				mgs.EXPECT().GetShippingAddressId(requestData.Address).Return("abcid")
				mus.EXPECT().
					CreateUserInDynamodb(
						requestData.FirstName,
						requestData.LastName,
						requestData.Username,
						requestData.Phone,
						requestData.Email,
						requestData.Password,
						domain.Admin,
						"abcid",
						requestData.Fax).
					Return(domain.User{}, nil)
				
			},
			expected: 201,
		},
		{
			name: "ErrorCreateRole",
			createStub: func(mus *mocks.MockUserService) {
				mus.EXPECT().
					CreateUserInDynamodb(
						requestData.FirstName,
						requestData.LastName,
						requestData.Username,
						requestData.Phone,
						requestData.Email,
						requestData.Password,
						domain.Admin,
						"abcid",
						requestData.Fax).
					Return(domain.User{}, nil)
			},
			expected: 201,
		},
		{
			name: "ErrorCreateUser",
			createStub: func(mus *mocks.MockUserService) {
				mus.EXPECT().
					CreateUserInDynamodb(
						requestData.FirstName,
						requestData.LastName,
						requestData.Username,
						requestData.Phone,
						requestData.Email,
						requestData.Password,
						domain.Admin,
						"abcid",
						requestData.Fax).
					Return(domain.User{}, fmt.Errorf("unable to insert record"))
			},
			expected: 500,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockUserService(mockCtrl)
			mockGrpc := mocks.NewMockGrpcHelper(mockCtrl)
			testcase.createStub(mockService, mockGrpc)


			router := app.SetupRouter(app.UserHandler{
				UserService: mockService,
			})

			data, _ := json.Marshal(requestData)

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(data))
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}

	// FailBindJSON
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockUserService(mockCtrl)

	router := app.SetupRouter(app.UserHandler{
		UserService: mockService,
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/user", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

*/

func TestHandleGetUserByID(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockUserService)
		expected   int
	}{
		{
			name: "SuccessGetUserByID",
			createStub: func(mus *mocks.MockUserService) {
				mus.EXPECT().
					GetUserById("xyx" /* id */).
					Return(&domain.User{}, nil)
			},
			expected: 200,
		},
		{
			name: "FailGetUserByID",
			createStub: func(mus *mocks.MockUserService) {
				mus.EXPECT().
					GetUserById("xyx" /* id */).
					Return(nil, fmt.Errorf("user not found"))
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockUserService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.UserHandler{
				UserService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/user/xyx", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}




func TestHandleGetAllUsers(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockUserService)
		expected   int
	}{
		{
			name: "SuccessGetAllUsers",
			createStub: func(mus *mocks.MockUserService) {
				mus.EXPECT().
					GetAllUsers().
					Return([]domain.User{}, nil)
			},
			expected: 200,
		},
		{
			name: "FailGetAllUsers",
			createStub: func(mus *mocks.MockUserService) {
				mus.EXPECT().
					GetAllUsers().
					Return(nil, fmt.Errorf("users not found"))
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockUserService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.UserHandler{
				UserService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}


func TestHandleDeleteUserById(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(*mocks.MockUserService)
		expected   int
	}{
		{
			name: "SuccessDeleteUserById",
			createStub: func(mus *mocks.MockUserService) {
				mus.EXPECT().
					DeleteUserById("xyx").
					Return(true, nil)
			},
			expected: 202,
		},
		{
			name: "FailDeleteUserById",
			createStub: func(mus *mocks.MockUserService) {
				mus.EXPECT().
					DeleteUserById("xyx").
					Return(false, fmt.Errorf("users not found"))
			},
			expected: 400,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockUserService(mockCtrl)
			testcase.createStub(mockService)

			router := app.SetupRouter(app.UserHandler{
				UserService: mockService,
			})

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/user/xyx", nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testcase.expected, recorder.Code)
		})
	}
}

