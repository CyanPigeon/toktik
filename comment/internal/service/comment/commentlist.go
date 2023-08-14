package service

import (
	diz "comment/internal/biz/comment"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "comment/api/toktik/comment"
	common "comment/api/toktik/common"
	utils "comment/internal/utils"
)

type CommentListService struct {
	pb.UnimplementedCommentListServer
	dizCommentService *diz.DizCommentServiceImpl
}

func NewCommentListService(dizCommentService *diz.DizCommentServiceImpl) *CommentListService {
	return &CommentListService{
		dizCommentService: dizCommentService,
	}
}

func (s *CommentListService) CommentListSrv(ctx context.Context, req *pb.CommentListRequest) (*pb.CommentListResponse, error) {
	log.Info(req.Token)
	commentList, mess, _ := s.dizCommentService.CommentListSrv(ctx, req)
	// dto
	dtoCommentList := make([]*common.Comment, 0)
	utils.DtoUtils(commentList, &dtoCommentList)
	return &pb.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   &mess,
		CommentList: dtoCommentList,
	}, nil
}
