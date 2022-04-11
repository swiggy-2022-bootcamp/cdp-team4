package main

import (
	"github.com/swiggy-2022-bootcamp/cdp-team4/Cart/app"
)

// @title Gin Swagger Example API
// @version 2.0
// @description Order Service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8005
// @BasePath /
// @schemes http
func main() {
	app.Start()
}