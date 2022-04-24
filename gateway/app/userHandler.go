package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	// "log"
	// "fmt"
	"encoding/json"
	"context"
	"time"
)

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

}
