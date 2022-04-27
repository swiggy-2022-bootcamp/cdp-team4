package app

import (
	"os"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/infra/logger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/swiggy-2022-bootcamp/cdp-team4/user/infra"	
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"		
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/user/app/protobuf"
)

var log logrus.Logger = *logger.GetLogger()

type Routes struct {
	router *gin.Engine
}

func SetupRouter(userHandler UserHandler) *gin.Engine {
	router := gin.Default()
	
	// health check route
	HealthCheckRouter(router)

	// user route
	UserRouter(router, userHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger User API"
}

func StartHttpServer() {
	err1 := godotenv.Load(".env")
	if err1 != nil {
		logrus.Fatal(err1)
		return
	}
	PORT := os.Getenv("PORT")


	userDynamodbRepository := infra.NewDynamoRepository()

	userHandler := UserHandler{
		UserService: domain.NewUserService(userDynamodbRepository),
		TestMode: false,
	}

	configureSwaggerDoc()
	router := SetupRouter(userHandler)

	router.Run(":" + PORT)
}

func setupServer() *grpc.Server {
	gs := grpc.NewServer()
	cs := NewUserGrpcServer()

	pb.RegisterUserServer(gs, cs)
	log.Debug("gRPC server registered!")

	reflection.Register(gs)
	return gs
}

// Function to start the gRPC server after getting the client
// object from setupServer function and reading the port number
// from .env file
func StartGrpcServer() {
	fmt.Println("grpc starting...")
	err := godotenv.Load(".env")
	// sometime it happens that .env is not present in project directory
	// as it is not pushed on github
	if err != nil {
		log.Fatal(err)
		fmt.Printf("error %v", err)
		return
	}
	PORT := os.Getenv("GRPC_PORT")

	gServer := setupServer()
	l, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
		return
	}

	gServer.Serve(l)
}