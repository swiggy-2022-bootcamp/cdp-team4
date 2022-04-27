package app

import (
	"fmt"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/app/protobuf"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/infra/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var productAdminHandler ProductAdminHandler
var log logrus.Logger = *logger.GetLogger()

// Function used to get the new http gin engine object
// after registering all the routers
func SetupRouter(productAdminHandler ProductAdminHandler) *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	ProductAdminRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func ConfigureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Product Admin API"
}

// Function to start an asynchronous gRPC server such that the inter service
// calls can be made and reading the port number
// from .env file
//
// it also intiliases the all repositories, services and handlers present
// in the this micro-service.
func startGrpcProductServer(pah ProductAdminHandler) {
	grpcServer := grpc.NewServer()
	productServer := NewProductGrpcServer()
	pb.RegisterProductServer(grpcServer, productServer)
	fmt.Println("register server")
	reflection.Register(grpcServer)
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error(err)
		return
	}
	GRPC_PRODUCT_PORT := os.Getenv("GRPC_PORT")
	fmt.Println(GRPC_PRODUCT_PORT)
	l, err := net.Listen("tcp", ":"+GRPC_PRODUCT_PORT)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print("grpc server started")
	grpcServer.Serve(l)

}

// Function to start the http server after getting the client
// object from setupServer function and reading the port number
// from .env file
//
// it also intiliases the all repositories, services and handlers present
// in the this micro-service.
//
// Also configures the Swagger UI:
// http//localhost:8004/swagger/index.html
func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	productAdminService := domain.NewProductAdminService(dynamoRepository)
	productAdminHandler = NewProductAdminHandler(productAdminService)

	go startGrpcProductServer(productAdminHandler)

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal(err)
		return
	}

	ConfigureSwaggerDoc()

	PORT := os.Getenv("PORT")
	router := SetupRouter(productAdminHandler)
	router.Run(":" + PORT)
}
