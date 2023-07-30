package data

import (
	"github.com/CyanPigeon/toktik/app/demo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// EntityRepoImpl 实体的读/写/查询操作接口实现。
// TODO 'EntityRepo'需要替换为对应的接口名，后缀统一为Impl
type EntityRepoImpl struct {
	client *DatabaseClient
	log    *log.Helper
}

// NewEntityRepoImpl 构造函数
// TODO 这个函数名改不改都行。无论是否修改，都需要删除该TODO。
func NewEntityRepoImpl(client *DatabaseClient, logger log.Logger) biz.EntityRepo {
	return &EntityRepoImpl{
		client: client,
		log:    log.NewHelper(logger),
	}
}

// TODO 为EntityRepoImpl实现接口EntityRepo
