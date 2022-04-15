package infra

import (
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"
)

type UserModel struct {
	UserID          string             `json:"user_id"`
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	Username        string             `json:"username"`
	Password        string             `json:"password"`
	Phone           string             `json:"phone"`
	Email           string             `json:"email"`
	Role            domain.Role        `json:"role"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}
