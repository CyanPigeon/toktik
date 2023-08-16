package user

import (
	"context"
	pb "user/api/toktik/user"
	"user/internal/biz"
)

type LoginService struct {
	b *biz.UserServiceBizImpl
	pb.UnimplementedUserLoginServer
}

func NewUserLoginService(bizsrv *biz.UserServiceBizImpl) *LoginService {
	return &LoginService{b: bizsrv}
}

func (s *LoginService) UserLoginSrv(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	service, _ := s.b.UserLoginService(ctx, req)
	return &service, nil
}
