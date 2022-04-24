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
