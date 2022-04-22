package infra

import (
	"fmt"

	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/protos"
	"google.golang.org/grpc"
)

func GetShippingGrpcClient() (protos.ShippingClient, error) {
	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("unable to connect with grpc server")
	}
	defer conn.Close()

	client := protos.NewShippingClient(conn)

	return client, nil
}
