package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/domain"
)

type CartHandler struct {
	CartService domain.CartService
}

type ProductRecordDTO struct {
	Product  string `json:"name"`
	Cost     int16  `json:"cost"`
	Quantity int    `json:"quantity"`
}

type CartRecordDTO struct {
	UserID   string             `json:"user_id"`
	Products []ProductRecordDTO `json:"products"`
}

// Create Cart
// @Summary      Create Cart
// @Description  This Handle allows admin to create new Cart
// @Tags         Cart
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /cart    [post]
func (ch CartHandler) HandleCart() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var cartDto CartRecordDTO
		if err := ctx.BindJSON(&cartDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		product_quantity := convertProductsDTOtoMaps(cartDto.Products)

		res, err := ch.CartService.CreateCart(
			cartDto.UserID,
			product_quantity,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Cart Record Added", "cart id": res})
	}
}

// Get Cart by ID
// @Summary      Get Cart by id
// @Description  This Handle returns Cart given cart id
// @Tags         Cart
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /cart/:id    [get]
func (ch CartHandler) HandleGetCartRecordByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
		res, err := ch.CartService.GetCartById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		cartdto := convertCartModeltoCartDTO(*res)
		ctx.JSON(http.StatusAccepted, gin.H{"record": cartdto})
	}
}

// Get All Cart records
// @Summary      Get All Cart records
// @Description  This Handle returns all of the carts
// @Tags         Cart
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /carts    [get]
func (ch CartHandler) HandleGetAllRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		record, err := ch.CartService.GetAllCarts()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		cartdto := make([]CartRecordDTO, 0)

		for _, rec := range record {
			respdto := convertCartModeltoCartDTO(rec)
			cartdto = append(cartdto, respdto)
		}

		ctx.JSON(http.StatusAccepted, gin.H{"record": cartdto})
	}
}

// Delete cart
// @Summary      Delete cart
// @Description  This Handle deletes cart given cart id
// @Tags         Cart
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /cart/:id   [delete]
func (ch CartHandler) HandleDeleteCartById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
		_, err := ch.CartService.DeleteCartById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted Succesfully"})
	}
}

func convertProductsDTOtoMaps(products []ProductRecordDTO) map[string]int {
	var product_quantity map[string]int = make(map[string]int)
	for _, product := range products {
		product_quantity[product.Product] = product.Quantity
	}
	return product_quantity
}

func convertCartModeltoCartDTO(cart domain.Cart) CartRecordDTO {
	var products []ProductRecordDTO
	for k, v := range cart.ProductsQuantity {
		products = append(products, ProductRecordDTO{Product: k, Quantity: v})
	}
	return CartRecordDTO{
		UserID:   cart.UserID,
		Products: products,
	}
}
