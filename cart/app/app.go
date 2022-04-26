package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/cart/app/protos"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/infra/gokafka"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var cartHandler CartHandler

func setupRouter(cartHandler CartHandler) *gin.Engine {
	router := gin.Default()
	// health check route
	HealthCheckRouter(router)
	CartRouter(router, cartHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func startGrpcCartServer(ch CartHandler) {
	grpcServer := grpc.NewServer()
	cartServer := NewCartGrpcServer()
	pb.RegisterCartServer(grpcServer, cartServer)
	fmt.Println("register server")
	reflection.Register(grpcServer)
	err := godotenv.Load(".env")
	if err != nil {
		// logrus.Error(err)
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

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Payment API"
}

func startKafkaConsumer(repo infra.CartDynamoRepository) {
	ctx := context.Background()
	go gokafka.EmptyCartAfterOrderProccessed(ctx, "checkout", repo)
}


func Start() {
	dynamoRepository := infra.NewDynamoRepository()
	cartHandler = CartHandler{CartService: domain.NewCartService(dynamoRepository)}

	go startGrpcCartServer(cartHandler)
	startKafkaConsumer(dynamoRepository)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	PORT := os.Getenv("PORT")

	configureSwaggerDoc()
	router := setupRouter(cartHandler)

	if err := router.Run(":" + PORT); err != nil {
		log.Fatal(err)
	}
}
