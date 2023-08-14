package service

import (
	diz "comment/internal/biz/comment"
	"context"

	pb "comment/api/toktik/comment"
	common "comment/api/toktik/common"
	utils "comment/internal/utils"
)

type CommentActionService struct {
	pb.UnimplementedCommentActionServer
	dizCommentService *diz.DizCommentServiceImpl
}

func NewCommentActionService(dizCommentService *diz.DizCommentServiceImpl) *CommentActionService {
	return &CommentActionService{
		dizCommentService: dizCommentService,
	}
}

func (s *CommentActionService) CommentActionSrv(ctx context.Context, req *pb.CommentActionRequest) (*pb.CommentActionResponse, error) {
	success, commentDto, mess, _ := s.dizCommentService.CommentActionSrv(ctx, req)
	code := int32(0)
	if !success {
		code = 1
	}
	// 异常或者删除评论
	if commentDto == nil {
		return &pb.CommentActionResponse{
			StatusCode: code,
			StatusMsg:  &mess,
			Comment:    nil,
		}, nil
	}
	// dto
	comment := common.Comment{}
	utils.DtoUtils(commentDto, &comment)
	return &pb.CommentActionResponse{
		StatusCode: code,
		StatusMsg:  &mess,
		Comment:    &comment,
	}, nil
}
