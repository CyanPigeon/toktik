package service

import (
	"context"
	"relation/internal/biz"
	"relation/internal/data/model"

	pb "relation/api/toktik/relation"
)

type FollowListService struct {
	pb.UnimplementedFollowListServer
	uc *biz.FollowUsecase
}

func NewFollowListService(uc *biz.FollowUsecase) *FollowListService {
	return &FollowListService{uc: uc}
}

func (s *FollowListService) FollowListSrv(ctx context.Context, req *pb.FollowListRequest) (*pb.FollowListResponse, error) {
	follow := model.Follow{UserUID: req.UserId}
	list, err := s.uc.GetFollowList(ctx, &follow, biz.QueryTypeFollow)
	var msg string
	if err != nil {
		msg = err.Error()
		return &pb.FollowListResponse{StatusCode: fail, StatusMsg: &msg}, nil
	}

	msg = "查询完成"
	return &pb.FollowListResponse{StatusCode: success, StatusMsg: &msg, UserList: list}, nil
}
