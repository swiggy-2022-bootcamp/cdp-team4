package app_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckRoute(t *testing.T) {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")

	healthHandler := app.HealthHandler{}

	v1.GET("/health", healthHandler.GetHealthStatus)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"status\":200,\"message\":\"ok\"}", w.Body.String())

}
