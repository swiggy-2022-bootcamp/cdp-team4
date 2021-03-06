package app

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/reward/app/protos"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/infra/logger"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/infra/gokafka"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log logrus.Logger = *logger.GetLogger()
var rewardHandler RewardHandler

func SetupRouter(rewardHandler RewardHandler) *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	rewardRouter(router, rewardHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Reward API"
}

func startGrpcRewardServer(rh RewardHandler) {
	grpcServer := grpc.NewServer()
	rewardServer := NewRewardGrpcServer()
	pb.RegisterRewardServer(grpcServer, rewardServer)
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

func startKafkaConsumer(repo infra.RewardDynamoRepository) {
	ctx := context.Background()
	go gokafka.UpdateRewardPoints(ctx, "payment", repo)
}

func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	rewardHandler = RewardHandler{RewardService: domain.NewRewardService(dynamoRepository)}

	go startGrpcRewardServer(rewardHandler)
	startKafkaConsumer(dynamoRepository)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("PORT")

	configureSwaggerDoc()
	router := SetupRouter(rewardHandler)

	router.Run(":" + PORT)
	// if err := r.Run(":3000"); err != nil {
	// 	log.Fatal(err)
	// }
}
