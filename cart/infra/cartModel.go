package infra

import (
	"time"
)

type Item struct {
	Name     string `json:"name"`
	Cost     int `json:"cost"`
	Quantity int `json:"quantity"`
}
type CartModel struct {
	Id        string          `json:"id"`
	UserID    string          `json:"user_id"`
	Items     map[string]Item `json:"items"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
