package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUserInDynamodb(string, string, string, string, string, string, Role) (User, error)
	GetUserById(string) (*User, error)
}

type service struct {
	userDynamodbRepository UserDynamoDBRepository
}


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return string(bytes), err
	}
	return string(bytes), nil
}


func (s service) CreateUserInDynamodb(firstName, lastName, username, phone, email, password string, role Role) (User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return User{}, err
	}
	user := NewUser(firstName, lastName, username, phone, email, hashedPassword, role)
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

func NewUserService(userDynamodbRepository UserDynamoDBRepository) UserService {
	return &service{
		userDynamodbRepository: userDynamodbRepository,
	}
}
