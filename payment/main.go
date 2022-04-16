package main

import "github.com/swiggy-2022-bootcamp/cdp-team4/payment/app"

// @title Payment API
// @version 1.0
// @description Payment Service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	app.Start(false /*testMode*/)
}
