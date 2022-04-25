package app

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/cart/app/protos"
	"google.golang.org/grpc"
)

func TestGrpcCartClient(t *testing.T) {
	// Set up connection with the grpc server

	conn, err := grpc.Dial("localhost:7006", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	client := pb.NewCartClient(conn)

	// Lets invoke the remote function from client on the server
	resp, err := client.GetCartByUserId(
		context.Background(),
		&pb.GetCartByUserIDRequest{
			UserId: "P",
		},
	)
	fmt.Println(resp)
	assert.Nil(t, err)
	//assert.Equal(t, "1", resp.UserId)
}
