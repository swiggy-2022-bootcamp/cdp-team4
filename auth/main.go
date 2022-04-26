package main

import "github.com/swiggy-2022-bootcamp/cdp-team4/auth/app"

// @title        Auth Microservice API
// @version      1.0
// @description  This is a authentication/authorization service.

// @contact.name   Murtaza Sadriwala
// @contact.email  murtaza896@gmail.com

// @host      localhost:8001
// @BasePath  /api/v1
func main() {
	app.Start(false)
}
