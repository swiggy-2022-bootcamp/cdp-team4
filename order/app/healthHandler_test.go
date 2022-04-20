package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/app"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra"
)

func TestHeathCheck(t *testing.T) {
	dynamoRepo := infra.NewDynamoRepository()
	service := domain.NewOrderService(dynamoRepo)
	orderHandler := app.NewOrderHandler(service)
	router := app.SetupRouter(orderHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Service is running\"}", w.Body.String())
}
