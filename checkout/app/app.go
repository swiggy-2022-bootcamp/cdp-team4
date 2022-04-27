package app

import (
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/infra/logger"
	protos "github.com/swiggy-2022-bootcamp/cdp-team4/checkout/protos/protoImpl"
)

var log logrus.Logger = *logger.GetLogger()

// Function used to get the new gRPC server object
// after registering checkout server implemented in domain
func setupServer() *grpc.Server {
	gs := grpc.NewServer()
	cs := NewCheckout()

	protos.RegisterCheckoutServer(gs, cs)
	log.Debug("gRPC server registered!")

	reflection.Register(gs)
	return gs
}

// Function to start the gRPC server after getting the client
// object from setupServer function and reading the port number
// from .env file
func Start() {
	err := godotenv.Load(".env")
	// sometime it happens that .env is not present in project directory
	// as it is not pushed on github
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("PORT")

	gServer := setupServer()
	l, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
		return
	}

	gServer.Serve(l)
}
