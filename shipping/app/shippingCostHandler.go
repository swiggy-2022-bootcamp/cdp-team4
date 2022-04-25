package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
// @Success      200  {string}  true
// @Failure      500  {number} 	http.StatusBadRequest
// @Router       /shippingcost    [post]
func (sh ShippingHandler) handleShippingCost() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var scDto ShippingCostRecordDTO
		if err := ctx.BindJSON(&scDto); err != nil {
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

		res, err := sh.ShippingCostService.CreateShippingCost(scDto.City, scDto.Cost)

		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("create Shipping Cost Record")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		responseDto := ResponseDTO{
			Status:  http.StatusOK,
			Message: "Shipping Cost Record Created",
		}
		log.WithFields(logrus.Fields{"data": res, "status": http.StatusCreated}).
			Info("Shipping Cost Record Created")
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Get Shipping Cost given city
// @Summary      Get Shipping Cost given city
// @Description  This Handle returns Shipping Cost given city
// @Tags         Shipping Cost
// @Produce      json
// @Param        city   path    int  true  "shipping cost city"
// @Success      200  {object}  ShippingCostRecordDTO
// @Failure      500  {number} 	http.StatusBadRequest
// @Router       /shippingcost/:city    [get]
func (sh ShippingHandler) HandleGetShippingCostByCity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		city := ctx.Param("city")
		res, err := sh.ShippingCostService.GetShippingCostByCity(city)

		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error Record Not Found")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		responseDto := ResponseDTO{
			Status: http.StatusOK,
			Data:   res,
		}
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Update Shipping Cost
// @Summary      Update Shipping Cost
// @Description  This Handle Update allows admin to update shipping cost of a city
// @Tags         Shipping Cost
// @Produce      json
// @Param 		 shippingCost body ShippingCostRecordDTO true "Update Shipping Cost"
// @Success      200  {string}  Shipping Cost record Updated
// @Failure      500  {number} 	http.StatusBadRequest
// @Router       /shippingcost    [put]
func (sh ShippingHandler) HandleUpdateShippingCostByCity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var scDto ShippingCostRecordDTO
		if err := ctx.BindJSON(&scDto); err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error binding")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		newshipCost := domain.ShippingCost{
			City:         scDto.City,
			ShippingCost: scDto.Cost,
		}

		ok, err := sh.ShippingCostService.UpdateShippingCost(newshipCost)
		if !ok {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error Updating Record")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}
		responseDto := ResponseDTO{
			Status:  http.StatusOK,
			Message: "Record Updated",
		}
		ctx.JSON(responseDto.Status, responseDto)
	}
}

// Delete Shipping Cost
// @Summary      Delete Shipping Cost
// @Description  This Handle deletes Shipping Cost of a city
// @Tags         Shipping Cost
// @Produce      json
// @Param        city   path    int  true  "shipping cost city"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {number} 	http.StatusBadRequest
// @Router       /shippingcost    [delete]
func (sh ShippingHandler) HandleDeleteShippingCostByCity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		city := ctx.Param("city")
		_, err := sh.ShippingCostService.DeleteShippingCostByCity(city)

		if err != nil {
			responseDto := ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Message,
			}
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("Error Deleting Record")
			ctx.JSON(responseDto.Status, responseDto)
			ctx.Abort()
			return
		}

		responseDto := ResponseDTO{
			Status:  http.StatusOK,
			Message: "Record Deleted",
		}
		ctx.JSON(responseDto.Status, responseDto)
	}
}
