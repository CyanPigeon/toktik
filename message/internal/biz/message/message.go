// Package message
// @Description
// @Author  Ymri  2023/8/13 20:59
// @Update 2023/8/13 20:59
package message

import (
	"context"
	pb "message/api/toktik/message"
	"message/internal/data"
	"message/internal/data/model"
	messageDao "message/internal/data/model/dao/message"
	userDao "message/internal/data/model/dao/user"
	"message/internal/data/model/dto"
	"message/internal/utils"

	"time"
)

type BizMessageService interface {
	MessageActionSrv(ctx context.Context, req *pb.MessageActionRequest) (bool, string, error)
	MessageHistorySrv(ctx context.Context, req *pb.MessageHistoryRequest) (bool, string, []*model.Message, error)
}

type BizMessageServiceImpl struct {
	BizMessageService
	db *data.Data
}

func NewBizMessageService(db *data.Data) *BizMessageServiceImpl {
	return &BizMessageServiceImpl{db: db}
}

func (s *BizMessageServiceImpl) MessageActionSrv(ctx context.Context, req *pb.MessageActionRequest) (bool, string, error) {
	// 只有消息发送
	var uid int64
	if req.Token != "" {
		uid = 1
	} else {
		uid = 2
	}
	if req.Content == "" {
		return false, "消息不能为空", nil
	}
	// 目前只支持消息发送
	if req.ActionType != 1 {
		return false, "不支持的消息类型", nil
	}
	// 检查接收的用户是否存在
	tempUserDao := userDao.Q.User
	exit, err := tempUserDao.WithContext(ctx).Where(tempUserDao.UID.Eq(req.ToUserId)).Count()
	if err != nil {
		return false, "查询用户失败", err
	}
	if exit == 0 {
		return false, "用户不存在", nil
	}

	newUid := utils.GenID()
	insertMessage := &model.Message{
		MessageID:   newUid,
		UIDSend:     uid,
		UIDReceive:  req.ToUserId,
		Content:     req.Content,
		CreatedBy:   uid,
		UpdatedBy:   uid,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	err = messageDao.Q.Message.WithContext(ctx).Create(insertMessage)
	if err != nil {
		return false, "发送消息失败", err
	}
	return true, "发送成功", nil
}
func (s *BizMessageServiceImpl) MessageHistorySrv(ctx context.Context, req *pb.MessageHistoryRequest) (bool, string, []*dto.MessageDto, error) {

	var uid int64
	if req.Token != "" {
		uid = 1
	} else {
		uid = 2
	}
	// 检查接收的用户是否存在
	tempUserDao := userDao.Q.User
	exit, err := tempUserDao.WithContext(ctx).Where(tempUserDao.UID.Eq(req.ToUserId)).Count()
	if err != nil {
		return false, "查询消息记录失败！", nil, err
	}
	if exit == 0 {
		return false, "查无此用户", nil, nil
	}
	// 查询消息记录
	tempMessageDao := messageDao.Q.Message
	retData := make([]*dto.MessageDto, 0)
	err = s.db.GormDB.Model(&dto.MessageDto{}).WithContext(ctx).
		Where(
			tempMessageDao.WithContext(ctx).Where(tempMessageDao.UIDSend.Eq(uid), tempMessageDao.UIDReceive.Eq(req.ToUserId))).Or(
		tempMessageDao.WithContext(ctx).Where(tempMessageDao.UIDSend.Eq(req.ToUserId), tempMessageDao.UIDReceive.Eq(uid))).Find(&retData).Error
	//messageList, err := s.db.GormDB.Model(&dto.MessageDto{}).WithContext(ctx).
	//	Where(tempMessageDao.UIDSend.Eq(uid), tempMessageDao.UIDReceive.Eq(req.ToUserId)).
	//	Or(tempMessageDao.WithContext(ctx).Where(tempMessageDao.UIDSend.Eq(uid))).Find()
	if err != nil {
		return false, "查询消息记录失败！", nil, err
	}
	for index, _ := range retData {
		retData[index].CreateTime = retData[index].CreatedTime.Format("2006-01-02 15:04:05")
	}
	return true, "查询成功", retData, nil
}
