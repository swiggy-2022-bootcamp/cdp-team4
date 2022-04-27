package app

import (
	"context"
	"fmt"

	pb "github.com/swiggy-2022-bootcamp/cdp-team4/user/app/protobuf"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/infra"	
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"	
	"google.golang.org/grpc"	
)

type GrpcServer struct {
	pb.UnimplementedUserServer
}

func (S *GrpcServer) mustEmbedUnimplementedUserServer() {}

func (S *GrpcServer) GetUserCity(ctx context.Context, in *pb.UserCityRequest) (*pb.UserCityResponse, error) {
	UserID := in.UserID

	fmt.Println("user id: ", UserID)

	userDynamodbRepository := infra.NewDynamoRepository()
	userService := domain.NewUserService(userDynamodbRepository)
	
	res, err := userService.GetUserById(UserID)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(SHIPPING_GRPC_URI, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// Create a client instance
	c := pb.NewShippingClient(conn)

	// Lets invoke the remote function from client on the server
	resp, err1 := c.GetShippingAddress(
		context.Background(),
		&pb.ShippingAddressRequest{
			ShippingAddressID: res.AddressID,
		},
	)
	
	if err1 != nil {
		return nil, nil
	}

	return &pb.UserCityResponse{
		City: resp.City,
	}, nil
}

func NewUserGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

