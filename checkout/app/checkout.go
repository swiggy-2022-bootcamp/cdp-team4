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

func getShippingCost(ctx context.Context, UserID string) (*protos.ShippingCostResponse, error) {
	// Get the User address city by ID 		[grpc call to User service]
	userCity, err := domain.GetUserCity(
		ctx,
		&protos.UserCityRequest{UserID: UserID},
	)
	if err != nil {
		log.WithFields(logrus.Fields{"response": userCity, "error": err}).
			Debug("user gRPC response")
		return nil, err
	}

	// Get the Shipping details by ID 		[grpc call to Shipping service]
	shippingCost, err := domain.GetShippingCost(
		ctx,
		&protos.ShippingCostRequest{City: userCity.City},
	)
	if err != nil {
		log.WithFields(logrus.Fields{"response": shippingCost, "error": err}).
			Debug("shipping cost gRPC response")
		return nil, err
	}

	return shippingCost, err
}

func getItemListAndCartPrice(
	ctx context.Context,
	UserID string,
) ([]*protos.OverviewResponse_Item, int, error) {
	// Get the User Cart details by ID 		[grpc call to Cart service]
	cartDetails, err := domain.GetCartDetails(
		ctx,
		&protos.GetCartByUserIDRequest{UserId: UserID},
	)
	if err != nil {
		log.WithFields(logrus.Fields{"response": cartDetails, "error": err}).
			Debug("shipping cost gRPC response")
		return nil, 0, err
	}

	// make Item list from cart products and find total price
	// by adding price of all the products
	var ItemList []*protos.OverviewResponse_Item
	var TotalItemPrice int
	for _, ele := range cartDetails.Items {
		ItemList = append(ItemList, &protos.OverviewResponse_Item{
			Id:    ele.ProductId,
			Name:  ele.ProductName,
			Price: ele.Price,
			Qty:   ele.Qty,
		})
		TotalItemPrice += int(ele.Price)
	}

	return ItemList, TotalItemPrice, nil
}

// function that calculates the total price by adding cart price and shipping cost
// and subtracting the rewards points if user have
func getTotalOrderPrice(cartPrice, shippingPrice, rewardPoints int) (int, int) {
	totalOrderPrice := cartPrice + shippingPrice

	// 1 amount of cash = 10 points
	rewardPointsToCash := int(float32(rewardPoints) * 0.1)
	tenPercentOfOrderPrice := int(float32(totalOrderPrice) * 0.1)

	rewardConsumed := 0

	// if user have more cash rewards  then order price then
	if tenPercentOfOrderPrice <= rewardPointsToCash {
		rewardPointsToCash -= tenPercentOfOrderPrice
		rewardConsumed = tenPercentOfOrderPrice
		tenPercentOfOrderPrice = 0
	} else {
		tenPercentOfOrderPrice -= rewardPointsToCash
		rewardConsumed = rewardPointsToCash
	}

	return tenPercentOfOrderPrice + int(float32(rewardPoints)*0.1), rewardConsumed
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
	shippingCost, err := getShippingCost(ctx, rq.GetUserID())
	if err != nil {
		log.WithFields(logrus.Fields{"response": shippingCost, "error": err}).
			Debug("shipping cost gRPC response")
		return nil, err
	}

	itemList, totalCartPrice, err := getItemListAndCartPrice(ctx, rq.GetUserID())
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).
			Debug("shipping cost gRPC response")
		return nil, err
	}

	// Get the Reward points details by ID 	  [grpc call to Reward service]
	rewardPoints, err := domain.GetRewardPoints(
		ctx,
		&protos.GetRewardPointsRequest{UserId: rq.GetUserID()},
	)
	if err != nil {
		log.WithFields(logrus.Fields{"response": rewardPoints, "error": err}).
			Debug("shipping cost gRPC response")
		return nil, err
	}

	totalOrderPrice, rewardPointsConsumed := getTotalOrderPrice(
		totalCartPrice,
		int(shippingCost.Cost),
		int(rewardPoints.RewardPoints),
	)

	return &protos.OverviewResponse{
		UserID:               rq.GetUserID(),
		TotalPrice:           int32(totalOrderPrice),
		ShippingPrice:        int32(shippingCost.Cost),
		RewardPointsConsumed: int32(rewardPointsConsumed),
		Item:                 itemList,
	}, nil
}
