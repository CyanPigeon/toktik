package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"relation/internal/biz"
	"relation/internal/data/model"
	dao "relation/internal/data/model/dao/follow"
)

var (
	ErrUpdateFail = errors.New("更新操作失败，未更新任何行")
)

type followRepoImpl struct {
	data  *Data
	log   *log.Helper
	query *dao.Query
}

// FindFollowList 查询列表
func (f followRepoImpl) FindFollowList(ctx context.Context, follow *model.Follow, action int) ([]*model.Follow, error) {
	query := f.query.Follow.WithContext(ctx)

	if action == biz.QueryTypeFollow {
		//查询关注的列表
		query = query.Where(
			f.query.Follow.UserUID.Eq(follow.UserUID), //他关注的人
			f.query.Follow.Delete.Is(false),
		)
	}

	if action == biz.QueryTypeFollower {
		//查询粉丝列表
		query = query.Where(
			f.query.Follow.FollowUID.Eq(follow.UserUID), //关注他的人
			f.query.Follow.Delete.Is(false),
		)
	}

	return query.Find()
}

// FindRelationExist 查询是否存在关系
func (f followRepoImpl) FindRelationExist(ctx context.Context, follow *model.Follow) (bool, error) {
	query := f.query.Follow.WithContext(ctx)

	_, err := query.Where(f.query.Follow.FollowUID.Eq(follow.FollowUID),
		f.query.Follow.UserUID.Eq(follow.UserUID), f.query.Follow.Delete.Is(false)).First()

	//未找到任何记录
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil

}

// UnFollow 取消关注
func (f followRepoImpl) UnFollow(ctx context.Context, follow *model.Follow) error {

	info, err := f.query.Follow.WithContext(ctx).Where(
		f.query.Follow.FollowUID.Eq(follow.FollowUID),
		f.query.Follow.UserUID.Eq(follow.UserUID),
	).Update(f.query.Follow.Delete, true)

	if err != nil {
		return err
	}

	if info.RowsAffected < 1 {
		return ErrUpdateFail
	}

	return nil
}

func (f followRepoImpl) Follow(ctx context.Context, follow *model.Follow) error {

	//检查是否存在旧的关注信息
	find, err := f.query.Follow.WithContext(ctx).Where(
		f.query.Follow.FollowUID.Eq(follow.FollowUID),
		f.query.Follow.UserUID.Eq(follow.UserUID),
		f.query.Follow.Delete.Is(true),
	).Find()

	if err != nil {
		return err
	}

	if len(find) == 1 {
		//存在已经删除的关注信息，更新为未删除
		info, err := f.query.Follow.WithContext(ctx).Where(f.query.Follow.FollowID.Eq(find[0].FollowID),
			f.query.Follow.Delete.Is(true)).Update(f.query.Follow.Delete, false)

		if err != nil {
			return err
		}

		if info.RowsAffected < 1 {
			return ErrUpdateFail
		}

		return nil
	}

	err = f.query.Follow.WithContext(ctx).Create(follow)

	if err != nil {
		return err
	}

	return nil

}

func NewFollowRepoImpl(data *Data, logger log.Logger) biz.FollowRepo {
	query := dao.Q
	return &followRepoImpl{
		data:  data,
		log:   log.NewHelper(logger),
		query: query,
	}
}
