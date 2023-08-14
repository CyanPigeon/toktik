package service

import (
	"context"

	pb "favorite/api/toktik/favorite"
	favService "favorite/internal/biz/favorite"
)

type LikeActionService struct {
	pb.UnimplementedLikeActionServer
	favService *favService.FavoriteServiceBiz
}

func NewLikeActionService(favService *favService.FavoriteServiceBiz) *LikeActionService {
	return &LikeActionService{
		favService: favService,
	}
}

func (s *LikeActionService) LikeActionSrv(ctx context.Context, req *pb.LikeActionRequest) (*pb.LikeActionResponse, error) {
	// 权限pass
	retBool, mess, _ := s.favService.LikeActionSrv(ctx, req)
	if retBool {
		var ACTION_SUCCESS = ""
		if req.ActionType == 1 {
			ACTION_SUCCESS = "点赞成功"
		} else {
			ACTION_SUCCESS = "取消点赞成功"
		}
		return &pb.LikeActionResponse{
			StatusCode: 0,
			StatusMsg:  &ACTION_SUCCESS,
		}, nil
	} else {
		ACTION_SUCCESS := mess
		return &pb.LikeActionResponse{
			StatusCode: -1,
			StatusMsg:  &ACTION_SUCCESS,
		}, nil
	}
}
