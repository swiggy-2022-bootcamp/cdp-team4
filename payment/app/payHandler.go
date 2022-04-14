package app

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
)

type PayHandler struct {
	PaymentService domain.PaymentService
}

func (ph PayHandler) handlePay() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (ph PayHandler) handleGetPayRecordByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

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
