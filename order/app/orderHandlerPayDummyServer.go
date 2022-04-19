package app

import (
	context "context"

	pb "github.com/swiggy-2022-bootcamp/cdp-team4/order/app/protobuf"
)

type Checkout struct {
	pb.UnimplementedCheckoutServer
}

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (ch *Checkout) mustEmbedUnimplementedCheckoutServer() {}

func (ch *Checkout) OrderOverview(
	ctx context.Context,
	rq *pb.OverviewRequest,
) (*pb.OverviewResponse, error) {
	// Get the User Cart details by ID 		[grpc call to Cart service]
	// Get the Reward points details by ID 	[grpc call to Reward service]
	// Get the Shipping details by ID 		[grpc call to Shipping service]

	return &pb.OverviewResponse{
		UserID:               rq.GetUserID(),
		TotalPrice:           15,
		ShippingPrice:        5,
		RewardPointsConsumed: 2,
		Item: []*pb.OverviewResponse_Item{
			{Name: "item1", Id: "item1", Price: 5, Qty: 3},
		},
	}, nil
}
