package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetRoleString(t *testing.T) {
	role := Admin
	var expected string = "admin"
	var actual string = role.String()

	assert.Equal(t, expected, actual)
}

func TestShouldReturnEnumIndexForRole(t *testing.T) {
	role := Admin
	var expected int = 1
	var actual int = role.EnumIndex()

	assert.Equal(t, expected, actual)
}

func TestShouldGetAdminEnumByIndex(t *testing.T) {
	var expected Role = Admin
	actual, err := GetEnumByIndex(0)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestShouldGetCustomerEnumByIndex(t *testing.T) {
	var expected Role = Customer
	actual, err := GetEnumByIndex(1)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestShouldReturnErrOnGetEnumByIndexForInvalidIndex(t *testing.T) {
	var expected Role = -1
	actual, err := GetEnumByIndex(1000)

	assert.Error(t, err.Error())
	assert.Equal(t, expected, actual)
}

func TestShouldReturnNewUser(t *testing.T) {
	id := "absbs"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password := "secret"
	role := Admin
	addressID := "addressid"
	fax := "18-1918-10101"

	user := NewUser(id, firstName, lastName, username, phone, email, password, role, addressID, fax)
	assert.Equal(t, firstName, user.FirstName)
	assert.Equal(t, lastName, user.LastName)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, phone, user.Phone)
	assert.Equal(t, email, user.Email) 
	assert.Equal(t, password, user.Password)
	assert.Equal(t, addressID, user.AddressID)
	assert.Equal(t, fax, user.Fax)
}

func TestShouldUpdateEmail(t *testing.T) {
	id := "absbs"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password := "secret"
	role := Admin
	addressID := "addressid"
	fax := "18-1918-10101"

	newEmail := "msadriwala.1198@gmail.com"
	user := NewUser(id, firstName, lastName, username, phone, email, password, role, addressID, fax)

	user.Email = newEmail

	assert.Equal(t, newEmail, user.Email)
}

func TestShouldUpdatePhone(t *testing.T) {
	id := "absbs"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password := "secret"
	role := Admin
	addressID := "addressid"
	fax := "18-1918-10101"

	newPhone := "9999955555"
	user := NewUser(id, firstName, lastName, username, phone, email, password, role, addressID, fax)

	user.Phone = newPhone

	assert.Equal(t, newPhone, user.Phone)
}

func TestShouldUpdateUsername(t *testing.T) {
	id := "absbs"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password := "secret"
	role := Admin
	addressID := "addressid"
	fax := "18-1918-10101"

	newUsername := "newUsername"
	user := NewUser(id, firstName, lastName, username, phone, email, password, role, addressID, fax)

	user.Username = newUsername

	assert.Equal(t, newUsername, user.Username)
}

func TestShouldUpdatePassword(t *testing.T) {
	id := "absbs"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password := "secret"
	role := Admin
	addressID := "addressid"
	fax := "18-1918-10101"

	newPassword := "newPass"
	user := NewUser(id, firstName, lastName, username, phone, email, password, role, addressID, fax)

	user.Password = newPassword

	assert.Equal(t, newPassword, user.Password)
}

func TestShouldUpdateFirstName(t *testing.T) {
	id := "absbs"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password := "secret"
	role := Admin
	addressID := "addressid"
	fax := "18-1918-10101"

	newFirstName := "SwastikNew"
	user := NewUser(id, firstName, lastName, username, phone, email, password, role, addressID, fax)

	user.FirstName = newFirstName

	assert.Equal(t, newFirstName, user.FirstName)
}

func TestShouldUpdateLastName(t *testing.T) {
	id := "absbs"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password := "secret"
	role := Admin
	addressID := "addressid"
	fax := "18-1918-10101"

	newLastName := "newLastName"
	user := NewUser(id, firstName, lastName, username, phone, email, password, role, addressID, fax)

	user.LastName = newLastName

	assert.Equal(t, newLastName, user.LastName)
}

func TestShouldUpdateRole(t *testing.T) {
	id := "absbs"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "1234567890"
	email := "swastiksahoo22@gmail.com"
	username := "swastik153"
	password := "secret"
	role := Admin
	addressID := "addressid"
	fax := "18-1918-10101"

	newRole := Customer
	user := NewUser(id, firstName, lastName, username, phone, email, password, role, addressID, fax)

	user.Role = newRole

	assert.Equal(t, newRole, user.Role)
}

func TestShouldMarshallJson(t *testing.T) {
	id := "absbs"
	firstName := "Swastik"
	lastName := "Sahoo"
	phone := "9900887766"
	email := "swastik123@gmail.com"
	username := "swastik123"
	password := "Pass!23"
	role := Admin
	addressID := "addressid"
	fax := "18-1918-10101"

	user := NewUser(id, firstName, lastName, username, phone, email, password, role, addressID, fax)

	expectedJson := "{\"address_id\":\"addressid\",\"email\":\"swastik123@gmail.com\",\"fax\":\"18-1918-10101\",\"firstName\":\"Swastik\",\"lastName\":\"Sahoo\",\"password\":\"Pass!23\",\"phone\":\"9900887766\",\"role\":1,\"user_id\":\"absbs\",\"username\":\"swastik123\"}"
	actualJson, _ := user.MarshalJSON()

	assert.Equal(t, expectedJson, string(actualJson))

}
