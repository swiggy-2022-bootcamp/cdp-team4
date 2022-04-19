package app

import (
	"context"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/order/app/protobuf"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra/gokafka"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var orderHandler OrderHandler
var log logrus.Logger = *logger.GetLogger()

func setupRouter() *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	OrderRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Order API"
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	orderHandler = OrderHandler{OrderService: domain.NewOrderService(dynamoRepository)}
	startKafkaConsumer(dynamoRepository)
	// grpcserver for testing
	//go testGrpcServer()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("ORDER_SERVICE_PORT")
	router := setupRouter()
	configureSwaggerDoc()

	router.Run(":" + PORT)
}

func startKafkaConsumer(repo infra.OrderDynamoRepository) {
	ctx := context.Background()
	go gokafka.StatusConsumer(ctx, "payment", repo)
}

func testGrpcServer() {
	gs := grpc.NewServer()
	ss := NewCheckout()
	pb.RegisterCheckoutServer(gs, ss)
	reflection.Register(gs)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	l, err := net.Listen("tcp", ":"+"7899")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	gs.Serve(l)
}
