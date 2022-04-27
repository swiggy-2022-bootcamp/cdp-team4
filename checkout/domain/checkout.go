package domain

import (
	context "context"

	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/infra/logger"
	protos "github.com/swiggy-2022-bootcamp/cdp-team4/checkout/protos/protoImpl"
)

var log logrus.Logger = *logger.GetLogger()

// function that makes gprc call to shipping service and gets back the shipping
// cost on the basis of city as given input
func GetShippingCost(
	ctx context.Context,
	request *protos.ShippingCostRequest,
) (*protos.ShippingCostResponse, error) {
	client, err := infra.GetShippingGrpcClient()
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("get shipping gRPC client")
		return nil, err
	}

	response, err := client.GetShippingCost(ctx, request)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("get shipping cost gRPC call")
		return nil, err
	}
	return response, nil
}

// function that makes gprc call to reward service and gets back the reward
// points for particular user on the basis of user id as given input
func GetRewardPoints(
	ctx context.Context,
	request *protos.GetRewardPointsRequest,
) (*protos.GetRewardPointsResponse, error) {
	client, err := infra.GetRewardGrpcClient()
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("get reward gRPC client")
		return nil, err
	}

	response, err := client.GetRewardPoints(ctx, request)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("get reward points gRPC call")
		return nil, err
	}

	return response, nil
}

// function that makes gprc call to cart service and gets back the cart details
//  for particular user on the basis of user id as given input
func GetCartDetails(
	ctx context.Context,
	request *protos.GetCartByUserIDRequest,
) (*protos.GetCartResponse, error) {
	client, err := infra.GetCartGrpcClient()
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("get cart details gRPC client")
		return nil, err
	}

	response, err := client.GetCartByUserId(ctx, request)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("get cart details gRPC call")
		return nil, err
	}

	return response, nil
}

// function that makes gprc call to user service and gets back the city
//  for particular user on the basis of user id as given input
func GetUserCity(
	ctx context.Context,
	request *protos.UserCityRequest,
) (*protos.UserCityResponse, error) {
	client, err := infra.GetUserGrpcClient()
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("get cart details gRPC client")
		return nil, err
	}

	response, err := client.GetUserCity(ctx, request)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("get cart details gRPC call")
		return nil, err
	}

	return response, nil
}

// function that gets city by grpc call to user service and then gets the
// shipping cost value
func GetShippingCostValue(ctx context.Context, UserID string) (*protos.ShippingCostResponse, error) {
	// Get the User address city by ID 		[grpc call to User service]
	userCity, err := GetUserCity(
		ctx,
		&protos.UserCityRequest{UserID: UserID},
	)
	if err != nil {
		log.WithFields(logrus.Fields{"response": userCity, "error": err}).
			Debug("user gRPC response")
		return nil, err
	}

	// Get the Shipping details by ID 		[grpc call to Shipping service]
	shippingCost, err := GetShippingCost(
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

// function that returns the list of the items present in the cart along with
// the total price of the cart
func GetItemListAndCartPrice(
	ctx context.Context,
	UserID string,
) ([]*protos.OverviewResponse_Item, int, error) {
	// Get the User Cart details by ID 		[grpc call to Cart service]
	cartDetails, err := GetCartDetails(
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
// rewards are only applicable on 10% price of total order price that is sum of
// shipping price and cart price.
func GetTotalOrderPrice(cartPrice, shippingPrice, rewardPoints int) (int, int) {
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
