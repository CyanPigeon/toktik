package data

import (
	"comment/internal/conf"
	commentDao "comment/internal/data/model/dao/comment"
	followDao "comment/internal/data/model/dao/follow"
	userDao "comment/internal/data/model/dao/user"
	videoDao "comment/internal/data/model/dao/video"
	snowInit "comment/internal/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGormDB)

// Data .
type Data struct {
	// add db
	GormDB *gorm.DB
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// init
	userDao.SetDefault(db)
	videoDao.SetDefault(db)
	commentDao.SetDefault(db)
	followDao.SetDefault(db)
	// TODO 分布式ID
	snowInit.Init(0)
	return &Data{GormDB: db}, cleanup, nil
}

func NewGormDB(c *conf.Data) (*gorm.DB, error) {
	dsn := c.Database.Source
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db = db.Debug()
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(time.Second * 25)
	return db, nil
}
