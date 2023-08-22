package utils

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Database struct {
	Db *gorm.DB
}

type IDBQuery interface {
	ExecuteRawSql(string, map[string]interface{}, interface{}) (bool, error)
	Query(interface{}, *interface{}) (bool, error)
	Insert(interface{}) (bool, error)
	GetDB() *gorm.DB
}

func InitDB(db *gorm.DB) *Database {
	return &Database{Db: db}
}

func (db Database) Query(comp interface{}, model *interface{}) (bool, error) {
	r := db.Db.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Where(comp).First(model)
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}

func (db Database) Insert(i interface{}) (bool, error) {
	r := db.Db.Clauses(clause.Locking{
		Strength: "ROW EXCLUSIVE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Create(&i)
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}

func (db Database) GetDB() *gorm.DB {
	return db.Db
}

func (db *Database) ExecuteRawSql(s string, m map[string]interface{}, model interface{}) (bool, error) {
	r := db.Db.Exec(s, m).Scan(&model)
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}
