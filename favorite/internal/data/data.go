package data

import (
	"favorite/internal/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"

	userDao "favorite/internal/data/model/dao/user"
	userLikeVideoDao "favorite/internal/data/model/dao/userVideo"
	videoDao "favorite/internal/data/model/dao/video"
	snowflake "favorite/internal/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
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

	// add dao
	videoDao.SetDefault(db)
	userDao.SetDefault(db)
	userLikeVideoDao.SetDefault(db)
	// init snowflake
	// TODO 分布式编号
	err := snowflake.Init(0)
	if err != nil {
		panic("snowflake init error")
	}
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
