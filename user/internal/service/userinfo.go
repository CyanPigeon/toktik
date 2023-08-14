package service

import (
	"context"
	"github.com/CyanPigeon/toktik/middleware"
	"user/api/toktik/common"
	pb "user/api/toktik/user"
)

type UserInfoService struct {
	pb.UnimplementedUserInfoServer
}

func NewUserInfoService() *UserInfoService {
	return &UserInfoService{}
}

func (s *UserInfoService) UserInfoSrv(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	v, j, err := middleware.ValidateToken(req.Token)
	errs := "invalid token"
	if !v || err != nil {
		return &pb.UserInfoResponse{
			StatusCode: 102,
			StatusMsg:  &errs,
		}, nil
	}
	// TODO: Db
	return &pb.UserInfoResponse{
		StatusCode: 0,
		User: &common.User{
			Id: j.UUID,
		},
	}, nil
}
