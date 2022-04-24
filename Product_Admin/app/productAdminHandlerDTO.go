package app

// data model of Productt record  used to parse the body of the http request
// and pass on the data to service layer that is going to redirect to
// infra layer and save it to database

type ProductSEOURLDTO struct {
	Keyword    string `json:"keyword"`
	LanguageID string `json:"language_id"`
	StoreID    string `json:"store_id"`
}

type ProductDescriptionDTO struct {
	LanguageID      string `json:"language_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeyword     string `json:"meta_keyword"`
	Tag             string `json:"tag"`
}

type ProductDTO struct {
	Model               string                  `json:"model"`
	Quantity            int64                   `json:"quantity"`
	Price               float64                 `json:"price"`
	ManufacturerID      string                  `json:"manufacturer_id"`
	SKU                 string                  `json:"sku"`
	ProductSEOURLs      []ProductSEOURLDTO      `json:"product_seo_url"`
	Points              int64                   `json:"points"`
	Reward              int64                   `json:"reward"`
	ImageURL            string                  `json:"image_url"`
	IsShippable         bool                    `json:"is_shippable"`
	Weight              float64                 `json:"weight"`
	Length              float64                 `json:"length"`
	Width               float64                 `json:"width"`
	Height              float64                 `json:"height"`
	MinimumQuantity     int64                   `json:"minimum_quantity"`
	RelatedProducts     []string                `json:"related_products"`
	ProductDescriptions []ProductDescriptionDTO `json:"product_description"`
	ProductCategories   []string                `json:"product_categories"`
}
