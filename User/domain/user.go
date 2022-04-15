package domain

import (
	"encoding/json"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/utils/errs"
)

type Role int

const (
	Admin    Role = iota // EnumIndex = 0
	Customer             // EnumIndex = 1
)

func (r Role) String() string {
	return [...]string{"admin", "seller", "customer"}[r]
}

func (r Role) EnumIndex() int {
	return int(r)
}

func GetEnumByIndex(idx int) (Role, *errs.AppError) {
	switch idx {
	case 0:
		return Admin, nil
	case 1:
		return Customer, nil
	default:
		return -1, errs.NewUnexpectedError("invalid enum index")
	}
}

type User struct {
	UserID          string             `json:"user_id"`
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	Username        string             `json:"username"`
	Password        string             `json:"password"`
	Phone           string             `json:"phone"`
	Email           string             `json:"email"`
	Role            Role               `json:"role"`
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"user_id":          u.UserID,
		"firstName":        u.FirstName,
		"lastName":         u.LastName,
		"email":            u.Email,
		"password":         u.Password,
		"username":         u.Username,
		"phone":            u.Phone,
		"role":             u.Role,
	})
}

func NewUser(userId, firstName, lastName, username, phone, email, password string, role Role) *User {
	return &User{
		UserID:			 userId,
		FirstName:       firstName,
		LastName:        lastName,
		Username:        username,
		Phone:           phone,
		Email:           email,
		Password:        password,
		Role:            role,
	}
}

type UserDynamoDBRepository interface {
	Save(User) (User, error)
	FindByID(string) (*User, error)
	DeleteByID(string) (bool, error)
	FindAll() ([]User, error)
}

