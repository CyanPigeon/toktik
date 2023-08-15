// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package user

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"user/internal/data/model"
)

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(&model.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewAsterisk(tableName)
	_user.UID = field.NewInt64(tableName, "uid")
	_user.Username = field.NewString(tableName, "username")
	_user.Password = field.NewString(tableName, "password")
	_user.Avatar = field.NewString(tableName, "avatar")
	_user.BackgroundImg = field.NewString(tableName, "background_img")
	_user.TotalFavorited = field.NewInt64(tableName, "total_favorited")
	_user.FellowCount = field.NewInt64(tableName, "fellow_count")
	_user.FellowerCount = field.NewInt64(tableName, "fellower_count")
	_user.VideoCount = field.NewInt64(tableName, "video_count")
	_user.FavoriteCount = field.NewInt64(tableName, "favorite_count")
	_user.Signature = field.NewString(tableName, "signature")
	_user.CreatedBy = field.NewInt64(tableName, "created_by")
	_user.CreatedTime = field.NewTime(tableName, "created_time")
	_user.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_user.UpdatedTime = field.NewTime(tableName, "updated_time")
	_user.Delete = field.NewBool(tableName, "delete")

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo userDo

	ALL            field.Asterisk
	UID            field.Int64  // UID
	Username       field.String // 用户昵称
	Password       field.String // 加密后的密码
	Avatar         field.String // 头像;头像地址，有默认头像
	BackgroundImg  field.String // 背景图片;有默认地址
	TotalFavorited field.Int64  // 总获赞数;被点赞数
	FellowCount    field.Int64  // 关注人数;关注数量，默认0
	FellowerCount  field.Int64  // 被关注人数;被关注数量，默认0
	VideoCount     field.Int64  // 作品数量;作品数量，默认0
	FavoriteCount  field.Int64  // 总点赞数;总点赞数
	Signature      field.String // 个人简介;个人简介
	CreatedBy      field.Int64  // 创建人;创建人
	CreatedTime    field.Time   // 创建时间;创建时间
	UpdatedBy      field.Int64  // 更新人
	UpdatedTime    field.Time   // 更新时间;更新时间
	Delete         field.Bool   // 是否删除;是否删除

	fieldMap map[string]field.Expr
}

func (u user) Table(newTableName string) *user {
	u.userDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user) updateTableName(table string) *user {
	u.ALL = field.NewAsterisk(table)
	u.UID = field.NewInt64(table, "uid")
	u.Username = field.NewString(table, "username")
	u.Password = field.NewString(table, "password")
	u.Avatar = field.NewString(table, "avatar")
	u.BackgroundImg = field.NewString(table, "background_img")
	u.TotalFavorited = field.NewInt64(table, "total_favorited")
	u.FellowCount = field.NewInt64(table, "fellow_count")
	u.FellowerCount = field.NewInt64(table, "fellower_count")
	u.VideoCount = field.NewInt64(table, "video_count")
	u.FavoriteCount = field.NewInt64(table, "favorite_count")
	u.Signature = field.NewString(table, "signature")
	u.CreatedBy = field.NewInt64(table, "created_by")
	u.CreatedTime = field.NewTime(table, "created_time")
	u.UpdatedBy = field.NewInt64(table, "updated_by")
	u.UpdatedTime = field.NewTime(table, "updated_time")
	u.Delete = field.NewBool(table, "delete")

	u.fillFieldMap()

	return u
}

func (u *user) WithContext(ctx context.Context) IUserDo { return u.userDo.WithContext(ctx) }

func (u user) TableName() string { return u.userDo.TableName() }

func (u user) Alias() string { return u.userDo.Alias() }

func (u user) Columns(cols ...field.Expr) gen.Columns { return u.userDo.Columns(cols...) }

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 16)
	u.fieldMap["uid"] = u.UID
	u.fieldMap["username"] = u.Username
	u.fieldMap["password"] = u.Password
	u.fieldMap["avatar"] = u.Avatar
	u.fieldMap["background_img"] = u.BackgroundImg
	u.fieldMap["total_favorited"] = u.TotalFavorited
	u.fieldMap["fellow_count"] = u.FellowCount
	u.fieldMap["fellower_count"] = u.FellowerCount
	u.fieldMap["video_count"] = u.VideoCount
	u.fieldMap["favorite_count"] = u.FavoriteCount
	u.fieldMap["signature"] = u.Signature
	u.fieldMap["created_by"] = u.CreatedBy
	u.fieldMap["created_time"] = u.CreatedTime
	u.fieldMap["updated_by"] = u.UpdatedBy
	u.fieldMap["updated_time"] = u.UpdatedTime
	u.fieldMap["delete"] = u.Delete
}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u user) replaceDB(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userDo struct{ gen.DO }

type IUserDo interface {
	gen.SubQuery
	Debug() IUserDo
	WithContext(ctx context.Context) IUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserDo
	WriteDB() IUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserDo
	Not(conds ...gen.Condition) IUserDo
	Or(conds ...gen.Condition) IUserDo
	Select(conds ...field.Expr) IUserDo
	Where(conds ...gen.Condition) IUserDo
	Order(conds ...field.Expr) IUserDo
	Distinct(cols ...field.Expr) IUserDo
	Omit(cols ...field.Expr) IUserDo
	Join(table schema.Tabler, on ...field.Expr) IUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserDo
	Group(cols ...field.Expr) IUserDo
	Having(conds ...gen.Condition) IUserDo
	Limit(limit int) IUserDo
	Offset(offset int) IUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo
	Unscoped() IUserDo
	Create(values ...*model.User) error
	CreateInBatches(values []*model.User, batchSize int) error
	Save(values ...*model.User) error
	First() (*model.User, error)
	Take() (*model.User, error)
	Last() (*model.User, error)
	Find() ([]*model.User, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User, err error)
	FindInBatches(result *[]*model.User, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.User) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserDo
	Assign(attrs ...field.AssignExpr) IUserDo
	Joins(fields ...field.RelationField) IUserDo
	Preload(fields ...field.RelationField) IUserDo
	FirstOrInit() (*model.User, error)
	FirstOrCreate() (*model.User, error)
	FindByPage(offset int, limit int) (result []*model.User, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userDo) Debug() IUserDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) IUserDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) ReadDB() IUserDo {
	return u.Clauses(dbresolver.Read)
}

func (u userDo) WriteDB() IUserDo {
	return u.Clauses(dbresolver.Write)
}

func (u userDo) Session(config *gorm.Session) IUserDo {
	return u.withDO(u.DO.Session(config))
}

func (u userDo) Clauses(conds ...clause.Expression) IUserDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Returning(value interface{}, columns ...string) IUserDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userDo) Not(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) IUserDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) IUserDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() IUserDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*model.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*model.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Take() (*model.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Last() (*model.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Find() ([]*model.User, error) {
	result, err := u.DO.Find()
	return result.([]*model.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User, err error) {
	buf := make([]*model.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*model.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userDo) Preload(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userDo) FirstOrInit() (*model.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FirstOrCreate() (*model.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*model.User, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userDo) Delete(models ...*model.User) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}
