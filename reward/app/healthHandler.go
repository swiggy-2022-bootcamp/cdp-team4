package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health Check
// @Summary      Health of Reward service
// @Description  Endpoint to check the health of Reward Microserice
// @Tags         Health
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /    [get]
func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("Log working")
		c.JSON(http.StatusOK, gin.H{"message": "Service is running"})
	}
}