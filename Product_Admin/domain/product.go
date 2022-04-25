package domain

type ProductSEOURL struct {
	Keyword    string
	LanguageID string
	StoreID    string
}

type ProductDescription struct {
	LanguageID      string
	Name            string
	Description     string
	MetaTitle       string
	MetaDescription string
	MetaKeyword     string
	Tag             string
}

type ProductCategory struct {
	CategoryID string
}

type Product struct {
	Id                  string
	Model               string
	Quantity            int64
	Price               float64
	ManufacturerID      string
	SKU                 string
	ProductSEOURLs      []ProductSEOURL
	Points              int64
	Reward              int64
	ImageURL            string
	IsShippable         bool
	Weight              float64
	Length              float64
	Width               float64
	Height              float64
	MinimumQuantity     int64
	RelatedProducts     []string
	ProductDescriptions []ProductDescription
	ProductCategories   []ProductCategory
}

type ProductAdminDynamoRepository interface {
	Insert(Product) (bool, error)
	Find() ([]Product, error)
	FindByID(string) (Product, error)
	UpdateItem(string, int64) (bool, error)
	DeleteByID(string) (bool, error)
}

func NewProductObject(model string, quantity int64, price float64, manufacturerID string, sku string,
	productSEOURLs []ProductSEOURL, points int64, reward int64, imageURL string, isShippable bool,
	weight float64, length float64, width float64, height float64, minimumQuantity int64,
	relatedProducts []string, productDescription []ProductDescription, productCategories []ProductCategory) *Product {
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