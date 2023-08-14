package service

import (
	"context"
	"encoding/json"
	common "feed/api/toktik/common"
	pb "feed/api/toktik/feed"
	feed "feed/internal/biz/feed"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type FeedService struct {
	pb.UnimplementedFeedServer
	feedServiceImpl feed.FeedServiceImpl
}

func NewFeedService(feedservice *feed.FeedServiceImpl) *FeedService {
	return &FeedService{
		feedServiceImpl: *feedservice,
	}
}

func (s *FeedService) FeedSrv(ctx context.Context, req *pb.FeedRequest) (*pb.FeedResponse, error) {
	// TODO token检查用户权限
	ret, err := s.feedServiceImpl.FeedSrv(ctx, req)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return nil, err
	}
	//a, _ := json.Marshal(ret)
	//log.Info(string(a))
	retVideo := []*common.Video{}
	// 这里转的时候丢失了
	retMsg := "success"

	DtoUtils(ret, &retVideo)
	//a, _ := json.Marshal(retVideo)
	//log.Info(string(a))
	nowTime := time.Now().Unix()
	//a, _ := json.Marshal(retVideo)
	//log.Info(string(a))
	return &pb.FeedResponse{
		StatusCode: 0,
		StatusMsg:  &retMsg,
		VideoList:  retVideo,
		NextTime:   &nowTime,
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
