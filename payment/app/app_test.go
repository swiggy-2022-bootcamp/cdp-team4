package app_test

import (
	"testing"

	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/app"
)

func TestConfigureSwaggerDoc(t *testing.T) {
	app.ConfigureSwaggerDoc()
}

func TestStart(t *testing.T) {
	app.Start(true /*testMode*/)
}
