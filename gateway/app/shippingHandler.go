package app

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	// "log"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type ShippingAddressRecordDTO struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	CountryID int    `json:"country_id"`
	PostCode  int    `json:"postcode"`
}

type ShippingCostRecordDTO struct {
	City string `json:"city"`
	Cost int    `json:"cost"`
}

type shippingHandler struct {
}

func (sh shippingHandler) GetShippingAddressByID(c *gin.Context) {

	shippingAddressServiceURI := os.Getenv("SHIPPING_SERVICE_URI") + "/shippingaddress/" + c.Param("id")

	req, err := http.NewRequest("GET", shippingAddressServiceURI, nil)
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

	var shippingAddressResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&shippingAddressResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(shippingAddressResDTO.Status, shippingAddressResDTO)
}

func (sh shippingHandler) UpdateShippingAddress(c *gin.Context) {
	shippingAddressServiceURI := os.Getenv("SHIPPING_SERVICE_URI") + "/shippingaddress/" + c.Param("id")

	var updateShippingAddress ShippingAddressRecordDTO

	if err := c.BindJSON(&updateShippingAddress); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("here1: %v", updateShippingAddress)

	shippingAddressJSON, err1 := json.Marshal(updateShippingAddress)

	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("PUT", shippingAddressServiceURI, bytes.NewBuffer(shippingAddressJSON))

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
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

	var shippingAddressResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&shippingAddressResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(shippingAddressResDTO.Status, shippingAddressResDTO)
}

func (sh shippingHandler) DeleteShippingAddress(c *gin.Context) {
	shippingAddressServiceURI := os.Getenv("SHIPPING_SERVICE_URI") + "/shippingaddress/" + c.Param("id")

	req, err2 := http.NewRequest("DELETE", shippingAddressServiceURI, nil)

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
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

	var shippingAddressResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&shippingAddressResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(shippingAddressResDTO.Status, shippingAddressResDTO)
	c.Abort()
}

func (sh shippingHandler) CreateShippingAddress(c *gin.Context) {
	shippingAddressServiceURI := os.Getenv("SHIPPING_SERVICE_URI") + "/shippingaddress"

	var newShippingAddress ShippingAddressRecordDTO

	if err := c.BindJSON(&newShippingAddress); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("here1: %v", newShippingAddress)

	shippingAddressJSON, err1 := json.Marshal(newShippingAddress)

	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("POST", shippingAddressServiceURI, bytes.NewBuffer(shippingAddressJSON))

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
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

	var shippingAddressResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&shippingAddressResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(shippingAddressResDTO.Status, shippingAddressResDTO)
}

func (sh shippingHandler) GetShippingCostByCity(c *gin.Context) {

	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}

	shippingCostServiceURI := os.Getenv("SHIPPING_SERVICE_URI") + "/shippingcost/" + c.Param("city")

	req, err := http.NewRequest("GET", shippingCostServiceURI, nil)
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

	var shippingCostResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&shippingCostResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(shippingCostResDTO.Status, shippingCostResDTO)
}

func (sh shippingHandler) UpdateShippingCost(c *gin.Context) {

	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}

	shippingCostServiceURI := os.Getenv("SHIPPING_SERVICE_URI") + "/shippingcost"

	var updateShippingCost ShippingCostRecordDTO

	if err := c.BindJSON(&updateShippingCost); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("here1: %v", updateShippingCost)

	shippingCostJSON, err1 := json.Marshal(updateShippingCost)

	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("PUT", shippingCostServiceURI, bytes.NewBuffer(shippingCostJSON))

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
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

	var shippingCostsResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&shippingCostsResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(shippingCostsResDTO.Status, shippingCostsResDTO)
}

func (sh shippingHandler) DeleteShippingCost(c *gin.Context) {
	shippingCostServiceURI := os.Getenv("SHIPPING_SERVICE_URI") + "/shippingcost/" + c.Param("city")

	req, err2 := http.NewRequest("DELETE", shippingCostServiceURI, nil)

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
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

	var shippingCostResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&shippingCostResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(shippingCostResDTO.Status, shippingCostResDTO)
	c.Abort()
}

func (sh shippingHandler) CreateShippingCost(c *gin.Context) {
	shippingAddressServiceURI := os.Getenv("SHIPPING_SERVICE_URI") + "/shippingcost"

	var newShippingCost ShippingCostRecordDTO

	if err := c.BindJSON(&newShippingCost); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("here1: %v", newShippingCost)

	shippingCostJSON, err1 := json.Marshal(newShippingCost)

	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("POST", shippingAddressServiceURI, bytes.NewBuffer(shippingCostJSON))

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
	}

	fmt.Printf("resp: %v", req)

	if err2 != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
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

	var shippingCostResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&shippingCostResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(shippingCostResDTO.Status, shippingCostResDTO)
}
