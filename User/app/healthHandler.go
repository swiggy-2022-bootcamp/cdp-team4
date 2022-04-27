package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health Check
// @Summary      Health of user service
// @Description  use tocheck whether user service is up and running or not
// @Tags         Health
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /    [get]
func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Service is running"})
	}
}