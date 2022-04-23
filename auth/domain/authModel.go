package domain

import (
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/utils/errs"
	"time"
)

type AuthModel struct {
	UserId    string    `json:"user_id"`
	Role      Role      `json:"role"`
	AuthToken string    `json:"auth_token"`
	IsExpired bool      `json:"is_expired"`
	ExpiresOn time.Time `json:"expires_on"`
}

type AuthRepository interface {
	FindByUserIdAndAuthToken(string, string) (*AuthModel, *errs.AppError)
	Save(AuthModel) *errs.AppError
}
