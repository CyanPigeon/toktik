package user

import (
	"context"
	pb "user/api/toktik/user"
	"user/internal/biz"
)

type UserInfoService struct {
	pb.UnimplementedUserInfoServer
	*biz.UserServiceBizImpl
}

func NewUserInfoService() *UserInfoService {
	return &UserInfoService{}
}

func (s *UserInfoService) UserInfoSrv(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	srv, _ := s.UserInfoSrv(ctx, req)
	return srv, nil
}
