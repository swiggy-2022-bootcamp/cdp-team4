package domain

import (
	context "context"
	"fmt"

	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/protos"
	"google.golang.org/grpc"
)

type Checkout struct {
	protos.UnimplementedCheckoutServer
}

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (ch *Checkout) OrderOverview(
	ctx context.Context,
	rq *protos.OverviewRequest,
) (*protos.OverviewResponse, error) {
	// Get the User Cart details by ID 		[grpc call to Cart service]
	// Get the Reward points details by ID 	[grpc call to Reward service]
	// Get the Shipping details by ID 		[grpc call to Shipping service]
	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("unable to connect with grpc server")
	}
	defer conn.Close()

	client := protos.NewShippingClient(conn)

	response, err := client.GetShippingCost(ctx, &protos.ShippingCostRequest{City: "Chennai"})
	if err != nil {
		return nil, err
	}

	fmt.Print(response, err)
	return &protos.OverviewResponse{
		UserID:               rq.GetUserID(),
		TotalPrice:           15,
		ShippingPrice:        5,
		RewardPointsConsumed: 2,
		Item: []*protos.OverviewResponse_Item{
			{Name: "item1", Id: "item1", Price: 5, Qty: "3"},
		},
	}, nil
}
