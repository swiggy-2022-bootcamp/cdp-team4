package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
)

type OrderHandler struct {
	OrderService domain.OrderService
}

type ProductRecordDTO struct {
	Product  string `json:"name"`
	Cost     int16  `json:"cost"`
	Quantity int    `json:"quantity"`
}

type OrderRecordDTO struct {
	UserID    string             `json:"user_id"`
	Status    string             `json:"status"`
	Products  []ProductRecordDTO `json:"products"`
	TotalCost int16              `json:"total_cost"`
}

// Create Order
// @Summary      Create Order
// @Description  This Handle allows admin to create new Order
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /order    [post]
func (oh OrderHandler) handleOrder() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var orderDto OrderRecordDTO
		if err := ctx.BindJSON(&orderDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		product_quantity, product_cost := convertProductsDTOtoMaps(orderDto.Products)

		res, err := oh.OrderService.CreateOrder(
			orderDto.UserID,
			orderDto.Status,
			product_quantity,
			product_cost,
			int(orderDto.TotalCost),
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Order Record Added", "order id": res})
	}
}

// Get Order by ID
// @Summary      Get Order by id
// @Description  This Handle returns Order given order id
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /order/:id    [get]
func (oh OrderHandler) HandleGetOrderRecordByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
		res, err := oh.OrderService.GetOrderById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		orderdto := convertOrderModeltoOrderDTO(*res)
		ctx.JSON(http.StatusAccepted, gin.H{"record": orderdto})
	}
}

// Get Order by User ID
// @Summary      Get Order by user id
// @Description  This Handle returns Order given user id
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /order/user/:userid    [get]
func (oh OrderHandler) HandleGetOrderRecordsByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("user_id")
		fmt.Println(id)
		record, err := oh.OrderService.GetOrderByUserId(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}

		orderdto := make([]OrderRecordDTO, 0)

		for _, rec := range record {
			respdto := convertOrderModeltoOrderDTO(rec)
			orderdto = append(orderdto, respdto)
		}

		ctx.JSON(http.StatusAccepted, gin.H{"record": orderdto})
	}
}

// Get Order by Status
// @Summary      Get Order by status
// @Description  This Handle returns Order given status
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /order/status/:status    [get]
func (oh OrderHandler) HandleGetOrderRecordsByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Param("status")
		record, err := oh.OrderService.GetOrderByStatus(status)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}

		orderdto := make([]OrderRecordDTO, 0)

		for _, rec := range record {
			respdto := convertOrderModeltoOrderDTO(rec)
			orderdto = append(orderdto, respdto)
		}

		ctx.JSON(http.StatusAccepted, gin.H{"record": orderdto})
	}
}

// Get All Order records
// @Summary      Get All Order records
// @Description  This Handle returns all of the orders
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /orders    [get]
func (oh OrderHandler) HandleGetAllRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		record, err := oh.OrderService.GetAllOrders()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		orderdto := make([]OrderRecordDTO, 0)

		for _, rec := range record {
			respdto := convertOrderModeltoOrderDTO(rec)
			orderdto = append(orderdto, respdto)
		}

		ctx.JSON(http.StatusAccepted, gin.H{"record": orderdto})
	}
}

// Update order status
// @Summary      Update order status
// @Description  This Handle Update order status given order id
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /order/status    [put]
func (oh OrderHandler) handleUpdateOrderStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestDTO struct {
			Id     string `json:"id"`
			Status string `json:"status"`
		}
		if err := ctx.BindJSON(&requestDTO); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ok, err := oh.OrderService.UpdateOrderStatus(requestDTO.Id, requestDTO.Status)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "order record updated"})
	}
}

// Delete order
// @Summary      Delete order
// @Description  This Handle deletes order given order id
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /order/:id   [delete]
func (oh OrderHandler) HandleDeleteOrderById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
		_, err := oh.OrderService.DeleteOrderById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted Succesfully"})
	}
}

func convertProductsDTOtoMaps(products []ProductRecordDTO) (map[string]int, map[string]int) {
	var product_quantity map[string]int = make(map[string]int)
	var product_cost map[string]int = make(map[string]int)
	for _, product := range products {
		product_quantity[product.Product] = product.Quantity
		product_cost[product.Product] = int(product.Cost)
	}
	return product_quantity, product_cost
}

func convertOrderModeltoOrderDTO(order domain.Order) OrderRecordDTO {
	var prodcts []ProductRecordDTO
	for k, v := range order.ProductsQuantity {
		prodcts = append(prodcts, ProductRecordDTO{Product: k, Cost: int16(order.ProductsCost[k]), Quantity: v})
	}
	return OrderRecordDTO{
		UserID:    order.UserID,
		Status:    order.Status,
		Products:  prodcts,
		TotalCost: int16(order.TotalCost),
	}
}
