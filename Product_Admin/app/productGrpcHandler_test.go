package app

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/app/protobuf"
	"google.golang.org/grpc"
)

func TestGrpcProductClient(t *testing.T) {
	// Set up connection with the grpc server

	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	client := pb.NewProductClient(conn)

	// Lets invoke the remote function from client on the server
	resp, err := client.GetProductAvailability(
		context.Background(),
		&pb.ProductAvailabilityGetRequest{
			ProductID:      "625cfa5833b5149cf2fbbc23",
			QuantityNeeded: 2,
		},
	)
	assert.Nil(t, err)
	assert.Equal(t, resp.IsAvailable, true)
}
