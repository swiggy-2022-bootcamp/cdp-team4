package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	// "log"
	"fmt"
	"encoding/json"
	"context"
	"time"
	"bytes"
)

type ShippingAddressDTO struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	City      string    `json:"city"`
	Address1  string    `json:"address_1"`
	Address2  string    `json:"address_2"`
	CountryID uint32    `json:"country_id"`
	PostCode  uint32    `json:"postcode"`
}

type UserDTO struct {
	FirstName string 				`json:"first_name"`
	LastName  string 				`json:"last_name"`
	Username  string 				`json:"username"`
	Password  string 				`json:"password"`
	Phone     string 				`json:"phone"`
	Email     string 				`json:"email"`
	Role      int    				`json:"role"`
	Address   ShippingAddressDTO 	`json:"address"`
	Fax		  string 				`json:"fax"`
}

type userHandler struct {
}

func (uh userHandler) GetAllUsers(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "true" {
		c.JSON(http.StatusOK, gin.H{"userId": c.Param("userId")})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
}

func (uh userHandler) GetUser(c *gin.Context) {

	userServiceURI := os.Getenv("USER_SERVICE_URI") + "/user/" + c.Param("userId")

	req, err := http.NewRequest("GET", userServiceURI, nil)
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

	var userResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&userResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(userResDTO.Status, userResDTO)
	c.Abort()
	return
}

func (uh userHandler) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"userId": c.Param("userId")})
	c.Abort()
	return
}

func (uh userHandler) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"userId": c.Param("userId")})
	c.Abort()
	return
}

func (uh userHandler) CreateUser(c *gin.Context) {
	userServiceURI := os.Getenv("USER_SERVICE_URI") + "/user" 

	var newUser UserDTO

	if err := c.BindJSON(&newUser); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	fmt.Printf("here1: %v", newUser)

	userJSON, err1 := json.Marshal(newUser)

	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("POST", userServiceURI, bytes.NewBuffer(userJSON))

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

	var userResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&userResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(userResDTO.Status, userResDTO)
	c.Abort()
	return
}
