package main

import (
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/app"
)

// @title Gin Swagger
// @version 2.0
// @description Order Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Suhas R
// @contact.email suhas7thfeb2000@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7000
// @BasePath /
// @schemes http
func main() {
	app.Start()
}
