package app

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	pb "github.com/swiggy-2022-bootcamp/cdp-team4/transaction/app/protos"
// 	"google.golang.org/grpc"
// )

// func TestGrpcTransactionClient(t *testing.T) {
// 	// Set up connection with the grpc server

// 	conn, err := grpc.Dial("localhost:7007", grpc.WithInsecure())
// 	if err != nil {
// 		fmt.Printf("Error while making connection, %v\n", err)
// 	}

// 	// Create a client instance
// 	client := pb.NewTransactionClient(conn)

// 	// Lets invoke the remote function from client on the server
// 	resp, err := client.GetTransactionPoints(
// 		context.Background(),
// 		&pb.GetTransactionPointsRequest{
// 			UserId: "b",
// 		},
// 	)
// 	fmt.Println(resp)
// 	assert.Nil(t, err)
// 	assert.Equal(t, uint32(165), resp.TransactionPoints)
// }
