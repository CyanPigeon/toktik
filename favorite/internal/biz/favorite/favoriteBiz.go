// Package favorite
// @Description
// @Author  Ymri  2023/8/12 21:04
// @Update 2023/8/12 21:04
package favorite

import (
	"context"
	pb "favorite/api/toktik/favorite"
	"favorite/internal/data"
	model "favorite/internal/data/model"
	userVideoDao "favorite/internal/data/model/dao/userVideo"
	videoDao "favorite/internal/data/model/dao/video"
	snowFlake "favorite/internal/utils"
	"github.com/go-kratos/kratos/v2/log"
	"sync"
	"time"
)

const (
	ACTION_TYPE_LIKE   = 1
	ACTION_TYPE_UNLIKE = 2
)

var (
	myMap = make(map[int]int, 1)
	// 声明一个全局的互斥锁
	// lock：一个全局的互斥锁。
	// sync：包，用于同步（synchornized）。
	// Mutex：互斥
	lock sync.Mutex
)

type FavoriteService interface {
	LikeActionSrv(context.Context, *pb.LikeActionRequest) (bool, string, error)
}

type FavoriteServiceBiz struct {
	FavoriteService
	db *data.Data
}

func NewFavoriteServiceBiz(db *data.Data) *FavoriteServiceBiz {
	return &FavoriteServiceBiz{
		db: db,
	}
}

// LikeActionSrv 点赞
func (t *FavoriteServiceBiz) LikeActionSrv(ctx context.Context, req *pb.LikeActionRequest) (bool, string, error) {
	// 权限pass
	var uid = int64(1)
	// TODO 根据token查询用户id
	//
	if req.Token == "" {
		uid = 1
	} else {
		uid = 2

	}
	// 检查参数
	if req.VideoId <= 0 {
		return false, "找不到视频", nil
	}
	var tempLike bool
	if ACTION_TYPE_LIKE == req.ActionType {
		tempLike = false
	} else if ACTION_TYPE_UNLIKE == req.ActionType {
		tempLike = true
	} else {
		return false, "请规范点赞类型", nil
	}
	// 查询是否点赞
	// 查询视频是否存在
	//  insert into student values(4,'d') on conflict(id) do update set name='as';
	tempVideoDao := &videoDao.Q.Video
	exit, err := videoDao.Q.Video.WithContext(ctx).Where(tempVideoDao.VideoID.Eq(req.VideoId)).Count()
	if exit != 1 || err != nil {
		return false, "找不到视频", nil
	}
	// 雪花ID生成

	// 这里是否加事务
	// 插入视频
	// 先查询，后更新或者插入
	var insertUserVideo *model.UserVideo
	tempUserVideoDao := &userVideoDao.Q.UserVideo
	insertUserVideo, err = tempUserVideoDao.WithContext(ctx).Where(tempUserVideoDao.UserUID.Eq(uid), tempUserVideoDao.VideoID.Eq(req.VideoId)).First()
	if err != nil {
		log.Error("insert user video error", err)
		return false, "网络堵塞，请稍后再试！", err
	}
	if insertUserVideo == nil {
		// insert
		snowId := snowFlake.GenID()
		insertUserVideo = &model.UserVideo{
			LikeID:      snowId,
			UserUID:     uid,
			VideoID:     req.VideoId,
			CreatedBy:   uid,
			UpdatedBy:   uid,
			Delete:      false,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		err = tempUserVideoDao.WithContext(ctx).Create(insertUserVideo)
		if err != nil {
			log.Error("insert user video error", err)
			return false, "网络堵塞，请稍后再试！", err
		}
	} else {
		if insertUserVideo.Delete == tempLike {
			return true, "", nil
		}
		insertUserVideo.Delete = tempLike
		_, err = tempUserVideoDao.WithContext(ctx).Where(tempUserVideoDao.LikeID.Eq(insertUserVideo.LikeID)).Update(tempUserVideoDao.Delete, tempLike)
		if err != nil {
			log.Error("insert user video error", err)
			return false, "网络堵塞，请稍后再试！", err
		}
	}

	// TODO 1.加锁 2.定时任务刷新到数据库
	//redis 更新视频信息 更新用户信息
	// 添加并发锁，但是针对于那张表只修改一个字段
	//更新视频  后续直接rpc调用
	lock.Lock()
	var updateSQl string
	if tempLike {
		// 取消
		updateSQl = "-1"
	} else {
		// ++
		updateSQl = "+1"
	}
	err = t.db.GormDB.Debug().Raw("update video set like_count = like_count "+updateSQl+" where video_id = ?", req.VideoId).Error
	err = t.db.GormDB.Debug().Raw("UPDATE tiktok_user SET total_favorited = total_favorited"+updateSQl+" WHERE uid = ?", uid).Error
	defer lock.Unlock()
	return true, "", nil
}

// LikeListSrv 点赞列表
func (t *FavoriteServiceBiz) LikeListSrv(ctx context.Context, req *pb.LikeListRequest) ([]*model.VideoDto, error) {
	var uid int64
	if req.Token == "" {
		uid = 1
	} else {
		uid = 2
	}
	// 不需要用户id,根据token就能够直接拿到用户id
	retVideoList := make([]*model.VideoDto, 0)
	// 查询所有的点赞过的视频
	err := t.db.GormDB.Debug().Model(&model.VideoDto{}).Preload("Author").Where("video_id in (?)", t.db.GormDB.Debug().Model(&model.UserVideo{}).Select("video_id").Where("user_uid = ? and delete=false", uid)).Find(&retVideoList).Error
	if err != nil {
		return nil, err
	}
	for _, video := range retVideoList {
		video.IsFavorite = true
	}
	return retVideoList, nil

}
