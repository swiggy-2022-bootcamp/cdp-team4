package domain
type Cart struct {
	Id               string         `json:"id"`
	UserID           string         `json:"user_id"`
	ProductsQuantity map[string]int `json:"products_quantity"`
}

func NewCart(userId string,products_quantity map[string]int) *Cart {
	return &Cart{
		UserID:           userId,
		ProductsQuantity: products_quantity,
	}
}