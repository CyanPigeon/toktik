package service

import (
	v1 "github.com/CyanPigeon/toktik/api/demo/v1"
	"github.com/CyanPigeon/toktik/app/demo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// DemoService Api的业务代码挂靠的结构体。
// TODO 结构体名需要根据Api名进行修改，后缀统一为Service
// TODO 如果有多个Api，请拆分为多个文件，一个文件一个Api
type DemoService struct {
	v1.UnimplementedDemoServer

	repo biz.EntityRepo
	log  *log.Helper
}

// NewDemoService 是DemoService的构造函数。
// TODO 函数名需要根据Api名进行修改，前缀统一为New
// TODO 如果有多个Api，请拆分为多个文件，一个文件一个Api
// TODO 重命名同时也需要修改service/service.go文件中ProviderSet指定的函数名
func NewService(repo biz.EntityRepo, logger log.Logger) *DemoService {
	return &DemoService{repo: repo, log: log.NewHelper(logger)}
}

// TODO 实现api/${api_name}/${version}/${api_name}_http.pb.go中的${ApiName}HTTPServer
// TODO   和api/${api_name}/${version}/${api_name}_grpc.pb.go中的${ApiName}Server接口。
// TODO 这两个接口为api/${api_name}/${version}/${api_name}.proto中service定义的接口。
// TODO 没搞懂的可以看主仓库的示例代码。
