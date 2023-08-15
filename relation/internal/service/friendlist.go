package service

import (
	"context"
	"relation/internal/biz"
	"relation/internal/data/model"

	pb "relation/api/toktik/relation"
)

type FriendListService struct {
	pb.UnimplementedFriendListServer
	uc *biz.FollowUsecase
}

func NewFriendListService(uc *biz.FollowUsecase) *FriendListService {
	return &FriendListService{uc: uc}
}

func (s *FriendListService) FriendListSrv(ctx context.Context, req *pb.FriendListRequest) (*pb.FriendListResponse, error) {
	follow := model.Follow{UserUID: req.UserId}
	list, err := s.uc.GetFollowList(ctx, &follow, biz.QueryTypeFollow)
	var msg string
	if err != nil {
		msg = err.Error()
		return &pb.FriendListResponse{StatusCode: fail, StatusMsg: &msg}, nil
	}

	msg = "查询完成"
	return &pb.FriendListResponse{StatusCode: success, StatusMsg: &msg, UserList: list}, nil
}
