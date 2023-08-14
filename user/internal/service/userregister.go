package service

import (
	"context"
	"github.com/CyanPigeon/toktik/middleware"
	"user/util"

	pb "user/api/toktik/user"
)

var sferr = util.SnowflakeInit(1145141919810)

type UserRegisterService struct {
	pb.UnimplementedUserRegisterServer
}

func NewUserRegisterService() *UserRegisterService {
	return &UserRegisterService{}
}

func (s *UserRegisterService) UserRegisterSrv(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	valid := util.ValidateParams([]string{req.Username, req.Password})
	em := "register failed."
	if !valid {
		return &pb.UserRegisterResponse{
			StatusCode: 101,
			StatusMsg:  &em,
		}, nil
	}
	if sferr != nil {
		em = sferr.Error()
		return &pb.UserRegisterResponse{
			StatusCode: 300,
			StatusMsg:  &em,
		}, nil
	}
	uuid := util.GenID()
	// TODO: Db
	t, _ := middleware.GenToken(uuid)
	return &pb.UserRegisterResponse{
		StatusCode: 0,
		UserId:     uuid,
		Token:      t,
	}, nil
}
