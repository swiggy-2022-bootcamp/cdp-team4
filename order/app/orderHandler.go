package app

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
)

type OrderHandler struct {
	OrderService domain.OrderService
}

func (oh OrderHandler) handlePay() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (oh OrderHandler) handleGetPayRecordByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (oh OrderHandler) handleGetPayRecordsByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (oh OrderHandler) handleUpdatePayStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
