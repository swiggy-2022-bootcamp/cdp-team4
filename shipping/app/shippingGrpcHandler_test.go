package app

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/app/protobuf"
	"google.golang.org/grpc"
)

func TestGrpcAddressClient(t *testing.T) {
	// Set up connection with the grpc server

	conn, err := grpc.Dial("localhost:"+os.Getenv("GRPC_SHIPPING_PORT"), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	c := pb.NewShippingClient(conn)

	// Lets invoke the remote function from client on the server
	resp, err := c.GetShippingAddress(
		context.Background(),
		&pb.ShippingAddressRequest{
			ShippingAddressID: "73e1285b-58d0-40ff-bf38-beda77712b5d",
		},
	)
	assert.Nil(t, err)
	assert.Equal(t, resp.Firstname, "Naveen")
	assert.Equal(t, resp.City, "Banglore")

	resp2, err := c.DeleteShippingAddress(context.Background(),
		&pb.ShippingAddressRequest{
			ShippingAddressID: "8c87aaa5-81b0-43a7-88e8-31b2026ea3b3",
		},
	)
	assert.Nil(t, err)
	assert.Equal(t, resp2.Confirm, true)
}

func TestGrpcCostClient(t *testing.T) {
	// Set up connection with the grpc server

	conn, err := grpc.Dial("localhost:7776", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	c := pb.NewShippingClient(conn)
	// Lets invoke the remote function from client on the server
	resp, err := c.GetShippingCost(
		context.Background(),
		&pb.ShippingCostRequest{
			City: "Chennai",
		},
	)
	assert.Nil(t, err)
	assert.Equal(t, resp.City, "Chennai")
	assert.Equal(t, int(resp.Cost), 199)
}
