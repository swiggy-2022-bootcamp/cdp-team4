package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/domain"
)

// ProductAdminHandler wraps up the ProductAdminServices along with
// all the handler methods of respective routes
type ProductAdminHandler struct {
	ProductAdminService domain.ProductAdminService
}

// Add product
// @Summary      Add Product
// @Description  This Handle allows admin to create a new product
// @Tags         Product Admin
// @Produce      json
// @Accept       json
// @Param        product body ProductDTO  true "product request structure"
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
// @Accept 		 json
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
// @Router       /products/{id}    [get]
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
// @Accept       json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/{id}    [put]
func (ph ProductAdminHandler) HandleUpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productDTO ProductDTO
		if err := ctx.BindJSON(&productDTO); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		var productSEOURL []domain.ProductSEOURL
		for _, item := range productDTO.ProductSEOURLs {
			productSEOURL = append(productSEOURL, domain.ProductSEOURL{Keyword: item.Keyword, LanguageID: item.LanguageID, StoreID: item.StoreID})
		}
		var productDescription []domain.ProductDescription
		for _, item := range productDTO.ProductDescriptions {
			productDescription = append(productDescription, domain.ProductDescription{LanguageID: item.LanguageID, Name: item.Name,
				Description: item.Description, MetaTitle: item.MetaTitle, MetaDescription: item.MetaDescription, MetaKeyword: item.MetaKeyword,
				Tag: item.Tag})
		}
		var productCategories []string
		for _, item := range productDTO.ProductCategories {
			productCategories = append(productCategories, item)
		}
		product := domain.NewProductObject(productDTO.Model, productDTO.Quantity, productDTO.Price, productDTO.ManufacturerID, productDTO.SKU,
			productSEOURL, productDTO.Points, productDTO.Reward, productDTO.ImageURL, productDTO.IsShippable, productDTO.Weight, productDTO.Length,
			productDTO.Width, productDTO.Height, productDTO.MinimumQuantity, productDTO.RelatedProducts, productDescription, productCategories)

		ok, err := ph.ProductAdminService.UpdateProduct(*product)
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
// @Accept       json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/{id}    [delete]
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

// Search product
// @Summary      Search product
// @Description  This Handles search of a product given category id
// @Tags         Product Admin
// @Produce      json
// @Accept       json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /products/search/{id}    [get]
func (ph ProductAdminHandler) HandleSearchByCategoryID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryId := ctx.Param("categoryid")
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

func NewProductAdminHandler(productAdminService domain.ProductAdminService) ProductAdminHandler {
	return ProductAdminHandler{
		ProductAdminService: productAdminService,
	}
}
