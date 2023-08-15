package service

import (
	"context"
	pb "relation/api/toktik/relation"
	"relation/internal/biz"
)

const (
	success = iota
	fail
)

type FollowActionService struct {
	pb.UnimplementedFollowActionServer
	uc *biz.FollowUsecase
}

func NewFollowActionService(uc *biz.FollowUsecase) *FollowActionService {
	return &FollowActionService{uc: uc}
}

func (s *FollowActionService) FollowActionSrv(ctx context.Context, req *pb.FollowActionRequest) (*pb.FollowActionResponse, error) {

	var msg string

	if req.ActionType == 1 {

		err := s.uc.Follow(ctx, req)

		if err != nil {
			msg = err.Error()
			return &pb.FollowActionResponse{StatusCode: fail, StatusMsg: &msg}, nil
		} else {
			msg = "关注成功"
			return &pb.FollowActionResponse{StatusCode: success, StatusMsg: &msg}, nil
		}

	} else if req.ActionType == 2 {

		err := s.uc.UnFollow(ctx, req)

		if err != nil {
			msg = err.Error()
			return &pb.FollowActionResponse{StatusCode: fail, StatusMsg: &msg}, nil
		} else {
			msg = "取关成功"
			return &pb.FollowActionResponse{StatusCode: success, StatusMsg: &msg}, nil
		}

	}

	msg = "非法请求"
	return &pb.FollowActionResponse{StatusCode: fail, StatusMsg: &msg}, nil

}
