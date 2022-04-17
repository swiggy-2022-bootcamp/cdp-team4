package main

import (
	"fmt"
	"net"
	"os"

	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// @title Checkout API
// @version 1.0
// @description Checkout Service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	// app.Start()
	gs := grpc.NewServer()
	cs := domain.NewCheckout()

	domain.RegisterCheckoutServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	gs.Serve(l)
}
