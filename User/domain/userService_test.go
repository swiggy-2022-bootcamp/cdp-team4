package domain_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/mocks"
	// "github.com/swiggy-2022-bootcamp/cdp-team4/user/utils/errs"

	"github.com/stretchr/testify/assert"
)

var mockUserRepo = mocks.UserDynamoDBRepository{}
var userService = domain.NewUserService(&mockUserRepo)

func TestShouldReturnNewUserService(t *testing.T) {
	userService := domain.NewUserService(nil)
	assert.NotNil(t, userService)
}

func TestShouldCreateNewUser(t *testing.T) {
	userID := "afshsjgj14151jou"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password, _ := domain.HashPassword("Pass!23")
	role := domain.Admin
	addressID := "addressid"
	fax := "18-1918-10101"


	user := domain.NewUser(userID, firstName, lastName, username, phone, email, password, role, addressID, fax)
	mockUserRepo.On("Save", mock.Anything).Return(*user, nil)
	userService.CreateUserInDynamodb(firstName, lastName, username, phone, email, password, role, addressID, fax)
	mockUserRepo.AssertNumberOfCalls(t, "Save", 1)
}

func TestShouldCreateNewInvalidUser(t *testing.T) {
	userID := "afshsjgj14151jou"
	firstName := ""
	lastName := ""
	phone := ""
	email := "swastiksahoo22"
	username := ""
	password := "Pas"
	role := domain.Admin
	addressID := ""
	fax := ""


	user := domain.NewUser(userID, firstName, lastName, username, phone, email, password, role, addressID, fax)
	mockUserRepo.On("Save", mock.Anything).Return(*user, nil)
	userService.CreateUserInDynamodb(firstName, lastName, username, phone, email, password, role, addressID, fax)
	// mockUserRepo.AssertNumberOfCalls(t, "Save", 1)
}

func TestShouldGetUserByUserId(t *testing.T) {
	userID := "afshsjgj14151jou"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password, _ := domain.HashPassword("Pass!23")
	role := domain.Admin
	addressID := "addressid"
	fax := "18-1918-10101"


	user := domain.NewUser(userID, firstName, lastName, username, phone, email, password, role, addressID, fax)
	mockUserRepo.On("FindByID", userID).Return(user, nil)
	var _, _ = userService.GetUserById(userID)
	mockUserRepo.AssertNumberOfCalls(t, "FindByID", 1)
}

func TestShouldGetAllUsers(t *testing.T) {
	userID := "afshsjgj14151jou"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password, _ := domain.HashPassword("Pass!23")
	role := domain.Admin
	addressID := "addressid"
	fax := "18-1918-10101"


	user := domain.NewUser(userID, firstName, lastName, username, phone, email, password, role, addressID, fax)
	userArr := []domain.User{*user}
	mockUserRepo.On("FindAll").Return(userArr, nil)
	var _, _ = userService.GetAllUsers()
	mockUserRepo.AssertNumberOfCalls(t, "FindAll", 1)
}


func TestShouldUpdateUserByUserId(t *testing.T) {
	userID := "1"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password, _ := domain.HashPassword("Pass!23")
	role := domain.Admin
	addressID := "addressid"
	fax := "18-1918-10101"

	mockUserRepo.On("UpdateById", mock.Anything).Return(true, nil)

	var _, err = userService.UpdateUserById(userID, firstName, lastName, username, phone, email, password, role, addressID, fax)
	assert.Nil(t, err)
}


func TestShouldDeleteUserByUserId(t *testing.T) {
	userId := "1"
	mockUserRepo.On("DeleteByID", userId).Return(true, nil)

	var _, err = userService.DeleteUserById(userId)
	assert.Nil(t, err)
}

