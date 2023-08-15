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
	srv, _ := s.b.UserInfoSrv(ctx, req)
	return &srv, nil
}
