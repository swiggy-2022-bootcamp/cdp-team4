package app

import "github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain"

type ValidationDTO struct {
	UserId string      `json:"user_id"`
	Role   domain.Role `json:"role"`
}
