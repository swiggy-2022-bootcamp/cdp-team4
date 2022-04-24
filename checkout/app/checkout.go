package app

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/domain"
	protos "github.com/swiggy-2022-bootcamp/cdp-team4/checkout/protos/protoImpl"
)

// Checkout is the client API for Checkout service.
type Checkout struct {
	protos.UnimplementedCheckoutServer
}

// Constructor method to get the checkout stuct object
func NewCheckout() *Checkout {
	return &Checkout{}
}

// gRPC service method which is going to calculate the overview of order
// by getting reward points of user from reward serice, cart details from
// cart service and shipping cost from shipping service.
//
// All the communication between  microservices is happened using gRPC call only
func (ch *Checkout) OrderOverview(
	ctx context.Context,
	rq *protos.OverviewRequest,
) (*protos.OverviewResponse, error) {
	// Get the User Cart details by ID 		[grpc call to Cart service]
	// Get the Shipping details by ID 		[grpc call to Shipping service]
	shippingCost, err := domain.GetShippingCost(ctx, &protos.ShippingCostRequest{City: "chennai"})
	if err != nil {
		log.WithFields(logrus.Fields{"response": shippingCost, "error": err}).Debug("shipping cost gRPC response")
		return nil, err
	}

	// Get the Reward points details by ID 	[grpc call to Reward service]
	rewardPoints, err := domain.GetRewardPoints(ctx, &protos.GetRewardPointsRequest{UserId: "userId"})
	if err != nil {
		log.WithFields(logrus.Fields{"response": rewardPoints, "error": err}).Debug("shipping cost gRPC response")
		return nil, err
	}

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
