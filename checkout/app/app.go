package app

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/infra/logger"
	// gin-swagger middleware
)

var log logrus.Logger = *logger.GetLogger()

func setupServer() *grpc.Server {
	gs := grpc.NewServer()
	cs := domain.NewCheckout()

	domain.RegisterCheckoutServer(gs, cs)
	log.Debug("gRPC server registered!")

	reflection.Register(gs)
	return gs
}

// func configureSwaggerDoc() {
// 	docs.SwaggerInfo.Title = "Swagger Checkout API"
// }

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("PORT")

	gServer := setupServer()
	l, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	gServer.Serve(l)
}
