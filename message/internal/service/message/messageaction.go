package service

import (
	"context"
	pb "message/api/toktik/message"
	"message/internal/biz/message"
)

type MessageActionService struct {
	pb.UnimplementedMessageActionServer
	messageServiceImpl *message.BizMessageServiceImpl
}

func NewMessageActionService(messageServiceImpl *message.BizMessageServiceImpl) *MessageActionService {
	return &MessageActionService{
		messageServiceImpl: messageServiceImpl,
	}
}

func (s *MessageActionService) MessageActionSrv(ctx context.Context, req *pb.MessageActionRequest) (*pb.MessageActionResponse, error) {
	success, mes, err := s.messageServiceImpl.MessageActionSrv(ctx, req)
	if success {
		return &pb.MessageActionResponse{
			StatusCode: 0,
			StatusMsg:  &mes,
		}, nil
	} else {
		return &pb.MessageActionResponse{
			StatusCode: 1,
			StatusMsg:  &mes,
		}, err
	}
}
