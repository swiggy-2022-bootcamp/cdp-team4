package app

import (
	"net/http"
	"strings"

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

type PaymentMethodDTO struct {
	Id      string
	Method  string
	Agree   string
	Comment string
}

func (ph PayHandler) HandlePay() gin.HandlerFunc {
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

func (ph PayHandler) HandleGetPayRecordByID() gin.HandlerFunc {
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
		id := ctx.Param("user_id")
		record, err := ph.PaymentService.GetPaymentAllRecordsByUserId(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": record})
	}
}

func (ph PayHandler) handleUpdatePayStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestDTO struct {
			Id     string
			Status string
		}
		if err := ctx.BindJSON(&requestDTO); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ok, err := ph.PaymentService.UpdatePaymentStatus(requestDTO.Id, requestDTO.Status)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "payment record updated"})
	}
}

func (ph PayHandler) handleAddPaymentMethods() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paymentMethodDTO PaymentMethodDTO
		if err := ctx.BindJSON(&paymentMethodDTO); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ok, err := ph.PaymentService.AddPaymentMethod(
			paymentMethodDTO.Id,
			paymentMethodDTO.Method,
			paymentMethodDTO.Agree,
			paymentMethodDTO.Comment,
		)

		if !ok {
			if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "method already exists"})
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "payment method added"})
	}
}

func (ph PayHandler) handleGetPaymentMethods() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		methods, err := ph.PaymentService.GetPaymentMethods(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": methods})
	}
}
