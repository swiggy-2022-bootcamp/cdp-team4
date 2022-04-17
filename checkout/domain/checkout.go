package domain

import (
	context "context"
	"fmt"
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
	fmt.Print("request received!")
	return &OverviewResponse{
		UserID:               "abs",
		TotalPrice:           15,
		ShippingPrice:        5,
		RewardPointsConsumed: 2,
		Item: []*OverviewResponse_Item{
			{Name: "item1", Id: "item1", Price: 5, Qty: "3"},
		},
	}, nil
}

func (ch *Checkout) mustEmbedUnimplementedCheckoutServer() {

}
