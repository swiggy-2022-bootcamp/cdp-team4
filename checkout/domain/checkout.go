package domain

import (
	context "context"

	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/infra/logger"
	"github.com/swiggy-2022-bootcamp/cdp-team4/checkout/protos"
)

var log logrus.Logger = *logger.GetLogger()

func GetShippingCost(ctx context.Context, request *protos.ShippingCostRequest) (*protos.ShippingCostResponse, error) {
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

func GetRewardPoints(ctx context.Context, request *protos.GetRewardPointsRequest) (*protos.GetRewardPointsResponse, error) {
	client, err := infra.GetRewardGrpcClient()
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("get reward gRPC client")
		return nil, err
	}

	response, err := client.GetRewardPoints(ctx, request)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Error("get reward cost gRPC call")
		return nil, err
	}

	return response, nil
}
