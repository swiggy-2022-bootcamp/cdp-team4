package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{"userId": c.Param("userId")})
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
