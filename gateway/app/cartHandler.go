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

type CartProductRecordDTO struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Cost     int    `json:"cost"`
	Quantity int    `json:"quantity"`
}

type CartRecordDTO struct {
	UserID string                 `json:"user_id"`
	Carts  []CartProductRecordDTO `json:"carts"`
}

type CartHandler struct {
}

func (ch CartHandler) GetCartByUserID(c *gin.Context) {

	cartServiceURI := os.Getenv("CART_SERVICE_URI") + "/cart/" + c.Param("userId")

	req, err := http.NewRequest("GET", cartServiceURI, nil)
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

	var cartResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&cartResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	// c.JSON(cartResDTO.Status, cartResDTO)
	c.JSON(200, res.Body)
	c.Abort()
}

func (ch CartHandler) CreateCart(c *gin.Context) {

	cartServiceURI := os.Getenv("CART_SERVICE_URI") + "/cart/" + c.Param("userId")

	var newCart CartRecordDTO

	if err := c.BindJSON(&newCart); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("here1: %v", newCart)

	cartJSON, err1 := json.Marshal(newCart)
	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("POST", cartServiceURI, bytes.NewBuffer(cartJSON))
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

	var cartResponseDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&cartResponseDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	// c.JSON(cartResponseDTO.Status, cartResponseDTO)
	c.JSON(200, res.Body)
	c.Abort()
}

func (ch CartHandler) UpdateCartByUserID(c *gin.Context) {

	cartServiceURI := os.Getenv("CART_SERVICE_URI") + "/cart/" + c.Param("userId")
	var newCart CartRecordDTO
	if err := c.BindJSON(&newCart); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("New Cart after update: %v", newCart)

	cartJSON, err1 := json.Marshal(newCart)
	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}
	req, err := http.NewRequest("PUT", cartServiceURI, bytes.NewBuffer(cartJSON))
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

	var cartResponseDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&cartResponseDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	// c.JSON(cartResponseDTO.Status, cartResponseDTO)
	c.JSON(200, res.Body)
	c.Abort()
}

func (ch CartHandler) DeleteCartByUserID(c *gin.Context) {

	cartServiceURI := os.Getenv("CART_SERVICE_URI") + "/cart/empty/" + c.Param("userId")

	req, err := http.NewRequest("DELETE", cartServiceURI, nil)
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

	var cartResponseDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&cartResponseDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	//c.JSON(cartResponseDTO.Status, cartResponseDTO)
	c.JSON(200, res.Body)
	c.Abort()
}

func (ch CartHandler) DeleteCartItemByUserId(c *gin.Context) {

	cartServiceURI := os.Getenv("CART_SERVICE_URI") + "/cart/" + c.Param("userId")

	type ProductIdList struct {
		ProductList []string `json:"product_list"`
	}

	var productIdList ProductIdList
	if err := c.BindJSON(&productIdList); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("List of product ids to remove from the cart: %v", productIdList)

	cartJSON, err1 := json.Marshal(productIdList)
	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}
	req, err := http.NewRequest("DELETE", cartServiceURI, bytes.NewBuffer(cartJSON))
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

	var cartResponseDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&cartResponseDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	// c.JSON(cartResponseDTO.Status, cartResponseDTO)
	c.JSON(200, res.Body)
	c.Abort()
}
