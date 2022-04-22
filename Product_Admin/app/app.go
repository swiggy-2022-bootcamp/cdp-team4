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

func setupRouter() *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	ProductAdminRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Product Admin API"
}

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

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	productAdminHandler = ProductAdminHandler{ProductAdminService: domain.NewProductAdminService(dynamoRepository)}

	go startGrpcProductServer(productAdminHandler)

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal(err)
		return
	}

	configureSwaggerDoc()

	PORT := os.Getenv("PORT")
	router := setupRouter()
	router.Run(":" + PORT)
}
