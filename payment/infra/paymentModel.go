package infra

import (
	"time"
)

type PayModel struct {
	Id          string    `json:"id"`
	Amount      int16     `json:"amount"`
	Currency    string    `json:"currency"`
	Status      string    `json:"status"`
	OrderID     string    `json:"order_id"`
	Method      string    `json:"method"`
	Description string    `json:"description"`
	Bank        string    `json:"bank"`
	Wallet      string    `json:"wallet"`
	VPA         string    `json:"vpa"`
	UserID      string    `json:"user_id"`
	Notes       []string  `json:"notes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaymentMethodModel struct {
	Id      string   `json:"id"`
	Methods []string `json:"methods"`
	Agree   string   `json:"agree"`
	Comment string   `json:"comment"`
}
