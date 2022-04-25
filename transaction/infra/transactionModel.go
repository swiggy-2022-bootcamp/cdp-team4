package infra

import (
	"time"
)

type TransactionModel struct {
	Id                string    `json:"id"`
	UserID            string    `json:"user_id"`
	TransactionPoints int       `json:"transaction_points"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
