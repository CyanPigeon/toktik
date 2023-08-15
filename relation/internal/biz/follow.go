package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"relation/api/toktik/common"
	pb "relation/api/toktik/relation"
	"relation/internal/data/model"
	"relation/internal/utils"
)

const (
	QueryTypeFollow = iota
	QueryTypeFollower
)

var (
	ErrRelationNotFound = errors.New("关注不存在")
	ErrRelationExists   = errors.New("已经关注了")
)

type FollowRepo interface {
	Follow(ctx context.Context, follow *model.Follow) error
	UnFollow(ctx context.Context, follow *model.Follow) error
	FindRelationExist(ctx context.Context, follow *model.Follow) (bool, error)
	FindFollowList(ctx context.Context, follow *model.Follow, action int) ([]*model.Follow, error)
}

type FollowUsecase struct {
	repo FollowRepo
	log  *log.Helper
}

func NewFollowUsecase(repo FollowRepo, logger log.Logger) *FollowUsecase {
	return &FollowUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Follow 关注
func (uc *FollowUsecase) Follow(ctx context.Context, req *pb.FollowActionRequest) error {
	token := req.Token
	var user *common.User
	if token != "" {
		//TODO 根据token获取操作者uid 测试操作
		user = GetUser(ctx, 1)
	}

	f := model.Follow{}

	f.UserUID = user.Id
	has, err := uc.HasRelation(ctx, &f) //检查是否已经存在关注
	if err != nil {
		return err
	}

	if has {
		// 关系存在 不可以再关注，返回错误
		return ErrRelationExists
	}

	f.FollowID = utils.GenID()
	f.CreatedBy = user.Id
	f.UpdatedBy = user.Id
	err = uc.repo.Follow(ctx, &f)
	if err != nil {
		return err
	}
	return nil
}

// UnFollow 取消关注
func (uc *FollowUsecase) UnFollow(ctx context.Context, req *pb.FollowActionRequest) error {
	token := req.Token
	var user *common.User
	if token != "" {
		//TODO 根据token获取操作者uid 测试操作
		user = GetUser(ctx, 1)
	}

	f := model.Follow{}
	f.UserUID = user.Id

	has, err := uc.HasRelation(ctx, &f)
	if err != nil {
		return err
	}

	if has {
		// 关系存在，可以取消
		f.Delete = true
		err := uc.repo.UnFollow(ctx, &f)
		if err != nil {
			return err
		}
		return nil
	}

	// 没有关系
	return ErrRelationNotFound
}

// HasRelation 查询关系，返回查询数量
func (uc *FollowUsecase) HasRelation(ctx context.Context, f *model.Follow) (bool, error) {
	return uc.repo.FindRelationExist(ctx, f)
}

// GetFollowList 获取关注列表，根据action判断获取的是粉丝列表还是关注列表，为了方便不直接传入req
func (uc *FollowUsecase) GetFollowList(ctx context.Context, f *model.Follow, action int) ([]*common.User, error) {
	follows, err := uc.repo.FindFollowList(ctx, f, action)
	if err != nil {
		return nil, err
	}
	var users []*common.User

	if action == QueryTypeFollow {
		for _, value := range follows {
			//TODO 根据follow信息查询关注者信息
			user := GetUser(ctx, value.UserUID)
			users = append(users, user)

		}
	}

	if action == QueryTypeFollower {
		for _, value := range follows {
			//TODO 根据follow信息查询粉丝信息
			user := GetUser(ctx, value.UserUID)
			users = append(users, user)

		}
	}

	return users, nil
}

// TODO 测试需要，在User服务搭建完成之前直接访问数据库
func GetUser(ctx context.Context, uid int64) *common.User {
	var user common.User
	user = common.User{Id: uid}
	return &user
}
