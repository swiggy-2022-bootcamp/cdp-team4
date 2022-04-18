package domain

import (
	context "context"
)

type Checkout struct {
}

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (ch *Checkout) OrderOverview(
	ctx context.Context,
	rq *OverviewRequest,
) (*OverviewResponse, error) {
	// Get the User Cart details by ID 		[grpc call to Cart service]
	// Get the Reward points details by ID 	[grpc call to Reward service]
	// Get the Shipping details by ID 		[grpc call to Shipping service]

	return &OverviewResponse{
		UserID:               rq.GetUserID(),
		TotalPrice:           15,
		ShippingPrice:        5,
		RewardPointsConsumed: 2,
		Item: []*OverviewResponse_Item{
			{Name: "item1", Id: "item1", Price: 5, Qty: "3"},
		},
	}, nil
}

func (ch *Checkout) mustEmbedUnimplementedCheckoutServer() {}
