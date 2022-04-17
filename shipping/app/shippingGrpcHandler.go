package app

import (
	"context"

	pb "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/app/protobuf"
)

// type GrpcCostServer struct {
// 	pb.UnimplementedShippingCostServer
// }

// func NewShippingGrpcCostServer() *GrpcCostServer {
// 	return &GrpcCostServer{}
// }

func (S *GrpcAddressServer) mustEmbedUnimplementedShippingCostServer() {}

func (S *GrpcAddressServer) GetShippingCost(ctx context.Context, in *pb.ShippingCostRequest) (out *pb.ShippingCostResponse, err1 error) {
	city := in.City
	res, err := shippingHandler.ShippingCostService.GetShippingCostByCity(city)
	if err != nil {
		return nil, err.AsMessage().Error()
	}
	return &pb.ShippingCostResponse{
		City: res.City,
		Cost: uint32(res.ShippingCost),
	}, nil
}

type GrpcAddressServer struct {
	pb.UnimplementedShippingServer
}

func NewShippingGrpcAddressServer() *GrpcAddressServer {
	return &GrpcAddressServer{}
}

func (S1 *GrpcAddressServer) mustEmbedUnimplementedShippingAddressServer() {}

func (S1 *GrpcAddressServer) GetShippingAddress(ctx context.Context, in *pb.ShippingAddressRequest) (out *pb.ShippingAddressResponse, err1 error) {
	id := in.ShippingAddressID
	res, err := shippingHandler.ShippingAddressService.GetShippingAddressById(id)
	if err != nil {
		return &pb.ShippingAddressResponse{}, err.AsMessage().Error()
	}
	return &pb.ShippingAddressResponse{
		Firstname: res.FirstName,
		Lastname:  res.LastName,
		City:      res.City,
		Address1:  res.Address1,
		Address2:  res.Address2,
		Countryid: uint32(res.CountryID),
		Postcode:  uint32(res.PostCode),
	}, nil
}
