package data

import (
	"github.com/CyanPigeon/toktik/app/demo/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDatabaseClient, NewEntityRepoImpl)

// DatabaseClient 数据客户端包装器
type DatabaseClient struct {
	// TODO 数据库客户端对象
}

// NewDatabaseClient 是DatabaseClient构造函数
func NewDatabaseClient(conf *conf.Data, logger log.Logger) (*DatabaseClient, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// conf.Data里面有数据库的连接信息，可以根据里面的数据连接数据库。
	// TODO 完成对数据库客户端的初始化工作
	return &DatabaseClient{}, cleanup, nil
}
