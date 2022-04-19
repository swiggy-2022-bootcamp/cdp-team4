package app

import (
	"context"

	pb "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/app/protobuf"
)

type GrpcServer struct {
	pb.UnimplementedShippingServer
}

func (S *GrpcServer) mustEmbedUnimplementedShippingServer() {}

func (S *GrpcServer) GetShippingCost(ctx context.Context, in *pb.ShippingCostRequest) (out *pb.ShippingCostResponse, err1 error) {
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

func NewShippingGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (S *GrpcServer) GetShippingAddress(ctx context.Context, in *pb.ShippingAddressRequest) (out *pb.ShippingAddressResponse, err1 error) {
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

func (S *GrpcServer) AddShippingAddress(ctx context.Context, in *pb.ShippingAddressAddRequest) (out *pb.ShippingAddressAddResponse, err1 error) {
	res, err := shippingHandler.ShippingAddressService.CreateShippingAddress(in.Firstname, in.Lastname, in.City, in.Address1, in.Address2, int(in.Countryid), int(in.Postcode))
	if err != nil {
		return &pb.ShippingAddressAddResponse{}, err.AsMessage().Error()
	}
	return &pb.ShippingAddressAddResponse{
		ShippingAddressID: res,
	}, nil
}

func (S *GrpcServer) DeleteShippingAddress(ctx context.Context, in *pb.ShippingAddressRequest) (out *pb.ShippingAddressDeleteResponse, err1 error) {
	res, err := shippingHandler.ShippingAddressService.DeleteShippingAddressById(in.ShippingAddressID)
	if err != nil {
		return &pb.ShippingAddressDeleteResponse{Confirm: false}, err.AsMessage().Error()
	}
	return &pb.ShippingAddressDeleteResponse{Confirm: res}, nil
}
