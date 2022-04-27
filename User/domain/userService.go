package domain

import (
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"unicode"
	"fmt"
)

type UserService interface {
	CreateUserInDynamodb(string, string, string, string, string, string, Role, string, string) (User, error)
	GetUserById(string) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUserById(user_id, firstName, lastName, username, phone, email, password string, role Role, addressId, fax string) (bool, error)
	DeleteUserById(string) (bool, error)
}

type service struct {
	userDynamodbRepository UserDynamoDBRepository
}

func NewUserService(userDynamodbRepository UserDynamoDBRepository) UserService {
	return &service{
		userDynamodbRepository: userDynamodbRepository,
	}
}


func (s service) CreateUserInDynamodb(firstName, lastName, username, phone, email, password string, role Role, addressId, fax string) (User, error) {
	id := _generateUniqueId()

	isValid, message := validateFeilds(firstName, lastName, username, phone, email, password, fax)

	if isValid != true{
		return User{}, fmt.Errorf(message)
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return User{}, err
	}
	user := NewUser(id, firstName, lastName, username, phone, email, hashedPassword, role, addressId, fax)
	persistedUser, err1 := s.userDynamodbRepository.Save(*user)
	if err1 != nil {
		return User{}, err1
	}
	return persistedUser, nil
}


func (s service) GetUserById(userId string) (*User, error) {
	res, err := s.userDynamodbRepository.FindByID(userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (s service) GetAllUsers() ([]User, error) {
	res, err := s.userDynamodbRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (s service) UpdateUserById(user_id, firstName, lastName, username, phone, email, password string, role Role, addressId, fax string) (bool, error) {
	user := NewUser(user_id, firstName, lastName, username, phone, email, password, role, addressId, fax)
	_, err := s.userDynamodbRepository.UpdateById(*user)
	if err != nil {
		return false, err
	}
	return true, err
}


func (s service) DeleteUserById(userId string) (bool, error) {
	_, err := s.userDynamodbRepository.DeleteByID(userId)
	if err != nil {
		return false, err
	}
	return true, err
}


func _generateUniqueId() string {
	return primitive.NewObjectID().Hex()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return string(bytes), err
	}
	return string(bytes), nil
}

func validateFeilds(firstName, lastName, username, phone, email, password, fax string) (bool, string){
	isValid := true
	message := ""

	//email
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if emailRegex.MatchString(email) != true {
		isValid = false
		message += "Invalid Email"
		message += "\n"
	}

	//password
	number := false
	upper := false
	special := false
	letters := 0
    for _, c := range password {
        switch {
        case unicode.IsNumber(c):
            number = true
			letters++
        case unicode.IsUpper(c):
            upper = true
            letters++
        case unicode.IsPunct(c) || unicode.IsSymbol(c):
            special = true
        case unicode.IsLetter(c) || c == ' ':
            letters++
        default:
            //return false, false, false, false
        }
    }
    sevenOrMore := letters >= 7

	if (number!=true) || (upper!=true) || (special!=true) || (sevenOrMore!=true){
		isValid = false
		message += "Invalid Password, should have the following: number, uppercase letter, special character and seven or more characters"
		message += "\n"
	}

	//firstname
	if (firstName == ""){
		isValid = false
		message += "Cannot have first name empty"
		message += "\n"
	}

	//lastname
	if (lastName == ""){
		isValid = false
		message += "Cannot have last name empty"
		message += "\n"
	}

	//username
	if (username == ""){
		isValid = false
		message += "Cannot have username empty"
		message += "\n"
	}

	//phone
	if (phone == ""){
		isValid = false
		message += "Cannot have phone number empty"
		message += "\n"
	}

	//fax
	if (fax == ""){
		isValid = false
		message += "Cannot have fax empty"
		message += "\n"
	}

	return isValid, message
}