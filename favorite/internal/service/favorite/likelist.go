package service

import (
	"context"
	"encoding/json"
	common "favorite/api/toktik/common"
	pb "favorite/api/toktik/favorite"
	favService "favorite/internal/biz/favorite"
	"github.com/go-kratos/kratos/v2/log"
)

type LikeListService struct {
	pb.UnimplementedLikeListServer
	favService *favService.FavoriteServiceBiz
}

func NewLikeListService(favService *favService.FavoriteServiceBiz) *LikeListService {
	return &LikeListService{
		favService: favService,
	}
}

func (s *LikeListService) LikeListSrv(ctx context.Context, req *pb.LikeListRequest) (*pb.LikeListResponse, error) {
	ret, err := s.favService.LikeListSrv(ctx, req)
	var msg string
	if err != nil {
		msg = "查询异常"
		return &pb.LikeListResponse{
			StatusCode: 0,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, nil
	}
	retVideoList := []*common.Video{}
	DtoUtils(ret, &retVideoList)
	msg = "查询成功"
	return &pb.LikeListResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		VideoList:  retVideoList,
	}, nil
}

// DtoUtils dto转换工具
func DtoUtils(inputClass any, outputClass any) {
	// 把输入转成json
	retJson, err := json.Marshal(inputClass)
	if err != nil {
		log.Error("json序列化转换失败")
		panic(err)
	}
	// json转成输出
	err = json.Unmarshal(retJson, &outputClass)

	if err != nil {
		log.Error("json反序列化转换失败")
		panic(err)
	}
}
