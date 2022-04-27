package app_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/app"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/utils/errs"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAuthToken(t *testing.T) {
	mockAuthService := mocks.NewAuthService(t)

	authHandler := app.AuthHandler{
		AuthService: mockAuthService,
	}

	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/login", authHandler.GetAuthToken)

	requestData := app.LoginDTO{
		Username: "abc",
		Password: "abc",
	}

	data, _ := json.Marshal(requestData)

	expectedResponse := app.ResponseDTO{
		Status: http.StatusOK,
		Data:   "test-auth-token",
	}

	mockAuthService.On("GenerateAuthToken", requestData.Username, requestData.Password).Return("test-auth-token", nil)
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewReader(data))
	router.ServeHTTP(recorder, req)

	assert.Equal(t, expectedResponse.Status, recorder.Code)
}

func TestGetAuthTokenWithInvalidRequestPayload(t *testing.T) {

	mockAuthService := mocks.NewAuthService(t)

	authHandler := app.AuthHandler{
		AuthService: mockAuthService,
	}

	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/login", authHandler.GetAuthToken)

	data := "{"

	customErr := errs.NewValidationError("Invalid request paylaod")
	expectedResponse := app.ResponseDTO{
		Status:  customErr.Code,
		Message: customErr.Message,
	}

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewReader([]byte(data)))
	router.ServeHTTP(recorder, req)

	assert.Equal(t, expectedResponse.Status, recorder.Code)
}

func TestValidateAuthToken(t *testing.T) {
	mockAuthService := mocks.NewAuthService(t)

	authHandler := app.AuthHandler{
		AuthService: mockAuthService,
	}

	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/validate", authHandler.ValidateAuthToken)

	testAuthToken := "test.auth.token"
	mockAuthModel := domain.AuthModel{
		UserId:    "test-user-id",
		Role:      1,
		AuthToken: testAuthToken,
		IsExpired: false,
		ExpiresOn: time.Now().Add(5 * time.Minute),
	}

	expectedResponse, _ := json.Marshal(app.ResponseDTO{
		Status:  http.StatusOK,
		Message: "Access Granted",
		Data: app.ValidationDTO{
			UserId: mockAuthModel.UserId,
			Role:   mockAuthModel.Role,
		},
	})
	mockAuthService.On("ValidateAuthToken", testAuthToken).Return(&mockAuthModel, nil)

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/validate", nil)
	req.Header.Set("Authorization", "Bearer "+testAuthToken)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, expectedResponse, recorder.Body.Bytes())
}

func TestValidateAuthTokenWithInvalidToken(t *testing.T) {
	mockAuthService := mocks.NewAuthService(t)

	authHandler := app.AuthHandler{AuthService: mockAuthService}

	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/validate", authHandler.ValidateAuthToken)

	testAuthToken := "test.auth"

	customErr := errs.NewAuthenticationError("Invalid token, Access Denied")

	expectedResponse, _ := json.Marshal(app.ResponseDTO{
		Status:  customErr.Code,
		Message: customErr.Message,
	})

	mockAuthService.On("ValidateAuthToken", testAuthToken).Return(nil, errs.NewValidationError("some error"))

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/validate", nil)
	req.Header.Set("Authorization", "Bearer "+testAuthToken)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, expectedResponse, recorder.Body.Bytes())
}

func TestInvalidateAuthToken(t *testing.T) {
	mockAuthService := mocks.NewAuthService(t)

	authHandler := app.AuthHandler{
		AuthService: mockAuthService,
	}

	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.POST("/logout", authHandler.InvalidateAuthToken)

	testAuthToken := "test.auth.token"
	mockAuthService.On("InvalidateAuthToken", testAuthToken).Return(nil)
	mockAuthService.On("ValidateAuthToken", testAuthToken).Return(nil, nil)

	expectedResponse, _ := json.Marshal(app.ResponseDTO{
		Status:  http.StatusOK,
		Message: "Logged out successfully",
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/logout", nil)
	req.Header.Set("Authorization", "Bearer "+testAuthToken)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, expectedResponse, recorder.Body.Bytes())

}

func TestInvalidateAuthTokenWithInvalidToken(t *testing.T) {
	mockAuthService := mocks.NewAuthService(t)

	authHandler := app.AuthHandler{AuthService: mockAuthService}

	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/logout", authHandler.InvalidateAuthToken)

	testAuthToken := "test.auth"

	customErr := errs.NewAuthenticationError("Invalid token, Access Denied")

	expectedResponse, _ := json.Marshal(app.ResponseDTO{
		Status:  customErr.Code,
		Message: customErr.Message,
	})

	mockAuthService.On("ValidateAuthToken", testAuthToken).Return(nil, errs.NewValidationError("some error"))

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/logout", nil)
	req.Header.Set("Authorization", "Bearer "+testAuthToken)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, expectedResponse, recorder.Body.Bytes())
}

func TestInvalidateAuthTokenWithInvalidUserId(t *testing.T) {
	mockAuthService := mocks.NewAuthService(t)

	authHandler := app.AuthHandler{AuthService: mockAuthService}

	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/logout", authHandler.InvalidateAuthToken)

	testAuthToken := "test.auth"

	customErr := errs.NewValidationError("Invalid token, Access Denied")

	expectedResponse, _ := json.Marshal(app.ResponseDTO{
		Status:  customErr.Code,
		Message: customErr.Message,
	})

	mockAuthService.On("ValidateAuthToken", testAuthToken).Return(nil, nil)
	mockAuthService.On("InvalidateAuthToken", testAuthToken).Return(errs.NewValidationError("Invalid token, Access Denied"))

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/logout", nil)
	req.Header.Set("Authorization", "Bearer "+testAuthToken)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, expectedResponse, recorder.Body.Bytes())
}
