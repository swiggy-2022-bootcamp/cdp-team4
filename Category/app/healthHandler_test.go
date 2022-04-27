package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/app"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Category/infra"
)

func TestHeathCheck(t *testing.T) {

}

func TestPingRoute(t *testing.T) {
	dynamoRepository := infra.NewDynamoRepository()
	categoryService := domain.NewCategoryService(dynamoRepository)
	categoryHandler := app.NewCategoryHandler(categoryService)

	router := app.SetupRouter(categoryHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Service is running\"}", w.Body.String())
}
