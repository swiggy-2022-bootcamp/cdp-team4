package domain

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

type ProductCategory struct {
	CategoryID string
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
	// ProductCategories   []ProductCategory
}

type ProductAdminDynamoRepository interface {
	Insert(Product) (bool, error)
	Find() ([]Product, error)
	FindByID(string) (Product, error)
	UpdateItem(string, int64) (bool, error)
	DeleteByID(string) (bool, error)
	GetProductAvailability(string,int64)(bool,error)
	FindByCategoryID(string)([]Product, error)
	FindByManufacturerID(string)([]Product, error)
	FindByKeyword(string)([]Product, error)
}

func NewProductObject(model string, quantity int64, price float64, manufacturerID string, sku string,
	productSEOURLs []ProductSEOURL, points int64, reward int64, imageURL string, isShippable bool,
	weight float64, length float64, width float64, height float64, minimumQuantity int64,
	relatedProducts []string, productDescription []ProductDescription, productCategories []string) *Product {
	return &Product{
		Id:                  _generateUniqueId(),
		Model:               model,
		Quantity:            quantity,
		Price:               price,
		ManufacturerID:      manufacturerID,
		SKU:                 sku,
		ProductSEOURLs:      productSEOURLs,
		Points:              points,
		Reward:              reward,
		ImageURL:            imageURL,
		IsShippable:         isShippable,
		Weight:              weight,
		Length:              length,
		Width:               width,
		Height:              height,
		MinimumQuantity:     minimumQuantity,
		RelatedProducts:     relatedProducts,
		ProductDescriptions: productDescription,
		ProductCategories:   productCategories,
	}
}
