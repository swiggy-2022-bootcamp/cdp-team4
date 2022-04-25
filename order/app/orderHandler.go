package app

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/order/app/protobuf"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
	"google.golang.org/grpc"
)

type OrderHandler struct {
	OrderService         domain.OrderService
	OrderOverviewService domain.OrderOverviewService
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

type InvoiceDTO struct {
	UserID                string             `json:"user_id"`
	Products              []ProductRecordDTO `json:"products"`
	Status                string             `json:"status"`
	TotalCost             int16              `json:"total_cost"`
	ShippingPrice         int16              `json:"shipping_price"`
	RewardspointsConsumed int16              `json:"reward_points"`
}

type OrderOverviewRecordDTO struct {
	OrderID  string         `json:"order_id"`
	Products map[string]int `json:"products"`
}

type RequestDTO struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

// Create Order
// @Summary      Create Order
// @Description  This Handle allows admin to create new Order
// @Tags         Order
// @Produce      json
// @Param 		 order body OrderRecordDTO true "Create order"
// @Success      200  {number}  http.StatusAccepted
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /order    [post]
func (oh OrderHandler) handleOrder() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var orderDto OrderRecordDTO
		if err := ctx.BindJSON(&orderDto); err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("bind json")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}

		product_quantity, product_cost := ConvertProductsDTOtoMaps(orderDto.Products)

		res, err := oh.OrderService.CreateOrder(
			orderDto.UserID,
			orderDto.Status,
			product_quantity,
			product_cost,
			int(orderDto.TotalCost),
		)
		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error while creating order!")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		responseDto := ResponseDTO{
			Status:  http.StatusOK,
			Data:    res,
			Message: "Order ID",
		}
		log.WithFields(logrus.Fields{"data": res, "status": http.StatusCreated}).
			Info("Order Added")
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Get Order by ID
// @Summary      Get Order by id
// @Description  This Handle returns Order given order id
// @Tags         Order
// @Produce      json
// @Param        id   path      int  true  "order id"
// @Success      200  {object}  OrderRecordDTO
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /order/:id    [get]
func (oh OrderHandler) HandleGetOrderRecordByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		res, err := oh.OrderService.GetOrderById(id)

		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Record not found!")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		orderdto := convertOrderModeltoOrderDTO(*res)
		responseDto := ResponseDTO{
			Status: http.StatusOK,
			Data:   orderdto,
		}
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Get Order by User ID
// @Summary      Get Order by user id
// @Description  This Handle returns Order given user id
// @Tags         Order
// @Produce      json
// @Param        user_id   path      int  true  "user id"
// @Success      200  {object}  []OrderRecordDTO
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /order/user/:userid    [get]
func (oh OrderHandler) HandleGetOrderRecordsByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("user_id")
		record, err := oh.OrderService.GetOrderByUserId(id)

		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Record not found!")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}

		orderdto := make([]OrderRecordDTO, 0)

		for _, rec := range record {
			respdto := convertOrderModeltoOrderDTO(rec)
			orderdto = append(orderdto, respdto)
		}

		responseDto := ResponseDTO{
			Status: http.StatusOK,
			Data:   orderdto,
		}
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Get Order by Status
// @Summary      Get Order by status
// @Description  This Handle returns Order given status
// @Tags         Order
// @Produce      json
// @Param        id   path    int  true  "status"
// @Success      200  {object}  []OrderRecordDTO
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /order/status/:status    [get]
func (oh OrderHandler) HandleGetOrderRecordsByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Param("status")
		record, err := oh.OrderService.GetOrderByStatus(status)

		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Record not found!")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}

		orderdto := make([]OrderRecordDTO, 0)

		for _, rec := range record {
			respdto := convertOrderModeltoOrderDTO(rec)
			orderdto = append(orderdto, respdto)
		}

		responseDto := ResponseDTO{
			Status: http.StatusOK,
			Data:   orderdto,
		}
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Get All Order records
// @Summary      Get All Order records
// @Description  This Handle returns all of the orders
// @Tags         Order
// @Produce      json
// @Param        id   path      int  true  "order id"
// @Success      200  {object}  []OrderRecordDTO
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /orders    [get]
func (oh OrderHandler) HandleGetAllRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		record, err := oh.OrderService.GetAllOrders()
		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Records not found!")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		orderdto := make([]OrderRecordDTO, 0)

		for _, rec := range record {
			respdto := convertOrderModeltoOrderDTO(rec)
			orderdto = append(orderdto, respdto)
		}

		responseDto := ResponseDTO{
			Status: http.StatusOK,
			Data:   orderdto,
		}
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Update order status
// @Summary      Update order status
// @Description  This Handle Update order status given order id
// @Tags         Order
// @Produce      json
// @Param 		 order body RequestDTO true "Update order"
// @Success      200  {number}  http.StatusAccepted
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /order/status    [put]
func (oh OrderHandler) handleUpdateOrderStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestDTO RequestDTO
		if err := ctx.BindJSON(&requestDTO); err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Bind Json")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}

		ok, err := oh.OrderService.UpdateOrderStatus(requestDTO.Id, requestDTO.Status)
		if !ok {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Unable to Update Order")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		responseDto := ResponseDTO{
			Status:  http.StatusOK,
			Message: "Order Status Updated",
		}
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Delete order
// @Summary      Delete order
// @Description  This Handle deletes order given order id
// @Tags         Order
// @Produce      json
// @Param        id   path      int  true  "order id"
// @Success      200  {number}  http.StatusAccepted
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /order/:id   [delete]
func (oh OrderHandler) HandleDeleteOrderById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		_, err := oh.OrderService.DeleteOrderById(id)

		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error Deleting Order")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		responseDto := ResponseDTO{
			Status:  http.StatusOK,
			Message: "Order Deleted Succesfully",
		}
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Confirm Order
// @Summary      Confirm Order
// @Description  This Handle adds order from checkout
// @Tags         Order
// @Produce      json
// @Param        userid   path      int  true  "user id"
// @Success      200  {object}  OrderConfirmResponseDTO
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /confirm/:userid   [post]
func (oh OrderHandler) HandleAddOrderFromCheckout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("user_id")

		err := godotenv.Load(".env")
		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Internal env file failed to load")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		PORT := os.Getenv("CHECKOUT_SERVICE_PORT")
		conn, err := grpc.Dial("0.0.0.0:"+PORT, grpc.WithInsecure())
		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			log.WithFields(logrus.Fields{"message": "Error while making connection", "status": http.StatusInternalServerError}).
				Error("Error while making grpc connection")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}

		// Create a client instance
		c := pb.NewCheckoutClient(conn)
		resp, err := c.OrderOverview(context.Background(), &pb.OverviewRequest{
			UserID: id,
		})
		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			log.WithFields(logrus.Fields{"message": "Error while getting grpc response", "status": http.StatusInternalServerError}).
				Error("Error while getting grpc response")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
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
		product_quantity, product_cost := ConvertProductsDTOtoMaps(newpdto)
		res, result := oh.OrderService.CreateOrder(
			resp.UserID,
			"Pending",
			product_quantity,
			product_cost,
			int(resp.TotalPrice),
		)
		if result != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error while creating order record")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}

		prod_id_quantity := make(map[string]int)

		for _, item := range resp.Item {
			prod_id_quantity[item.Id] = int(item.Qty)
		}

		neworder := domain.OrderOverview{
			OrderID:            res,
			ProductsIdQuantity: prod_id_quantity,
		}

		_, err1 := oh.OrderOverviewService.CreateOrderOverview(neworder)
		if err1 != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error while creating order overview record")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}

		responseDto := ResponseDTO{
			Status: http.StatusOK,
			Data: OrderConfirmResponseDTO{
				UserID:                resp.UserID,
				OrderID:               res,
				Status:                "Pending",
				TotalCost:             int16(resp.TotalPrice),
				ShippingPrice:         int16(resp.ShippingPrice),
				RewardspointsConsumed: int16(resp.RewardPointsConsumed),
			},
		}
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Get Order Invoice
// @Summary      Get Order Invoice given order id
// @Description  This generated invoice given order id
// @Tags         Order
// @Produce      json
// @Param        orderid   path      int  true  "order id"
// @Success      200  {object}  InvoiceDTO
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /order/invoice/:orderid   [get]
func (oh OrderHandler) HandleGetOrderInvoice() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("order_id")

		order, err := oh.OrderService.GetOrderById(id)
		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error while retriving order record for invoice")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}

		err1 := godotenv.Load(".env")
		if err1 != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error while retriving order record for invoice, .env file unable to load")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		PORT := os.Getenv("CHECKOUT_SERVICE_PORT")
		conn, err2 := grpc.Dial("localhost:"+PORT, grpc.WithInsecure())
		if err2 != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error while making grpc connection to checkout")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}

		// Create a client instance
		c := pb.NewCheckoutClient(conn)
		resp, err3 := c.OrderOverview(context.Background(), &pb.OverviewRequest{
			UserID: order.UserID,
		})
		if err3 != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error while making grpc getting grpc response from checkout")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		var products []ProductRecordDTO
		for _, item := range resp.Item {
			newproduct := ProductRecordDTO{
				Product:  item.Name,
				Quantity: int(item.Qty),
				Cost:     int16(item.Price),
			}
			products = append(products, newproduct)
		}
		ctx.JSON(http.StatusAccepted, gin.H{
			"Invoice": InvoiceDTO{
				Products:              products,
				UserID:                order.UserID,
				Status:                order.Status,
				TotalCost:             int16(order.TotalCost),
				RewardspointsConsumed: int16(resp.RewardPointsConsumed),
				ShippingPrice:         int16(resp.ShippingPrice),
			},
		})

		responseDto := ResponseDTO{
			Status: http.StatusOK,
			Data: InvoiceDTO{
				Products:              products,
				UserID:                order.UserID,
				Status:                order.Status,
				TotalCost:             int16(order.TotalCost),
				RewardspointsConsumed: int16(resp.RewardPointsConsumed),
				ShippingPrice:         int16(resp.ShippingPrice),
			},
			Message: "Invoice",
		}
		ctx.JSON(responseDto.Status, responseDto)
	}

}

// Function that converts Product DTO to products quantity and products cost map which
// are directly stored in dynamodb
func ConvertProductsDTOtoMaps(products []ProductRecordDTO) (map[string]int, map[string]int) {
	var product_quantity map[string]int = make(map[string]int)
	var product_cost map[string]int = make(map[string]int)
	for _, product := range products {
		product_quantity[product.Product] = product.Quantity
		product_cost[product.Product] = int(product.Cost)
	}
	return product_quantity, product_cost
}

// Function that converts Order Model to Order DTO which is
// returned as a response to the caller
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

func NewOrderHandler(orderService domain.OrderService, orderOverviewService domain.OrderOverviewService) OrderHandler {
	return OrderHandler{
		OrderService:         orderService,
		OrderOverviewService: orderOverviewService,
	}
}
