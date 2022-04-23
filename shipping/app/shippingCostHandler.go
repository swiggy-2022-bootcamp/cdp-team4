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

// Create Shipping Cost
// @Summary      Create Shipping Cost
// @Description  This Handle allows admin to create new Shipping Cost of a city
// @Tags         Shipping Cost
// @Produce      json
// @Param 		 shippingCost body ShippingCostRecordDTO true "Create Shipping Cost"
// @Success      202  {string}  true
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingcost    [post]
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

// Get Shipping Cost given city
// @Summary      Get Shipping Cost given city
// @Description  This Handle returns Shipping Cost given city
// @Tags         Shipping Cost
// @Produce      json
// @Param        city   path    int  true  "shipping cost city"
// @Success      202  {object}  ShippingCostRecordDTO
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingcost/:city    [get]
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

// Update Shipping Cost
// @Summary      Update Shipping Cost
// @Description  This Handle Update allows admin to update shipping cost of a city
// @Tags         Shipping Cost
// @Produce      json
// @Param 		 shippingCost body ShippingCostRecordDTO true "Update Shipping Cost"
// @Success      202  {string}  Shipping Cost record Updated
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingcost    [put]
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

// Delete Shipping Cost
// @Summary      Delete Shipping Cost
// @Description  This Handle deletes Shipping Cost of a city
// @Tags         Shipping Cost
// @Produce      json
// @Param        city   path    int  true  "shipping cost city"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingcost    [delete]
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
