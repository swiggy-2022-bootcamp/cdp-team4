package domain

//Todo - quantity can be removed from it
type ProductSEOURL struct {
	Keyword    string `json:"keyword"`
	LanguageID string `json:"language_id"`
	StoreID    string `json:"store_id"`
}

type ProductDescription struct {
	LanguageID      string `json:"language_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeyword     string `json:"meta_keyword"`
	Tag             string `json:"tag"`
}

type Product struct {
	Id                  string               `json:"id"`
	Model               string               `json:"model"`
	Quantity            int64                `json:"quantity"`
	Price               float64              `json:"price"`
	ManufacturerID      string               `json:"manufacturer_id"`
	SKU                 string               `json:"sku"`
	ProductSEOURLs      []ProductSEOURL      `json:"product_seo_url"`
	Points              int64                `json:"points"`
	Reward              int64                `json:"reward"`
	ImageURL            string               `json:"image_url"`
	IsShippable         bool                 `json:"is_shippable"`
	Weight              float64              `json:"weight"`
	Length              float64              `json:"length"`
	Width               float64              `json:"width"`
	Height              float64              `json:"height"`
	MinimumQuantity     int64                `json:"minimum_quantity"`
	RelatedProducts     []string             `json:"related_products"`
	ProductDescriptions []ProductDescription `json:"product_description"`
	ProductCategories   []string             `json:"product_categories"`
}

//go:generate mockgen -source=ProductFrontStoreDynamoRepository
type ProductFrontStoreDynamoRepository interface {
	Find() ([]Product, error)
	FindByProductID(string) (Product, error)
	FindByCategoryID(string) ([]Product, error)
}
