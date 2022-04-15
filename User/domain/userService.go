package domain

import (
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "fmt"
)

type UserService interface {
	CreateUserInDynamodb(string, string, string, string, string, string, Role) (User, error)
	GetUserById(string) (*User, error)
	GetAllUsers() ([]User, error)
	DeleteUserById(string) (bool, *error)
}

type service struct {
	userDynamodbRepository UserDynamoDBRepository
}

func NewUserService(userDynamodbRepository UserDynamoDBRepository) UserService {
	return &service{
		userDynamodbRepository: userDynamodbRepository,
	}
}


func (s service) CreateUserInDynamodb(firstName, lastName, username, phone, email, password string, role Role) (User, error) {
	id := _generateUniqueId()
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return User{}, err
	}
	user := NewUser(id, firstName, lastName, username, phone, email, hashedPassword, role)
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


func (s service) DeleteUserById(userId string) (bool, *error) {
	_, err := s.userDynamodbRepository.DeleteByID(userId)
	if err != nil {
		return false, &err
	}
	return true, nil
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