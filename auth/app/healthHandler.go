package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {
}

// GetHealthStatus godoc
// @Schemes
// @Description  Returns health status
// @Tags         health
// @Produce      json
// @Success      200  {object}  ResponseDTO
// @Router       /health [get]
func (hh HealthHandler) GetHealthStatus(c *gin.Context) {
	message := "ok"
	responseDto := ResponseDTO{
		Status:  http.StatusOK,
		Message: message,
	}
	c.JSON(responseDto.Status, responseDto)
}
