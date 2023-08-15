package user

import (
	"context"
	pb "user/api/toktik/user"
	"user/internal/biz"
)

type UserLoginService struct {
	*biz.UserServiceBizImpl
	pb.UnimplementedUserLoginServer
}

func NewUserLoginService() *UserLoginService {
	return &UserLoginService{}
}

func (s *UserLoginService) UserLoginSrv(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	srv, _ := s.UserLoginSrv(ctx, req)
	return srv, nil
}
