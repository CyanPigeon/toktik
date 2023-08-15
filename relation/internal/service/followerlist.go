package service

import (
	"context"
	"relation/internal/biz"
	"relation/internal/data/model"

	pb "relation/api/toktik/relation"
)

type FollowerListService struct {
	pb.UnimplementedFollowerListServer
	uc *biz.FollowUsecase
}

func NewFollowerListService(uc *biz.FollowUsecase) *FollowerListService {
	return &FollowerListService{uc: uc}
}

func (s *FollowerListService) FollowerListSrv(ctx context.Context, req *pb.FollowerListRequest) (*pb.FollowerListResponse, error) {
	follow := model.Follow{UserUID: req.UserId}
	list, err := s.uc.GetFollowList(ctx, &follow, biz.QueryTypeFollower)
	var msg string
	if err != nil {
		msg = err.Error()
		return &pb.FollowerListResponse{StatusCode: fail, StatusMsg: &msg}, nil
	}

	msg = "查询完成"
	return &pb.FollowerListResponse{StatusCode: success, StatusMsg: &msg, UserList: list}, nil

}
