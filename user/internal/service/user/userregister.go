package user

import (
	"context"
	"user/internal/biz"
	"user/util"

	pb "user/api/toktik/user"
)

var sferr = util.SnowflakeInit(0)

type RegisterService struct {
	b *biz.UserServiceBizImpl
	pb.UnimplementedUserRegisterServer
}

func NewUserRegisterService(bizsrv *biz.UserServiceBizImpl) *RegisterService {
	return &RegisterService{b: bizsrv}
}

func (s *RegisterService) UserRegisterSrv(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	srv, _ := s.b.UserRegisterSrv(ctx, req)
	return &srv, nil
}
