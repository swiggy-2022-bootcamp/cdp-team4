package app

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	// "log"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

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

type orderHandler struct {
}

func (oh orderHandler) GetOrderByID(c *gin.Context) {

	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}

	orderServiceURI := os.Getenv("ORDER_SERVICE_URI") + "/order/" + c.Param("id")

	req, err := http.NewRequest("GET", orderServiceURI, nil)
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var orderResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&orderResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(orderResDTO.Status, orderResDTO)
}

func (oh orderHandler) GetOrderByUserID(c *gin.Context) {

	orderServiceURI := os.Getenv("ORDER_SERVICE_URI") + "/order/user/" + c.Param("user_id")

	req, err := http.NewRequest("GET", orderServiceURI, nil)
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var orderResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&orderResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(orderResDTO.Status, orderResDTO)
}

func (oh orderHandler) GetOrderByStatus(c *gin.Context) {

	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}

	orderServiceURI := os.Getenv("ORDER_SERVICE_URI") + "/order/status/" + c.Param("status")

	req, err := http.NewRequest("GET", orderServiceURI, nil)
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var orderResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&orderResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(orderResDTO.Status, orderResDTO)
}

func (oh orderHandler) GetOrderInvoice(c *gin.Context) {

	orderServiceURI := os.Getenv("ORDER_SERVICE_URI") + "/order/invoice/" + c.Param("order_id")

	req, err := http.NewRequest("GET", orderServiceURI, nil)
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var orderResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&orderResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(orderResDTO.Status, orderResDTO)
}

func (oh orderHandler) GetAllOrders(c *gin.Context) {

	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}

	orderServiceURI := os.Getenv("ORDER_SERVICE_URI") + "/orders"

	req, err := http.NewRequest("GET", orderServiceURI, nil)
	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var orderResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&orderResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(orderResDTO.Status, orderResDTO)
}

func (oh orderHandler) UpdateOrder(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	orderServiceURI := os.Getenv("ORDER_SERVICE_URI") + "/order/status"

	var updateOrder RequestDTO

	if err := c.BindJSON(&updateOrder); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("here1: %v", updateOrder)

	orderJSON, err1 := json.Marshal(updateOrder)

	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("PUT", orderServiceURI, bytes.NewBuffer(orderJSON))

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var orderResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&orderResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(orderResDTO.Status, orderResDTO)
}

func (oh orderHandler) DeleteOrder(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	orderServiceURI := os.Getenv("ORDER_SERVICE_URI") + "/order/" + c.Param("id")

	req, err2 := http.NewRequest("DELETE", orderServiceURI, nil)

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var orderResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&orderResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(orderResDTO.Status, orderResDTO)
	c.Abort()
}

func (oh orderHandler) CreateOrder(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	orderServiceURI := os.Getenv("ORDER_SERVICE_URI") + "/order"

	var newOrder OrderRecordDTO

	if err := c.BindJSON(&newOrder); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("here1: %v", newOrder)

	orderJSON, err1 := json.Marshal(newOrder)

	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("POST", orderServiceURI, bytes.NewBuffer(orderJSON))

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var orderResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&orderResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(orderResDTO.Status, orderResDTO)
}

func (oh orderHandler) ConfirmOrder(c *gin.Context) {
	orderServiceURI := os.Getenv("ORDER_SERVICE_URI") + "/confirm/" + c.Param("user_id")

	var newOrder OrderRecordDTO

	if err := c.BindJSON(&newOrder); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("here1: %v", newOrder)

	orderJSON, err1 := json.Marshal(newOrder)

	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("POST", orderServiceURI, bytes.NewBuffer(orderJSON))

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err1 := client.Do(req)
	if err1 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	var orderResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&orderResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(orderResDTO.Status, orderResDTO)
}
