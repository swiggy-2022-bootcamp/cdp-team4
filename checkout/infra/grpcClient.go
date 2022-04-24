package infra

import (
	"fmt"

	protos "github.com/swiggy-2022-bootcamp/cdp-team4/checkout/protos/protoImpl"
	"google.golang.org/grpc"
)

// Function to get the gRPC client object of shipping service
// Dialing to shipping service without any security as Dialup option
func GetShippingGrpcClient() (protos.ShippingClient, error) {
	conn, err := grpc.Dial("localhost:7012", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("unable to connect with grpc server")
	}

	client := protos.NewShippingClient(conn)
	return client, nil
}

// Function to get the gRPC client object of reward service
func GetRewardGrpcClient() (protos.RewardClient, error) {
	conn, err := grpc.Dial("localhost:7010", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("unable to connect with grpc server")
	}

	client := protos.NewRewardClient(conn)
	return client, nil
}

// Function to get the gRPC client object of cart service
func GetCartGrpcClient() (protos.CartClient, error) {
	conn, err := grpc.Dial("localhost:7006", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("unable to connect with grpc server")
	}

	client := protos.NewCartClient(conn)
	return client, nil
}
