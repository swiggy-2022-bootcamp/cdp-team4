package app

import (
	"github.com/gin-gonic/gin"
)

type productHandler struct {
}

func (ph productHandler) GetAllProducts(c *gin.Context) {
	pah := productAdminHandler{}
	pfsh := productFrontStoreHandler{}

	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		pfsh.GetAllProducts(c)
		return
	} else {
		pah.GetAllProducts(c)
		return
	}
}

func (ph productHandler) GetProductByID(c *gin.Context) {
	pah := productAdminHandler{}
	pfsh := productFrontStoreHandler{}

	isAdmin := c.Request.Header.Get("admin")
	if isAdmin == "false" {
		pfsh.GetProductByID(c)
		return
	} else {
		pah.GetProductByID(c)
		return
	}
}
