package app_test

import (
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/app"
	"testing"
)

func TestShouldStartApp(t *testing.T) {
	app.Start(true)
}
