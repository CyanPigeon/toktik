package utils

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Database struct {
	Db *gorm.DB
}

type IDBQuery interface {
	ExecuteRawSql(string, map[string]interface{}, interface{}) error
	Query(interface{}, *interface{}) error
	Insert(interface{}) error
	GetDB() *gorm.DB
}

// InitDB 首先调用InitDB，传入当前业务的数据库，返回IDBQuery
func InitDB(db *gorm.DB) *Database {
	return &Database{Db: db}
}

// GetDB 返回调用InitDB时传入的DB
func (db *Database) GetDB() *gorm.DB {
	return db.Db
}

// Query 在当前数据库查询,comp为比较,model为返回的模型
//
// 如:
//
// Query(user.Q.User.UID.Eq(1),&User)
//
// 返回的查询结果将存储在User中,查询失败将返回error
func (db *Database) Query(comp interface{}, model *interface{}) error {
	r := db.Db.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Where(comp).First(model)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// Insert 在当前数据库插入,comp为比较,model为返回的模型
//
// 如:
//
//	Insert(&User{
//		xxx : xxx
//		})
//
// 插入成功将返回nil,否则返回error
func (db *Database) Insert(data interface{}) error {
	r := db.Db.Clauses(clause.Locking{
		Strength: "ROW EXCLUSIVE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Create(&data)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// ExecuteRawSql 直接执行SQL语句,不建议用,不保证稳定性
//
// s为RawSql语句，m为用于替换RawSql中?部分的map，model为返回的模型
func (db *Database) ExecuteRawSql(s string, m map[string]interface{}, model interface{}) error {
	r := db.Db.Exec(s, m).Scan(&model)
	if r.Error != nil {
		return r.Error
	}
	return nil
}
