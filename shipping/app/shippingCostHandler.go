package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
)

type ShippingCostRecordDTO struct {
	City string `json:"city"`
	Cost int    `json:"cost"`
}

// Create Order
// @Summary      Create Order
// @Description  This Handle allows admin to create new Order
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /order    [post]
func (sh ShippingHandler) handleShippingCost() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var scDto ShippingCostRecordDTO
		if err := ctx.BindJSON(&scDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := sh.ShippingCostService.CreateShippingCost(scDto.City, scDto.Cost)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": res})
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
func (sh ShippingHandler) HandleGetShippingCostByCity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		city := ctx.Param("city")
		res, err := sh.ShippingCostService.GetShippingCostByCity(city)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": res})
	}
}

// Update order status
// @Summary      Update order status
// @Description  This Handle Update order status given order id
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /orders    [put]
func (sh ShippingHandler) HandleUpdateShippingCostByCity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var scDto ShippingCostRecordDTO
		if err := ctx.BindJSON(&scDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		newshipCost := domain.ShippingCost{
			City:         scDto.City,
			ShippingCost: scDto.Cost,
		}

		ok, err := sh.ShippingCostService.UpdateShippingCost(newshipCost)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Shipping Cost record Updated"})
	}
}

// Delete order
// @Summary      Delete order
// @Description  This Handle deletes order given order id
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /orders    [delete]
func (sh ShippingHandler) HandleDeleteShippingCostByCity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		city := ctx.Param("city")
		_, err := sh.ShippingCostService.DeleteShippingCostByCity(city)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted Succesfully"})
	}
}
