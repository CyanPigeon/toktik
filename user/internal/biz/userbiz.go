package biz

import (
	"context"
	"github.com/CyanPigeon/toktik/middleware"
	"user/api/toktik/common"
	pb "user/api/toktik/user"
	"user/internal/data"
	followDao "user/internal/data/follow"
	"user/internal/data/model"
	userDao "user/internal/data/user"
	"user/util"
)

type UserServiceBiz interface {
	UserRegisterSrv(ctx context.Context, req *pb.UserRegisterRequest) (pb.UserRegisterResponse, error)
	UserLoginSrv(ctx context.Context, req *pb.UserLoginRequest) (pb.UserLoginResponse, error)
	UserInfoSrv(ctx context.Context, req *pb.UserInfoRequest) (pb.UserInfoResponse, error)
}

type UserServiceBizImpl struct {
	UserServiceBiz
	db *data.Data
}

func NewUserServiceImpl(db *data.Data) *UserServiceBizImpl {
	return &UserServiceBizImpl{
		db: db,
	}
}
func (u *UserServiceBizImpl) UserRegisterSrv(ctx context.Context, req *pb.UserRegisterRequest) (pb.UserRegisterResponse, error) {
	paramsOk := util.ValidateParams([]string{req.Username, req.Password})
	em := "invalid username or password"
	if !paramsOk {
		return pb.UserRegisterResponse{
			StatusCode: 101,
			StatusMsg:  &em,
		}, nil
	}
	err := util.SnowflakeInit(0)
	if err != nil {
		em = err.Error()
		return pb.UserRegisterResponse{
			StatusCode: 300,
			StatusMsg:  &em,
		}, err
	}
	q := userDao.Q.WithContext(ctx).User
	uuid := util.GenID()
	encPassword, err := util.AESEncrypt(req.Password)
	if err != nil {
		em = err.Error()
		return pb.UserRegisterResponse{
			StatusCode: 300,
			StatusMsg:  &em,
		}, err
	}
	err = q.Create(&model.User{
		UID:      uuid,
		Username: req.Username,
		Password: encPassword,
	})
	if err != nil {
		em = err.Error()
		return pb.UserRegisterResponse{
			StatusCode: 300,
			StatusMsg:  &em,
		}, nil
	}
	t, _ := middleware.GenToken(uuid)
	return pb.UserRegisterResponse{
		StatusCode: 0,
		UserId:     uuid,
		Token:      t,
	}, nil
}

func (u *UserServiceBizImpl) UserLoginSrv(ctx context.Context, req *pb.UserLoginRequest) (pb.UserLoginResponse, error) {
	paramsOk := util.ValidateParams([]string{req.Username, req.Password})
	em := "invalid username or password"
	if !paramsOk {
		return pb.UserLoginResponse{
			StatusCode: 101,
			StatusMsg:  &em,
		}, nil
	}
	q := userDao.Q.WithContext(ctx).User
	first, err := q.Where(userDao.Q.User.Username.Eq(req.Username)).First()
	if err != nil {
		return pb.UserLoginResponse{
			StatusCode: 101,
			StatusMsg:  &em,
		}, nil
	}
	p, err := util.AESDecrypt(first.Password)
	if first.Password != p || err != nil {
		return pb.UserLoginResponse{
			StatusCode: 101,
			StatusMsg:  &em,
		}, nil
	}
	t, _ := middleware.GenToken(first.UID)
	return pb.UserLoginResponse{
		StatusCode: 0,
		UserId:     first.UID,
		Token:      t,
	}, nil
}

func (u *UserServiceBizImpl) UserInfoSrv(ctx context.Context, req *pb.UserInfoRequest) (pb.UserInfoResponse, error) {
	v, j, _ := middleware.ValidateToken(req.Token)
	q := userDao.Q.WithContext(ctx).User
	first, err := q.Where(userDao.Q.User.UID.Eq(req.UserId)).First()
	if !v {
		if err != nil {
			em := err.Error()
			return pb.UserInfoResponse{
				StatusCode: 201,
				StatusMsg:  &em,
			}, nil
		}
		return pb.UserInfoResponse{
			StatusCode: 0,
			StatusMsg:  nil,
			User: &common.User{
				Id:              first.UID,
				Name:            first.Username,
				FollowCount:     &first.FellowCount,
				FollowerCount:   &first.FellowerCount,
				IsFollow:        false,
				Avatar:          &first.Avatar,
				BackgroundImage: &first.BackgroundImg,
				Signature:       &first.Signature,
				TotalFavorited:  &first.TotalFavorited,
				WorkCount:       &first.VideoCount,
				FavoriteCount:   &first.FavoriteCount,
			},
		}, nil
	}
	fq := followDao.Q.Follow
	c, err := fq.WithContext(ctx).Where(fq.UserUID.Eq(j.UUID), fq.FollowID.Eq(req.UserId)).Count()

	return pb.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		User: &common.User{
			Id:              first.UID,
			Name:            first.Username,
			FollowCount:     &first.FellowCount,
			FollowerCount:   &first.FellowerCount,
			IsFollow:        err == nil && c > 0,
			Avatar:          &first.Avatar,
			BackgroundImage: &first.BackgroundImg,
			Signature:       &first.Signature,
			TotalFavorited:  &first.TotalFavorited,
			WorkCount:       &first.VideoCount,
			FavoriteCount:   &first.FavoriteCount,
		},
	}, nil
}