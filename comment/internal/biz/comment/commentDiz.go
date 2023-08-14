// Package comment
// @Description
// @Author  Ymri  2023/8/13 14:38
// @Update 2023/8/13 14:38
package comment

import (
	pb "comment/api/toktik/comment"
	"comment/internal/data"
	"comment/internal/data/model"
	commentDao "comment/internal/data/model/dao/comment"
	followDao "comment/internal/data/model/dao/follow"
	videoDao "comment/internal/data/model/dao/video"
	dto "comment/internal/data/model/dto"
	utils "comment/internal/utils"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type DizCommentService interface {
	// CommentListSrv 评论列表
	CommentListSrv(ctx context.Context, req *pb.CommentListRequest) ([]*dto.CommentDto, error)
	// CommentActionSrv 评论操作 删除后返回nil
	CommentActionSrv(ctx context.Context, req *pb.CommentActionRequest) (*dto.CommentDto, error)
}

type DizCommentServiceImpl struct {
	DizCommentService
	db *data.Data
}

func NewDizCommentServiceImpl(db *data.Data) *DizCommentServiceImpl {
	return &DizCommentServiceImpl{
		db: db,
	}
}

func (s *DizCommentServiceImpl) CommentListSrv(ctx context.Context, req *pb.CommentListRequest) ([]*dto.CommentDto, string, error) {
	var commentList []*dto.CommentDto
	var uid int64
	if req.Token == "" {
		uid = 1
	} else {
		uid = 2
	}

	tempFollow := followDao.Q.Follow
	fmt.Println(req.Token)
	// vid 检查
	s.db.GormDB.Model(&dto.CommentDto{}).Preload("User").Where("video_id = ? and delete = false", req.VideoId).Find(&commentList)
	for index, v := range commentList {
		// 月份和日
		commentList[index].CreateData = v.CreatedTime.Format("01-02")
		// 查询是否关注
		exit, err := tempFollow.WithContext(ctx).Where(tempFollow.UserUID.Eq(uid), tempFollow.FollowUID.Eq(v.User.UID), tempFollow.Delete.Value(false)).Count()
		if err != nil {
			continue
		}
		if exit == 0 {
			commentList[index].User.IsFollow = false
		} else {
			commentList[index].User.IsFollow = true
		}
	}
	return commentList, "刷新评论成功...", nil
}

func (s *DizCommentServiceImpl) CommentActionSrv(ctx context.Context, req *pb.CommentActionRequest) (bool, *dto.CommentDto, string, error) {
	var uid int64
	var retData *dto.CommentDto
	if req.Token == "" {
		uid = 1
	} else {
		uid = 2
	}

	tempVideo := &videoDao.Q.Video
	exit, err := tempVideo.WithContext(ctx).Where(tempVideo.VideoID.Eq(req.VideoId), tempVideo.Delete.Value(false)).Count()
	if err != nil {
		return false, nil, "查询异常", err
	}
	if exit == 0 {
		return false, nil, "找不到视频(" + string(req.VideoId) + ")", nil
	}
	// 空文本过滤
	if req.CommentText == nil || *req.CommentText == "" {
		return false, nil, "不能发表空评论！", nil
	}
	// 评论操作 1 发布 2 删除
	if req.ActionType == 1 {
		snowId := utils.GenID()
		inserData := &model.Comment{
			CommentID:   snowId,
			VideoID:     req.VideoId,
			UserUID:     uid,
			Content:     *req.CommentText,
			CreatedBy:   uid,
			UpdatedBy:   uid,
			Delete:      false,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		err := commentDao.Q.Comment.WithContext(ctx).Create(inserData)
		if err != nil {
			log.Error("insert comment fail", err)
			return false, nil, "评论异常！", err
		}
		// 再联合查一遍
		tempComment := &commentDao.Q.Comment
		err = s.db.GormDB.Model(retData).Preload("User").Where(tempComment.CommentID.Eq(snowId)).First(&retData).Error
		// 查询是否关注
		retData.User.IsFollow = true
		if err != nil {
			log.Error("insert comment fail", err)
			return false, nil, "评论异常！", err
		}
		// 月份和日
		retData.CreateData = retData.CreatedTime.Format("01-02")
		return true, retData, "评论成功", nil
	} else if req.ActionType == 2 {
		// 只能删自己的评论
		commentId := req.CommentId
		tempComment := &commentDao.Q.Comment
		selectCommen, err := tempComment.WithContext(ctx).Where(tempComment.CommentID.Eq(*commentId)).First()
		if err != nil {
			log.Error("select cmment err:", err)
			return false, nil, "找不到该评论！", err
		}
		if selectCommen.Delete {
			return false, nil, "找不到该评论或已删除", nil
		}
		if selectCommen.UserUID != uid {
			return false, nil, "该评论不属于您，删除失败！", nil
		}
		// 删除
		row, err := tempComment.WithContext(ctx).Where(tempComment.CommentID.Eq(*commentId)).Update(tempComment.Delete, true)
		if row.RowsAffected == 0 && err != nil {
			log.Error("delete comment err:", err)
			return false, nil, "删除评论异常！", err
		} else {
			return true, nil, "删除成功！", nil
		}
	} else {
		return false, nil, "请正确操作！", nil
	}
}
