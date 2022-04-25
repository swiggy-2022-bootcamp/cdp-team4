package app

import (
	// "context"
	// "fmt"
	// "testing"

	// "github.com/stretchr/testify/assert"
	// pb "github.com/swiggy-2022-bootcamp/cdp-team4/reward/app/protos"
	// "google.golang.org/grpc"
)

// func TestGrpcRewardClient(t *testing.T) {
// 	// Set up connection with the grpc server

// 	conn, err := grpc.Dial("localhost:7010", grpc.WithInsecure())
// 	if err != nil {
// 		fmt.Printf("Error while making connection, %v\n", err)
// 	}

// 	// Create a client instance
// 	client := pb.NewRewardClient(conn)

// 	// Lets invoke the remote function from client on the server
// 	resp, err := client.GetRewardPoints(
// 		context.Background(),
// 		&pb.GetRewardPointsRequest{
// 			UserId: "1",
// 		},
// 	)
// 	assert.Nil(t, err)
// 	assert.Equal(t, uint32(10), resp.RewardPoints)
// }
