package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"relation/internal/conf"
	dao "relation/internal/data/model/dao/follow"
	"relation/internal/utils"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewFollowRepoImpl)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	err := utils.Init(1)
	if err != nil {
		panic("snowflake init error")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: c.Database.Source}), &gorm.Config{})
	dao.SetDefault(db)
	// 后续直接用就行
	if err != nil {
		panic(err)
	}
	return &Data{db: db}, cleanup, nil
}
