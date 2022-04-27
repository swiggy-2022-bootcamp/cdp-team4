package app

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type PaymentRecordDTO struct {
	Amount      int16
	Currency    string
	Status      string
	OrderID     string
	UserID      string
	Method      string
	Description string
	VPA         string
	Notes       []string
}

type PaymentMethodDTO struct {
	Id      string
	Method  string
	Agree   string
	Comment string
}

type PaymentHandler struct {
}

func (ph PaymentHandler) InitiatePayment(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "true" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	var paymentDto PaymentRecordDTO

	if err := c.BindJSON(&paymentDto); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	paymentServiceURI := os.Getenv("PAYMENT_SERVICE_URI") + "/pay"
	paymentJSON, _ := json.Marshal(paymentDto)

	req, err := http.NewRequest("POST", paymentServiceURI, bytes.NewBuffer(paymentJSON))
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

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(200, res.Body)
	c.Abort()
}

func (ph PaymentHandler) AddPaymentMethod(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "true" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	var paymentDto PaymentMethodDTO

	if err := c.BindJSON(&paymentDto); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	paymentServiceURI := os.Getenv("PAYMENT_SERVICE_URI") + "/pay/paymentMethods"
	paymentJSON, _ := json.Marshal(paymentDto)

	req, err := http.NewRequest("POST", paymentServiceURI, bytes.NewBuffer(paymentJSON))
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

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(200, res.Body)
	c.Abort()
}

func (ph PaymentHandler) GetPaymentMethod(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "true" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}

	paymentServiceURI := os.Getenv("PAYMENT_SERVICE_URI") + "/pay/paymentMethods/" + c.Param("userId")

	req, err := http.NewRequest("GET", paymentServiceURI, nil)
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

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(200, res.Body)
	c.Abort()
}
