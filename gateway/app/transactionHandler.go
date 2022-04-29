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

type TransactionRecordDTO struct {
	UserID            string `json:"user_id"`
	TransactionPoints int    `json:"transaction_points"`
}

type TransactionHandler struct {
}

func (th TransactionHandler) GetTransactionByUserID(c *gin.Context) {

	transactionServiceURI := os.Getenv("TRANSACTION_SERVICE_URI") + "/transaction/" + c.Param("userId")

	req, err := http.NewRequest("GET", transactionServiceURI, nil)
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

	var transactionResDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&transactionResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	c.JSON(200, res.Body)
	//c.JSON(transactionResDTO.Status, transactionResDTO)
}

func (th TransactionHandler) UpdateTransaction(c *gin.Context) {
	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	transactionServiceURI := os.Getenv("TRANSACTION_SERVICE_URI") + "/transaction/" + c.Param("userId")

	var updateTransaction TransactionRecordDTO

	if err := c.BindJSON(&updateTransaction); err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	updateTransaction.UserID = c.Param("userId")

	fmt.Printf("here1: %v", updateTransaction)

	transactionJSON, err1 := json.Marshal(updateTransaction)

	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
	}

	req, err2 := http.NewRequest("PUT", transactionServiceURI, bytes.NewBuffer(transactionJSON))

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

	var transactionResDTO ResponseDTO
	err := json.NewDecoder(res.Body).Decode(&transactionResDTO)

	if err != nil {
		responseDto := ResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	c.JSON(200, res.Body)
	//c.JSON(transactionResDTO.Status, transactionResDTO)
}
