package app

import (
	"context"

	pb "github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/app/protobuf"
)

type GrpcServer struct {
	pb.UnimplementedProductServer
}

func (server *GrpcServer) GetProductAvailability(ctx context.Context, req *pb.ProductAvailabilityGetRequest) (*pb.ProductAvailabilityGetResponse, error) {
	res, err := productAdminHandler.ProductAdminService.GetProductAvailability(req.ProductID, int64(req.QuantityNeeded))
	if err != nil {
		return nil, err
	}
	return &pb.ProductAvailabilityGetResponse{
		IsAvailable: res,
	}, nil
}

func NewProductGrpcServer() *GrpcServer {
	return &GrpcServer{}
}
