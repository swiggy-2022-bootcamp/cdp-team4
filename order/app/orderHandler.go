package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/order/app/protobuf"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
	"google.golang.org/grpc"
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
	OrderID   string             `json:"order_id"`
	Status    string             `json:"status"`
	Products  []ProductRecordDTO `json:"products"`
	TotalCost int16              `json:"total_cost"`
}

type OrderConfirmResponseDTO struct {
	UserID                string `json:"user_id"`
	OrderID               string `json:"order_id"`
	Status                string `json:"status"`
	TotalCost             int16  `json:"total_cost"`
	ShippingPrice         int16  `json:"shipping_price"`
	RewardspointsConsumed int16  `json:"reward_points"`
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
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("Error while creating order")
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
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("Error while creating order")
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
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
		if id == "" {
			log.WithFields(logrus.Fields{"message": "Invalid ID", "status": http.StatusBadRequest}).
				Error("Error while Getting order by order id")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		fmt.Println(id)
		res, err := oh.OrderService.GetOrderById(id)

		if err != nil {
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("Record not found")
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
		if id == "" {
			log.WithFields(logrus.Fields{"message": "Invalid ID", "status": http.StatusBadRequest}).
				Error("Error while getting order by user id")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		fmt.Println(id)
		record, err := oh.OrderService.GetOrderByUserId(id)

		if err != nil {
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("Record not found")
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
		if status == "" {
			log.WithFields(logrus.Fields{"message": "No Status param", "status": http.StatusBadRequest}).
				Error("Record not found")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "No Status param"})
			return
		}
		record, err := oh.OrderService.GetOrderByStatus(status)

		if err != nil {
			log.WithFields(logrus.Fields{"message": "Record not found", "status": http.StatusBadRequest}).
				Error("Record not found")
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
			log.WithFields(logrus.Fields{"message": "Record not found", "status": http.StatusBadRequest}).
				Error("Record not found")
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
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ok, err := oh.OrderService.UpdateOrderStatus(requestDTO.Id, requestDTO.Status)
		if !ok {
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error(err.Error())
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
		if id == "" {
			log.WithFields(logrus.Fields{"message": "Invalid ID", "status": http.StatusBadRequest}).
				Error("Invalid ID")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		fmt.Println(id)
		_, err := oh.OrderService.DeleteOrderById(id)

		if err != nil {
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadGateway}).
				Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted Succesfully"})
	}
}

func (oh OrderHandler) HandleAddOrderFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("user_id")
		if id == "" {
			log.WithFields(logrus.Fields{"message": "Invalid ID", "status": http.StatusBadRequest}).
				Error("Invalid ID")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		conn, err := grpc.Dial("localhost:7899", grpc.WithInsecure())
		if err != nil {
			log.WithFields(logrus.Fields{"message": "Error while making connection", "status": http.StatusBadGateway}).
				Error("Error while making connection")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error while making connection"})
			return
		}

		// Create a client instance
		c := pb.NewCheckoutClient(conn)
		resp, err := c.OrderOverview(context.Background(), &pb.OverviewRequest{
			UserID: id,
		})
		if err != nil {
			log.WithFields(logrus.Fields{"message": "Error while getting grpc response", "status": http.StatusBadGateway}).
				Error("Error while getting grpc response")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error while getting grpc response"})
			return
		}
		var newpdto []ProductRecordDTO
		for _, prod := range resp.Item {
			newpdto = append(newpdto, ProductRecordDTO{
				Product:  prod.Name,
				Quantity: int(prod.Qty),
				Cost:     int16(prod.Price),
			})
		}
		product_quantity, product_cost := convertProductsDTOtoMaps(newpdto)
		res, result := oh.OrderService.CreateOrder(
			resp.UserID,
			"Pending",
			product_quantity,
			product_cost,
			int(resp.TotalPrice),
		)
		if result != nil {
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadGateway}).
				Error("Error while creating order")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error while creating order"})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": OrderConfirmResponseDTO{
			UserID:                resp.UserID,
			OrderID:               res,
			Status:                "Pending",
			TotalCost:             int16(resp.TotalPrice),
			ShippingPrice:         int16(resp.ShippingPrice),
			RewardspointsConsumed: int16(resp.RewardPointsConsumed),
		}})
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
		OrderID:   order.ID,
		TotalCost: int16(order.TotalCost),
	}
}
