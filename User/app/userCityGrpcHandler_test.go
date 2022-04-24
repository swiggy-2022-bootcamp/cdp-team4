package app

// import (
// 	pb "github.com/swiggy-2022-bootcamp/cdp-team4/user/app/protobuf"
// 	"fmt"
// 	"google.golang.org/grpc"
// 	"context"
// 	"testing"
// )


// func TestGrpcAddressClient(t *testing.T) {
// 	// Set up connection with the grpc server

// 	conn, err := grpc.Dial("localhost:7002", grpc.WithInsecure())
// 	if err != nil {
// 		fmt.Printf("Error while making connection, %v\n", err)
// 	}

// 	// Create a client instance
// 	c := pb.NewUserClient(conn)

// 	// Lets invoke the remote function from client on the server
// 	resp, err := c.GetUserCity(
// 		context.Background(),
// 		&pb.UserCityRequest{
// 			UserID: "626525f965a27e0d14ff3257",
// 		},
// 	)

// 	t.Logf("response is:  %s\n", resp)
	
// }
