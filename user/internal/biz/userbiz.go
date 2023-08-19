package biz

import (
	"context"
	"github.com/CyanPigeon/toktik/middleware"
	"gorm.io/gorm/clause"
	"user/api/toktik/common"
	pb "user/api/toktik/user"
	"user/internal/data"
	followDao "user/internal/data/follow"
	"user/internal/data/model"
	userDao "user/internal/data/user"
	"user/util"
)

type UserServiceBiz interface {
	UserRegisterService(ctx context.Context, req *pb.UserRegisterRequest) (pb.UserRegisterResponse, error)
	UserLoginService(ctx context.Context, req *pb.UserLoginRequest) (pb.UserLoginResponse, error)
	UserInfoService(ctx context.Context, req *pb.UserInfoRequest) (pb.UserInfoResponse, error)
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
func (u *UserServiceBizImpl) UserRegisterService(ctx context.Context, req *pb.UserRegisterRequest) (pb.UserRegisterResponse, error) {
	paramsOk := util.ValidateParams([]string{req.Username, req.Password})
	em := "invalid username or password"
	if !paramsOk {
		return pb.UserRegisterResponse{
			StatusCode: 101,
			StatusMsg:  &em,
		}, nil
	}
	q := userDao.Q.WithContext(ctx).User
	first, _ := q.Where(userDao.Q.User.Username.Eq(req.Username)).First()
	if first != nil {
		em = "user exist."
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
	uuid := util.GenID()
	encPassword, err := util.AESEncrypt(req.Password)
	if err != nil {
		em = err.Error()
		return pb.UserRegisterResponse{
			StatusCode: 300,
			StatusMsg:  &em,
		}, err
	}
	c, _ := q.Count()
	uuid += c
	u.db.GormDB.Clauses(clause.Locking{
		Strength: "RowExclusiveLock",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Create(&model.User{
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

func (u *UserServiceBizImpl) UserLoginService(ctx context.Context, req *pb.UserLoginRequest) (pb.UserLoginResponse, error) {
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
	if req.Password != p || err != nil {
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

func (u *UserServiceBizImpl) UserInfoService(ctx context.Context, req *pb.UserInfoRequest) (pb.UserInfoResponse, error) {
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
				FollowCount:     &first.FollowCount,
				FollowerCount:   &first.FollowerCount,
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
			FollowCount:     &first.FollowCount,
			FollowerCount:   &first.FollowerCount,
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
