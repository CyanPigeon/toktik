package user

import (
	"context"
	"user/internal/biz"
	"user/util"

	pb "user/api/toktik/user"
)

var sferr = util.SnowflakeInit(0)

type UserRegisterService struct {
	biz.UserServiceBizImpl
	pb.UnimplementedUserRegisterServer
}

func NewUserRegisterService() *UserRegisterService {
	return &UserRegisterService{}
}

func (s *UserRegisterService) UserRegisterSrv(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	srv, _ := s.UserRegisterSrv(ctx, req)
	return srv, nil
}
