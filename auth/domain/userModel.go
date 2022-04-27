package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/auth/utils/errs"

type Role int

const (
	Admin Role = iota + 1
	Customer
)

func (r Role) String() string {
	return [...]string{"admin", "customer"}[r-1]
}

func (r Role) EnumIndex() int {
	return int(r)
}

func GetEnumByIndex(idx int) (Role, *errs.AppError) {
	switch idx {
	case 1:
		return Admin, nil
	case 2:
		return Customer, nil
	default:
		return -1, errs.NewUnexpectedError("invalid enum index")
	}
}

type UserModel struct {
	UserId   string `json:"user_id"`
	Role     Role   `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	FindByUsername(string) (*UserModel, *errs.AppError)
}
