package service

import (
	"context"
	"message/api/toktik/common"
	"message/internal/biz/message"
	"message/internal/utils"

	pb "message/api/toktik/message"
)

type MessageHistoryService struct {
	pb.UnimplementedMessageHistoryServer
	bizMessageservice *message.BizMessageServiceImpl
}

func NewMessageHistoryService(bizMessageservice *message.BizMessageServiceImpl) *MessageHistoryService {
	return &MessageHistoryService{
		bizMessageservice: bizMessageservice,
	}
}

func (s *MessageHistoryService) MessageHistorySrv(ctx context.Context, req *pb.MessageHistoryRequest) (*pb.MessageHistoryResponse, error) {
	success, mes, messageList, _ := s.bizMessageservice.MessageHistorySrv(ctx, req)
	if success {
		// 对象转换
		messageDtoList := []*common.Message{}
		utils.DtoUtils(messageList, &messageDtoList)
		return &pb.MessageHistoryResponse{
			StatusCode:  0,
			StatusMsg:   &mes,
			MessageList: messageDtoList,
		}, nil
	} else {
		return &pb.MessageHistoryResponse{
			StatusCode:  1,
			StatusMsg:   &mes,
			MessageList: []*common.Message{},
		}, nil
	}
}
