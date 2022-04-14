package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
)

type PayHandler struct {
	PaymentService domain.PaymentService
}

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

func (ph PayHandler) handlePay() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paymentDto PaymentRecordDTO

		if err := ctx.BindJSON(&paymentDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ok, err := ph.PaymentService.CreateDynamoPaymentRecord(
			paymentDto.Amount,
			paymentDto.Currency,
			paymentDto.Status,
			paymentDto.OrderID,
			paymentDto.UserID,
			paymentDto.Method,
			paymentDto.Description,
			paymentDto.VPA,
			paymentDto.Notes,
		)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "payment record added"})
	}
}

func (ph PayHandler) handleGetPayRecordByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		record, err := ph.PaymentService.GetPaymentRecordById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": record})
	}
}

func (ph PayHandler) handleGetPayRecordsByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (ph PayHandler) handleUpdatePayStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (ph PayHandler) handleAddPaymentMethods() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (ph PayHandler) handleGetPaymentMethods() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
