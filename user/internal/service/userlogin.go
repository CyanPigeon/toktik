package service

import (
	"context"
	"github.com/CyanPigeon/toktik/middleware"
	"user/util"

	pb "user/api/toktik/user"
)

type UserLoginService struct {
	pb.UnimplementedUserLoginServer
}

func NewUserLoginService() *UserLoginService {
	return &UserLoginService{}
}

func (s *UserLoginService) UserLoginSrv(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	valid := util.ValidateParams([]string{req.Username, req.Password})
	errs := "login failed."
	if !valid {
		return &pb.UserLoginResponse{
			StatusCode: 101,
			StatusMsg:  &errs,
		}, nil
	}
	// TODO: Db
	// 0 is a placeholder
	t, _ := middleware.GenToken(0)
	return &pb.UserLoginResponse{
		StatusCode: 0,
		UserId:     0,
		Token:      t,
	}, nil
}
