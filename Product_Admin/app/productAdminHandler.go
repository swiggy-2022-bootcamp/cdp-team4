package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/domain"
)

type ProductAdminHandler struct {
	ProductAdminService domain.ProductAdminService
}

// Add product
// @Summary      Add Product
// @Description  This Handle allows admin to create a new product
// @Tags         Product Admin
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/    [post]
func (ph ProductAdminHandler) HandleAddProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product ProductDTO
		if err := ctx.BindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		var productSEOURL []domain.ProductSEOURL
		for _, item := range product.ProductSEOURLs {
			productSEOURL = append(productSEOURL, domain.ProductSEOURL{Keyword: item.Keyword, LanguageID: item.LanguageID, StoreID: item.StoreID})
		}
		var productDescription []domain.ProductDescription
		for _, item := range product.ProductDescriptions {
			productDescription = append(productDescription, domain.ProductDescription{LanguageID: item.LanguageID, Name: item.Name,
				Description: item.Description, MetaTitle: item.MetaTitle, MetaDescription: item.MetaDescription, MetaKeyword: item.MetaKeyword,
				Tag: item.Tag})
		}
		var productCategories []string
		for _, item := range product.ProductCategories {
			productCategories = append(productCategories, item)
		}
		result, err := ph.ProductAdminService.CreateDynamoProductAdminRecord(product.Model, product.Quantity, product.Price,
			product.ManufacturerID, product.SKU, productSEOURL, product.Points, product.Reward, product.ImageURL,
			product.IsShippable, product.Weight, product.Length, product.Width, product.Height, product.MinimumQuantity,
			product.RelatedProducts, productDescription, productCategories)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to add product")
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "New product added", "product id": result})
		log.WithFields(logrus.Fields{"product id": result, "status": http.StatusOK}).
			Info("New product added")
	}
}

// Get all products
// @Summary      Get all Products
// @Description  This Handle allows admin to fetch all the products in the datastore
// @Tags         Product Admin
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/    [get]
func (ph ProductAdminHandler) HandleGetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		records, err := ph.ProductAdminService.GetProduct()
		fmt.Print("handlegetallproducts ", records, err)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to get products")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"records": records})
		log.WithFields(logrus.Fields{"products": records, "status": http.StatusOK}).
			Info("All products fetched")
	}
}

// Get product by ID
// @Summary      Get Product details by Id
// @Description  This Handle allows admin to get a product, given Id
// @Tags         Product Admin
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/    [post]
func (ph ProductAdminHandler) HandleGetProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		product, err := ph.ProductAdminService.GetProductById(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to get product")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"product": product})
		log.WithFields(logrus.Fields{"product id": product, "status": http.StatusOK}).
			Info("Product fetched by Id")
	}
}

// Update product details
// @Summary      Update product details
// @Description  This Handles Updation of product details given product id
// @Tags         Product Admin
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/:id    [put]
func (ph ProductAdminHandler) HandleUpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Todo - update any field of product
		//currently it supports only updation of product quantity
		var requestDTO RequestDTO
		if err := ctx.BindJSON(&requestDTO); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ok, err := ph.ProductAdminService.UpdateProduct(requestDTO.Id, requestDTO.QuantityToBeReduced)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to update product details")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "product record updated"})
		log.WithFields(logrus.Fields{"status": http.StatusOK}).Info("Product updated")
	}
}

// Delete product
// @Summary      Delete product
// @Description  This Handles deletion of a product given product id
// @Tags         Product Admin
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/:id    [delete]
func (ph ProductAdminHandler) HandleDeleteProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		_, err := ph.ProductAdminService.DeleteProductById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to delete product")
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted Succesfully"})
		log.WithFields(logrus.Fields{"status": http.StatusOK}).Info("Product deleted successfully")
	}
}

// Search products
// @Summary      Search product by category id
// @Description  This Handles searching a product given categoryid
// @Tags         Product Admin
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/:id    [delete]
func (ph ProductAdminHandler) HandleSearchByCategoryID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryId := ctx.Param("id")
		res, err := ph.ProductAdminService.GetProductByCategoryId(categoryId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to find product")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": res, "status": http.StatusOK})
		log.WithFields(logrus.Fields{"status": http.StatusOK, "products": res}).Info("Product deleted successfully")
	}
}

func (ph ProductAdminHandler) HandleSearchByManufacturerID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		manufacturerId := ctx.Param("id")
		res, err := ph.ProductAdminService.GetProductByManufacturerId(manufacturerId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to find product")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": res})
		log.WithFields(logrus.Fields{"status": http.StatusOK, "products": res}).Info("Product deleted successfully")
	}
}

func (ph ProductAdminHandler) HandleSearchByKeyword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyword := ctx.Param("keyword")
		res, err := ph.ProductAdminService.GetProductByKeyword(keyword)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to find product")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": res})
		log.WithFields(logrus.Fields{"status": http.StatusOK, "products": res}).Info("Product deleted successfully")
	}
}

/**********DTO***********/
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

type RequestDTO struct {
	Id                  string `json:"id"`
	QuantityToBeReduced int64  `json:"quantity_to_be_reduced"`
}
