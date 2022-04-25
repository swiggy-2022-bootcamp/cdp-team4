package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/app"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/infra"
)

func TestHeathCheck(t *testing.T) {

}

func TestPingRoute(t *testing.T) {
	dynamoRepository := infra.NewDynamoRepository()
	productFrontStoreService := domain.NewProductFrontStoreService(dynamoRepository)
	productFrontStoreHandler := app.NewProductFrontStoreHandler(productFrontStoreService)

	router := app.SetupRouter(productFrontStoreHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Service is running\"}", w.Body.String())
}
