package infra

import (
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain"
	"time"
)

type AuthModel struct {
	UserId    string      `json:"user_id"`
	AuthToken string      `json:"auth_token"`
	Role      domain.Role `json:"role"`
	IsExpired bool        `json:"is_expired"`
	ExpiresOn time.Time   `json:"expires_on"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
