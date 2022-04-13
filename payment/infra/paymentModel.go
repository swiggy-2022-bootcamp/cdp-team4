package infra

import (
	"time"
)

type PayModel struct {
	Id          string
	Amount      int16
	Currency    string
	Status      string
	OrderID     string
	Method      string
	Description string
	Bank        string
	Wallet      string
	VPA         string
	UserID      string
	Notes       []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
