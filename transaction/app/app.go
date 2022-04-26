package app

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/transaction/app/protos"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/infra/gokafka"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/infra/logger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var log logrus.Logger = *logger.GetLogger()
var transactionHandler TransactionHandler

func setupRouter() *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	transactionRouter(router, transactionHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Transaction API"
}

func startGrpcTransactionServer(th TransactionHandler) {
	grpcServer := grpc.NewServer()
	transactionServer := NewTransactionGrpcServer()
	pb.RegisterTransactionServer(grpcServer, transactionServer)
	fmt.Println("register server")
	reflection.Register(grpcServer)
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error(err)
		return
	}
	GRPC_REWARD_PORT := os.Getenv("GRPC_PORT")
	fmt.Println(GRPC_REWARD_PORT)
	l, err := net.Listen("tcp", ":"+GRPC_REWARD_PORT)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print("grpc server started")
	grpcServer.Serve(l)
}

func startKafkaConsumer(repo infra.TransactionDynamoRepository) {
	ctx := context.Background()
	go gokafka.UpdateTransactionPoints(ctx, "payment", repo)
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	transactionHandler = TransactionHandler{TransactionService: domain.NewTransactionService(dynamoRepository)}

	go startGrpcTransactionServer(transactionHandler)
	startKafkaConsumer(dynamoRepository)
	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("PORT")

	configureSwaggerDoc()
	router := setupRouter()

	router.Run(":" + PORT)
	// if err := r.Run(":3000"); err != nil {
	// 	log.Fatal(err)
	// }
}
