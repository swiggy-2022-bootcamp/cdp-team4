package app

import (
	"context"

	pb "github.com/swiggy-2022-bootcamp/cdp-team4/reward/app/protos"
)

type GrpcServer struct {
	pb.UnimplementedRewardServer
}

func (server *GrpcServer) GetRewardPoints(ctx context.Context, req *pb.GetRewardPointsRequest) (*pb.GetRewardPointsResponse, error) {
	res, err := rewardHandler.RewardService.GetRewardByUserId(req.UserId)
	if err != nil {
		return nil, err.Error()
	}
	return &pb.GetRewardPointsResponse{
		RewardPoints: uint32(res.RewardPoints),
	}, nil
}

func NewRewardGrpcServer() *GrpcServer {
	return &GrpcServer{}
}
