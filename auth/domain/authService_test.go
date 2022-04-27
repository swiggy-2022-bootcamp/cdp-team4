package domain_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/utils/errs"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"testing"
	"time"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func TestShouldGenerateAuthToken(t *testing.T) {
	var mockUserRepo = mocks.UserRepository{}
	var mockAuthRepo = mocks.AuthRepository{}
	var authServiceTest = domain.NewAuthService(&mockUserRepo, &mockAuthRepo)

	username := "abc"
	password := "123"

	mockUserModel := &domain.UserModel{
		UserId:   "aabbcchhjjkkk",
		Role:     1,
		Username: username,
		Password: HashPassword(password),
	}

	mockUserRepo.On("FindByUsername", username).Return(mockUserModel, nil)
	mockAuthRepo.On("Save", mock.Anything).Return(nil)
	authToken, _ := authServiceTest.GenerateAuthToken(username, password)
	assert.NotEmpty(t, authToken)
}

func TestShouldReturnErrorWhileGeneratingTokenForInvalidPassword(t *testing.T) {
	var mockUserRepo = mocks.UserRepository{}
	var mockAuthRepo = mocks.AuthRepository{}
	var authServiceTest = domain.NewAuthService(&mockUserRepo, &mockAuthRepo)

	username := "abc"
	password := "123"

	mockUserModel := &domain.UserModel{
		UserId:   "aabbcchhjjkkk",
		Role:     1,
		Username: username,
		Password: "wrong password",
	}

	mockUserRepo.On("FindByUsername", username).Return(mockUserModel, nil)
	mockAuthRepo.On("Save", mock.Anything).Return(nil)
	authToken, err := authServiceTest.GenerateAuthToken(username, password)
	assert.Empty(t, authToken)
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, http.StatusUnauthorized)
}

func TestShouldReturnErrorForInvalidUsername(t *testing.T) {
	var mockUserRepo = mocks.UserRepository{}
	var mockAuthRepo = mocks.AuthRepository{}
	var authServiceTest = domain.NewAuthService(&mockUserRepo, &mockAuthRepo)

	username := "abc"
	password := "123"

	mockUserRepo.On("FindByUsername", username).Return(nil, errs.NewUnexpectedError("Username not found"))
	mockAuthRepo.On("Save", mock.Anything).Return(nil)
	authToken, err := authServiceTest.GenerateAuthToken(username, password)
	assert.Empty(t, authToken)
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, http.StatusInternalServerError)
}

func TestShouldValidateAuthToken(t *testing.T) {
	var mockUserRepo = mocks.UserRepository{}
	var mockAuthRepo = mocks.AuthRepository{}
	var authServiceTest = domain.NewAuthService(&mockUserRepo, &mockAuthRepo)

	username := "abc"
	password := "123"

	mockUserModel := &domain.UserModel{
		UserId:   "aabbcchhjjkkk",
		Role:     1,
		Username: username,
		Password: HashPassword(password),
	}

	mockUserRepo.On("FindByUsername", username).Return(mockUserModel, nil)
	mockAuthRepo.On("Save", mock.Anything).Return(nil)
	authToken, _ := authServiceTest.GenerateAuthToken(username, password)

	mockAuthModel := domain.AuthModel{
		UserId:    "absconds",
		Role:      1,
		AuthToken: authToken,
		ExpiresOn: time.Now().Add(5 * time.Minute),
	}

	mockAuthRepo.On("FindByAuthToken", authToken).Return(&mockAuthModel, nil)

	actualAuthModel, err := authServiceTest.ValidateAuthToken(authToken)
	assert.Nil(t, err)
	assert.Equal(t, mockAuthModel, *actualAuthModel)
}

func TestShouldReturnErrorForInvalidAuthToken(t *testing.T) {
	var mockUserRepo = mocks.UserRepository{}
	var mockAuthRepo = mocks.AuthRepository{}
	var authServiceTest = domain.NewAuthService(&mockUserRepo, &mockAuthRepo)

	authToken := "abc.abc.abc"
	mockAuthModel := domain.AuthModel{
		UserId:    "absconds",
		Role:      1,
		AuthToken: authToken,
		ExpiresOn: time.Now().Add(5 * time.Minute),
	}

	mockAuthRepo.On("FindByAuthToken", authToken).Return(&mockAuthModel, nil)

	actualAuthModel, err := authServiceTest.ValidateAuthToken(authToken)
	assert.NotNil(t, err)
	assert.Nil(t, actualAuthModel)
}

func TestShouldReturnErrorForAuthTokenNotFound(t *testing.T) {
	var mockUserRepo = mocks.UserRepository{}
	var mockAuthRepo = mocks.AuthRepository{}
	var authServiceTest = domain.NewAuthService(&mockUserRepo, &mockAuthRepo)

	username := "abc"
	password := "123"

	mockUserModel := &domain.UserModel{
		UserId:   "aabbcchhjjkkk",
		Role:     1,
		Username: username,
		Password: HashPassword(password),
	}

	mockUserRepo.On("FindByUsername", username).Return(mockUserModel, nil)
	mockAuthRepo.On("Save", mock.Anything).Return(nil)
	authToken, _ := authServiceTest.GenerateAuthToken(username, password)

	mockAuthRepo.On("FindByAuthToken", authToken).Return(nil, errs.NewNotFoundError("Auth Token not found"))

	actualAuthModel, err := authServiceTest.ValidateAuthToken(authToken)
	assert.NotNil(t, err)
	assert.Nil(t, actualAuthModel)
}

func TestShouldInvalidateAuthToken(t *testing.T) {
	var mockUserRepo = mocks.UserRepository{}
	var mockAuthRepo = mocks.AuthRepository{}
	var authServiceTest = domain.NewAuthService(&mockUserRepo, &mockAuthRepo)

	username := "abc"
	password := "123"

	mockUserModel := &domain.UserModel{
		UserId:   "aabbcchhjjkkk",
		Role:     1,
		Username: username,
		Password: HashPassword(password),
	}

	mockUserRepo.On("FindByUsername", username).Return(mockUserModel, nil)
	mockAuthRepo.On("Save", mock.Anything).Return(nil)
	authToken, _ := authServiceTest.GenerateAuthToken(username, password)

	mockAuthModel := domain.AuthModel{
		UserId:    "absconds",
		Role:      1,
		AuthToken: authToken,
		ExpiresOn: time.Now().Add(5 * time.Minute),
	}

	mockAuthRepo.On("FindByAuthToken", authToken).Return(&mockAuthModel, nil)
	mockAuthRepo.On("Save", mock.Anything).Return(nil)
	err := authServiceTest.InvalidateAuthToken(authToken)
	assert.Nil(t, err)
}
