package app

import (
	"context"
	"fmt"

	pb "github.com/swiggy-2022-bootcamp/cdp-team4/transaction/app/protos"
)

type GrpcServer struct {
	pb.UnimplementedTransactionServer
}

func (server *GrpcServer) GetTransactionPoints(ctx context.Context, req *pb.GetTransactionPointsRequest) (*pb.GetTransactionPointsResponse, error) {
	res, err := transactionHandler.TransactionService.GetTransactionByUserId(req.UserId)
	fmt.Println(res, err)
	if err != nil {
		return nil, err.Error()
	}
	return &pb.GetTransactionPointsResponse{
		TransactionPoints: uint32(res.TransactionPoints),
	}, nil
}

func NewTransactionGrpcServer() *GrpcServer {
	return &GrpcServer{}
}
