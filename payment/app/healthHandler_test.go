package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/app"
)

func TestHeathCheck(t *testing.T) {

	router := app.SetupRouter(paymentHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Service is running\"}", w.Body.String())
}
