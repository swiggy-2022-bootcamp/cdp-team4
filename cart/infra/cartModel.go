package infra
import (
	"time"
)

type CartModel struct {
	Id               string         `json:"id"`
	UserID           string         `json:"user_id"`
	ProductsQuantity map[string]int `json:"products_quantity"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}
