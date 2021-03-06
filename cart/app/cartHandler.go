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
	Id       string `json:"id"`
	Name     string `json:"name"`
	Cost     int    `json:"cost"`
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
// @Schemes
// @Accept json
// @Produce      json
// @Param   req  body CartRecordDTO true "Cart details"
// @Success      200  string    Cart record Added
// @Failure      400  string   	Bad request
// @Router       /cart    [post]
func (ch CartHandler) HandleCreateCart() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var cartDto CartRecordDTO
		if err := ctx.BindJSON(&cartDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		domainItemMap := convertProductsDTOtoMaps(cartDto.Products)

		res, err := ch.CartService.CreateCart(
			cartDto.UserID,
			domainItemMap,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Cart Record Added", "cart id": res})
	}
}

// Update Cart
// @Summary      Update Cart items
// @Description  This Handle allows to add new items/update items to cart
// @Tags         Cart
// @Schemes
// @Accept json
// @Produce      json
// @Param   req  query int true "User id"
// @Param   req  body CartRecordDTO true "Cart details"
// @Success      200 string  Cart Record Update
// @Failure      400  string   	Bad request
// @Router       /cart/:userId     [put]
func (ch CartHandler) HandleUpdateCartItemByUserId() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("userId")
		fmt.Println(id)
		var cartDto CartRecordDTO
		if err := ctx.BindJSON(&cartDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		domainItemMap := convertProductsDTOtoMaps(cartDto.Products)

		res, err := ch.CartService.UpdateCartItemsByUserId(
			cartDto.UserID,
			domainItemMap,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Cart Record Update", "Success Status": res})
	}
}

// Delete Cart items
// @Summary      Delete Cart items
// @Description  This Handle allows to Delete items to cart
// @Tags         Cart
// @Schemes
// @Accept json
// @Produce      json
// @Param   req  query int true "User id"
// @Param   req  body CartRecordDTO true "Cart details"
// @Success      200 string  Cart Record Update
// @Failure      400  string   	Bad request
// @Router       /cart/:userId    [delete]
func (ch CartHandler) HandleDeleteCartItemByUserId() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("userId")
		fmt.Println(id)
		type DeleteItemfromCartRequest struct{
			ProductList []string `json:"product_list"`
		}
		var reqBody DeleteItemfromCartRequest
		if err := ctx.BindJSON(&reqBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		fmt.Println(reqBody)
		res, err := ch.CartService.DeleteCartItemByUserId(
			id,
			reqBody.ProductList,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Cart Record Update","Success Status": res})
	}
}


// // Get Cart by ID
// // @Summary      Get Cart by id
// // @Description  This Handle returns Cart given cart id
// // @Tags         Cart
// // @Produce      json
// // @Success      200  {object}  map[string]interface{}
// // @Failure      400  {number} 	http.StatusBadRequest
// // @Router       /cart/:id    [get]
// func (ch CartHandler) HandleGetCartRecordByID() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		id := ctx.Param("id")
// 		fmt.Println(id)
// 		res, err := ch.CartService.GetCartById(id)
//
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
// 			return
// 		}
// 		cartdto := convertCartModeltoCartDTO(*res)
// 		ctx.JSON(http.StatusAccepted, gin.H{"record": cartdto})
// 	}
// }


// Get Cart by UserId
// @Summary      Get Cart by UserId
// @Description  This Handle returns Cart given cart UserId
// @Tags         Cart
// @Produce      json
// @Param   req  query int true "User id"
// @Success      200  {object}  domain.Cart
// @Failure      400  string   	Bad request
// @Router       /cart/:userId    [get]
func (ch CartHandler) HandleGetCartRecordByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("userId")
		fmt.Println(userId)
		res, err := ch.CartService.GetCartByUserId(userId)

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
// @Success      200  {object}  []domain.Cart
// @Failure      400  string    Record not found
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

// // Delete cart by cart ID
// // @Summary      Delete cart
// // @Description  This Handle deletes cart given cart id
// // @Tags         Cart
// // @Produce      json
// // @Success      200  {object}  map[string]interface{}
// // @Failure      400  string   	Bad request
// // @Router       /cart/:id   [delete]
// func (ch CartHandler) HandleDeleteCartById() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		id := ctx.Param("id")
// 		fmt.Println(id)
// 		_, err := ch.CartService.DeleteCartById(id)

// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 			return
// 		}

// 		ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted Succesfully"})
// 	}
// }



// Delete cart By User Id
// @Summary      Delete cart By User Id
// @Description  This Handle deletes cart given User ID
// @Tags         Cart
// @Produce      json
// @Param   req  query int true "User id"
// @Success      200  string Deleted Succesfully
// @Failure      400  string   	Bad request
// @Router       /cart/:userId   [delete]
func (ch CartHandler) HandleDeleteCartByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("userId")
		fmt.Println(id)
		_, err := ch.CartService.DeleteCartByUserId(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted Succesfully"})
	}
}

func convertProductsDTOtoMaps(products []ProductRecordDTO) map[string]domain.Item {
	var domainItemMap map[string]domain.Item = make(map[string]domain.Item)
	for _, product := range products {
		var singleItem domain.Item
		singleItem.Cost=product.Cost
		singleItem.Quantity=product.Quantity
		singleItem.Name=product.Name	
		domainItemMap[product.Id] = singleItem
	}
	return domainItemMap
}

func convertCartModeltoCartDTO(cart domain.Cart) CartRecordDTO {
	var products []ProductRecordDTO
	for k, v := range cart.Items {
		products = append(products, ProductRecordDTO{Id:k,Name:v.Name,Cost:v.Cost,Quantity:v.Quantity})
	}
	return CartRecordDTO{
		UserID:   cart.UserID,
		Products: products,
	}
}
