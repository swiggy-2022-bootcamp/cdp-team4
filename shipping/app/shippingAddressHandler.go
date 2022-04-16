package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
)

type ShippingHandler struct {
	ShippingAddressService domain.ShippingAddressService
	ShippingCostService    domain.ShippingCostService
}

type ShippingAddressRecordDTO struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	CountryID int    `json:"country_id"`
	PostCode  int    `json:"postcode"`
}

// Create Shipping Address
// @Summary      Create Shipping Address
// @Description  This Handler allow user to create new Shipping Address
// @Tags         Shipping Address
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress    [post]
func (sh ShippingHandler) handleShippingAddress() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var saDto ShippingAddressRecordDTO
		if err := ctx.BindJSON(&saDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := sh.ShippingAddressService.CreateShippingAddress(
			saDto.FirstName,
			saDto.LastName,
			saDto.City,
			saDto.Address1,
			saDto.Address2,
			saDto.CountryID,
			saDto.PostCode,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Shipping Address Record Added", "Shipping Address id": res})
	}
}

// Get Shipping Address by Id
// @Summary      Get Shipping Address by id
// @Description  This Handle returns shippingAddress given id
// @Tags         Shipping Address
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress/:id    [get]
func (sh ShippingHandler) HandleGetShippingAddrssByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
		res, err := sh.ShippingAddressService.GetShippingAddressById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": res})
	}
}

// Update Shipping Address
// @Summary      Update Shipping Address
// @Description  This Handle Update shippingAddress given id
// @Tags         Shipping Address
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress/:id     [put]
func (sh ShippingHandler) HandleUpdateShippingAddressByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var saDto ShippingAddressRecordDTO
		id := ctx.Param("id")
		if err := ctx.BindJSON(&saDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		newshipAddr := convertShippingAddressDTOtoShippingAddressModel(saDto)

		ok, err := sh.ShippingAddressService.UpdateShippingAddressById(id, newshipAddr)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Shipping Address record updated"})
	}
}

// Delete Shipping Address
// @Summary      Delete Shipping Address
// @Description  This Handle deletes Delete Shipping Address given sid
// @Tags         Shipping Address
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress/:id   [delete]
func (sh ShippingHandler) HandleDeleteShippingAddressById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
		_, err := sh.ShippingAddressService.DeleteShippingAddressById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted Succesfully"})
	}
}

func convertShippingAddressDTOtoShippingAddressModel(saDto ShippingAddressRecordDTO) domain.ShippingAddress {

	return domain.ShippingAddress{
		FirstName: saDto.FirstName,
		LastName:  saDto.LastName,
		City:      saDto.City,
		Address1:  saDto.Address1,
		Address2:  saDto.Address2,
		PostCode:  saDto.PostCode,
		CountryID: saDto.CountryID,
	}
}
