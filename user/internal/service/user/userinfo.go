package user

import (
	"context"
	pb "user/api/toktik/user"
	"user/internal/biz"
)

type InfoService struct {
	b *biz.UserServiceBizImpl
	pb.UnimplementedUserInfoServer
}

func NewUserInfoService(bizsrv *biz.UserServiceBizImpl) *InfoService {
	return &InfoService{b: bizsrv}
}

func (s *InfoService) UserInfoSrv(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	service, _ := s.b.UserInfoService(ctx, req)
	return &service, nil
}
