package app

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/app/protobuf"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/shipping/infra/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log logrus.Logger = *logger.GetLogger()
var shippingHandler ShippingHandler

func setupRouter() *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	ShippingRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Order API"
}

func Start() {
	dynamoRepository := infra.NewDynamoShippingAddressRepository()
	dynamoRepository1 := infra.NewShippingCostDynamoRepository()
	shippingHandler = ShippingHandler{
		ShippingAddressService: domain.NewShippingAddressService(dynamoRepository),
		ShippingCostService:    domain.NewShippingCostService(dynamoRepository1),
	}
	log.WithFields(logrus.Fields{"message": "message", "status": http.StatusBadRequest}).Error("Error Check")
	go startGrpcCostServer(shippingHandler)
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(err)
		return
	}
	PORT := os.Getenv("SHIPPING_SERVICE_PORT")
	router := setupRouter()
	configureSwaggerDoc()

	router.Run(":" + PORT)
}

func startGrpcCostServer(sh ShippingHandler) {
	gs := grpc.NewServer()
	ss := NewShippingGrpcServer()
	pb.RegisterShippingServer(gs, ss)
	reflection.Register(gs)
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(err)
		return
	}
	GRPC_COST_PORT := os.Getenv("GRPC_PORT")
	l, err := net.Listen("tcp", ":"+GRPC_COST_PORT)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	gs.Serve(l)
}

// func startGrpcAddressServer(sh ShippingHandler) {
// 	gs := grpc.NewServer()
// 	ss := NewShippingGrpcAddressServer()
// 	pb.RegisterShippingAddressServer(gs, ss)
// 	reflection.Register(gs)
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Print(err)
// 		return
// 	}
// 	GRPC_ADDRESSS_PORT := os.Getenv("GRPC_ADDRESSS_PORT")
// 	l, err := net.Listen("tcp", ":"+GRPC_ADDRESSS_PORT)
// 	if err != nil {
// 		fmt.Print(err)
// 		os.Exit(1)
// 	}

// 	gs.Serve(l)
// }
