package app

import (
	"context"

	pb "github.com/swiggy-2022-bootcamp/cdp-team4/cart/app/protos"
)

type GrpcServer struct {
	pb.UnimplementedCartServer
}

func (server *GrpcServer) GetCartByUserId(ctx context.Context, req *pb.GetCartByUserIDRequest) (*pb.GetCartResponse, error) {
	res, err := cartHandler.CartService.GetCartByUserId(req.UserId)
	//fmt.Println(res)
	if err != nil {
		return nil, err.Error()
	}
	var itemList []*pb.GetCartResponse_Item
	for key,value :=range res.Items{
		var singleItem pb.GetCartResponse_Item
		singleItem.ProductId=key
		singleItem.ProductName=value.Name
		singleItem.Qty=int32(value.Quantity)
		singleItem.Price=int32(value.Cost)
		itemList = append(itemList, &singleItem)
	}
	return &pb.GetCartResponse{
		UserId: res.UserID,
		Items: itemList,
	}, nil
}

func NewCartGrpcServer() *GrpcServer {
	return &GrpcServer{}
}
