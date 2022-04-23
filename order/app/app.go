package app

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra/gokafka"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra/logger"
)

var log logrus.Logger = *logger.GetLogger()

func SetupRouter(orderHandler OrderHandler) *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	OrderRouter(router, orderHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Order API"
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	dynamoRepositoryOrderOverview := infra.NewDynomoOrderOverviewRepository()
	orderHandler := NewOrderHandler(domain.NewOrderService(dynamoRepository), domain.NewOrderOverviewService(dynamoRepositoryOrderOverview))
	startKafkaConsumer(dynamoRepository, dynamoRepositoryOrderOverview)
	// grpcserver for testing
	//go testGrpcServer()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err, "start")
		return
	}
	PORT := os.Getenv("ORDER_SERVICE_PORT")
	router := SetupRouter(orderHandler)
	configureSwaggerDoc()

	router.Run(":" + PORT)
}

func startKafkaConsumer(repo infra.OrderDynamoRepository, repo1 infra.OrderDynamoRepository) {
	ctx := context.Background()
	go gokafka.StatusConsumer(ctx, "payment", repo, repo1)
}

// func testGrpcServer() {
// 	gs := grpc.NewServer()
// 	ss := NewCheckout()
// 	pb.RegisterCheckoutServer(gs, ss)
// 	reflection.Register(gs)
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	l, err := net.Listen("tcp", ":"+"7899")
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(1)
// 	}

// 	gs.Serve(l)
// }
