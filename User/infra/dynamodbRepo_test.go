package infra_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/infra"
)

var testUserService = infra.NewDynamoRepository()
var insertedid string
var inserteduserid string

func TestShouldCreateNewUserinDynamoDB(t *testing.T) {
	userID := "afshsjgj14151joi"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password, _ := domain.HashPassword("Pass!23")
	role := domain.Admin

	user := domain.NewUser(userID, firstName, lastName, username, phone, email, password, role)
	res, err := testUserService.Save(*user)
	t.Logf("Inserted user is %s\n", res)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}


func TestShouldGetUserByUserIdDynamoDB(t *testing.T) {
	userID := "afshsjgj14151joi"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password, _ := domain.HashPassword("Pass!23")
	role := domain.Admin

	t.Logf("Inserted User Id is %s Reading\n", userID)
	res, err := testUserService.FindByID(userID)
	t.Logf("Read %v", res)

	user := domain.NewUser(userID, firstName, lastName, username, phone, email, password, role)

	assert.NotNil(t, res)
	assert.Nil(t, err)

	assert.Equal(t, res.FirstName, user.FirstName)
	assert.Equal(t, res.LastName, user.LastName)
	assert.Equal(t, res.Username, user.Username)
	assert.Equal(t, res.Phone, user.Phone)
	assert.Equal(t, res.Email, user.Email)
	assert.Equal(t, res.Role, user.Role)
}

func TestShouldUpdateUserByIdDynamoDB(t *testing.T) {
	userID := "afshsjgj14151joi"
	firstName := "SwastikU"
	lastName := "SahooU"
	phone := "1234567890U"
	email := "swastiksahoo22@gmail.comU"
	username := "swastik153U"
	password, _ := domain.HashPassword("Pass!23U")
	role := domain.Admin

	user := domain.NewUser(userID, firstName, lastName, username, phone, email, password, role)

	t.Logf("Inserted User Id is %s Reading\n", userID)
	res, err := testUserService.UpdateById(*user)
	t.Logf("Read %v", res)

	assert.NotNil(t, res)
	assert.Nil(t, err)

	assert.Equal(t, res,true)
}

func TestShouldUpdateUserByIdDynamoDB2(t *testing.T) {
	userID := "afshsjgj14151joi"
	firstName := ""
	lastName := ""
	phone := ""
	email := ""
	username := ""
	password := ""
	role := domain.Admin

	user := domain.NewUser(userID, firstName, lastName, username, phone, email, password, role)

	t.Logf("Inserted User Id is %s Reading\n", userID)
	res, err := testUserService.UpdateById(*user)
	t.Logf("Read %v", res)

	assert.NotNil(t, res)
	assert.Nil(t, err)

	assert.Equal(t, res,true)
}

func TestShouldGetAllUsersDynamoDB(t *testing.T) {
	res, err := testUserService.FindAll()
	assert.NotNil(t, res)
	assert.Nil(t, err)
}


func TestShouldDeleteUserByIdDynamoDB(t *testing.T) {
	userID := "afshsjgj14151joi"
	res, err := testUserService.DeleteByID(userID)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}