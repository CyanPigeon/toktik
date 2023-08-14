// Package feed
// @Description
// @Author  Ymri  2023/8/11 15:36
// @Update 2023/8/11 15:36
package feed

import (
	"context"
	pb "feed/api/toktik/feed"
	"feed/internal/data"
	model "feed/internal/data/model"
	toktikVideo "feed/internal/data/model"
	followDao "feed/internal/data/model/dao/follow"
	userVideoDao "feed/internal/data/model/dao/userVideo"
	"time"
)

// LIMIT 一次视频拉去最大值
const LIMIT = 30

type FeedServicebiz interface {
	FeedSrv(context.Context, *pb.FeedRequest) ([]toktikVideo.TokTikVideo, error)
}

func NewFeedServiceImpl(db *data.Data) *FeedServiceImpl {
	return &FeedServiceImpl{
		db: db,
	}
}

type FeedServiceImpl struct {
	feedService FeedServicebiz
	db          *data.Data // add db
}

// FeedSrv 获取feed流
func (t *FeedServiceImpl) FeedSrv(ctx context.Context, req *pb.FeedRequest) (*[]toktikVideo.TokTikVideo, error) {
	// 查询
	//TODO
	//  是否该用户关注
	// 该视屏用户是否点赞
	selectModel := &model.TokTikVideo{}
	retList := []model.TokTikVideo{}
	if req.LatestTime == nil {
		// 没有默认当前时间
		nowTime := time.Now().Unix()
		req.LatestTime = &nowTime
	}
	// 按照时间倒序
	t.db.GormDB.Model(selectModel).Debug().Joins("Author").
		Where("video.created_time<to_timestamp(?)", req.LatestTime).Limit(30).
		Find(&retList)
	// 如果没有用户状态，关注和点赞默认为空，
	// 如果能够查询到用户状态， 则挨个查询条件
	// 根据token查询用户id
	var uid int64 = -1
	if req.Token != nil {
		// TODO
		//  查询用户uid
		uid = 1
	} else {
		uid = -1
	}
	uid = 1
	tempFollowDao := followDao.Q.Follow
	tempLikeDao := userVideoDao.Q.UserVideo
	for index, v := range retList {
		if uid == -1 {
			// 未登录，关注和点赞默认为false
			retList[index].Author.IsFollow = false
			retList[index].IsFavorite = false
		} else {
			// 是否点关注
			count, _ := tempFollowDao.WithContext(ctx).Where(tempFollowDao.UserUID.Eq(uid), tempFollowDao.FollowUID.Eq(v.Author.UID)).Count()
			if count != 0 {
				retList[index].Author.IsFollow = true
			} else {
				retList[index].Author.IsFollow = false
			}
			// 是否点赞
			count, _ = tempLikeDao.WithContext(ctx).Where(tempLikeDao.UserUID.Eq(uid), tempLikeDao.VideoID.Eq(v.VideoID)).Count()
			if count != 0 {
				retList[index].IsFavorite = true
			} else {
				retList[index].IsFavorite = false
			}
		}
	}
	if retList == nil {
		return nil, nil
	}
	return &retList, nil
}
