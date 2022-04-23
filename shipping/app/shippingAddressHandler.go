package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
// @Param 		 shippingAddress body ShippingAddressRecordDTO true "Create Shipping Address"
// @Success		 202  string    Shipping Address record added
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingadress    [post]
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
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("create Shipping Address Record")
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
// @Param        id   path      int  true  "shipping address id"
// @Success      202  {object}  ShippingAddressRecordDTO
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress/:id    [get]
func (sh ShippingHandler) HandleGetShippingAddrssByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		//fmt.Println(id)
		res, err := sh.ShippingAddressService.GetShippingAddressById(id)

		if err != nil {
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("Get Shipping Address Record")
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
// @Param        id   path      int  true  "shipping address id"
// @Param 		 shippingAddress body ShippingAddressRecordDTO true "Update Shipping Address"
// @Success      202  {number}  http.StatusAccepted
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
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("Update Shipping Address Record")
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
// @Param        id   path      int  true  "shipping address id"
// @Success      202  {number}  http.StatusAccepted
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress/:id   [delete]
func (sh ShippingHandler) HandleDeleteShippingAddressById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		//fmt.Println(id)
		_, err := sh.ShippingAddressService.DeleteShippingAddressById(id)

		if err != nil {
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("Delete Shipping Address Record")
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

func NewShippingHandler(shippingAddressService domain.ShippingAddressService, shippingCostService domain.ShippingCostService) ShippingHandler {
	return ShippingHandler{
		ShippingAddressService: shippingAddressService,
		ShippingCostService:    shippingCostService,
	}
}
