// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package userVideo

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"feed/internal/data/model"
)

func newUserVideo(db *gorm.DB, opts ...gen.DOOption) userVideo {
	_userVideo := userVideo{}

	_userVideo.userVideoDo.UseDB(db, opts...)
	_userVideo.userVideoDo.UseModel(&model.UserVideo{})

	tableName := _userVideo.userVideoDo.TableName()
	_userVideo.ALL = field.NewAsterisk(tableName)
	_userVideo.LikeID = field.NewInt64(tableName, "like_id")
	_userVideo.UserUID = field.NewInt64(tableName, "user_uid")
	_userVideo.VideoID = field.NewInt64(tableName, "video_id")
	_userVideo.CreatedBy = field.NewInt64(tableName, "created_by")
	_userVideo.CreatedTime = field.NewTime(tableName, "created_time")
	_userVideo.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_userVideo.UpdatedTime = field.NewTime(tableName, "updated_time")
	_userVideo.Delete = field.NewBool(tableName, "delete")

	_userVideo.fillFieldMap()

	return _userVideo
}

type userVideo struct {
	userVideoDo userVideoDo

	ALL         field.Asterisk
	LikeID      field.Int64 // 点赞ID;点赞ID
	UserUID     field.Int64 // 用户UID;用户UID
	VideoID     field.Int64 // 视频ID;视频ID
	CreatedBy   field.Int64 // 创建人
	CreatedTime field.Time  // 创建时间
	UpdatedBy   field.Int64 // 更新人
	UpdatedTime field.Time  // 更新时间
	Delete      field.Bool  // 是否删除

	fieldMap map[string]field.Expr
}

func (u userVideo) Table(newTableName string) *userVideo {
	u.userVideoDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userVideo) As(alias string) *userVideo {
	u.userVideoDo.DO = *(u.userVideoDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userVideo) updateTableName(table string) *userVideo {
	u.ALL = field.NewAsterisk(table)
	u.LikeID = field.NewInt64(table, "like_id")
	u.UserUID = field.NewInt64(table, "user_uid")
	u.VideoID = field.NewInt64(table, "video_id")
	u.CreatedBy = field.NewInt64(table, "created_by")
	u.CreatedTime = field.NewTime(table, "created_time")
	u.UpdatedBy = field.NewInt64(table, "updated_by")
	u.UpdatedTime = field.NewTime(table, "updated_time")
	u.Delete = field.NewBool(table, "delete")

	u.fillFieldMap()

	return u
}

func (u *userVideo) WithContext(ctx context.Context) IUserVideoDo {
	return u.userVideoDo.WithContext(ctx)
}

func (u userVideo) TableName() string { return u.userVideoDo.TableName() }

func (u userVideo) Alias() string { return u.userVideoDo.Alias() }

func (u userVideo) Columns(cols ...field.Expr) gen.Columns { return u.userVideoDo.Columns(cols...) }

func (u *userVideo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userVideo) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 8)
	u.fieldMap["like_id"] = u.LikeID
	u.fieldMap["user_uid"] = u.UserUID
	u.fieldMap["video_id"] = u.VideoID
	u.fieldMap["created_by"] = u.CreatedBy
	u.fieldMap["created_time"] = u.CreatedTime
	u.fieldMap["updated_by"] = u.UpdatedBy
	u.fieldMap["updated_time"] = u.UpdatedTime
	u.fieldMap["delete"] = u.Delete
}

func (u userVideo) clone(db *gorm.DB) userVideo {
	u.userVideoDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userVideo) replaceDB(db *gorm.DB) userVideo {
	u.userVideoDo.ReplaceDB(db)
	return u
}

type userVideoDo struct{ gen.DO }

type IUserVideoDo interface {
	gen.SubQuery
	Debug() IUserVideoDo
	WithContext(ctx context.Context) IUserVideoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserVideoDo
	WriteDB() IUserVideoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserVideoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserVideoDo
	Not(conds ...gen.Condition) IUserVideoDo
	Or(conds ...gen.Condition) IUserVideoDo
	Select(conds ...field.Expr) IUserVideoDo
	Where(conds ...gen.Condition) IUserVideoDo
	Order(conds ...field.Expr) IUserVideoDo
	Distinct(cols ...field.Expr) IUserVideoDo
	Omit(cols ...field.Expr) IUserVideoDo
	Join(table schema.Tabler, on ...field.Expr) IUserVideoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserVideoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserVideoDo
	Group(cols ...field.Expr) IUserVideoDo
	Having(conds ...gen.Condition) IUserVideoDo
	Limit(limit int) IUserVideoDo
	Offset(offset int) IUserVideoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserVideoDo
	Unscoped() IUserVideoDo
	Create(values ...*model.UserVideo) error
	CreateInBatches(values []*model.UserVideo, batchSize int) error
	Save(values ...*model.UserVideo) error
	First() (*model.UserVideo, error)
	Take() (*model.UserVideo, error)
	Last() (*model.UserVideo, error)
	Find() ([]*model.UserVideo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserVideo, err error)
	FindInBatches(result *[]*model.UserVideo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserVideo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserVideoDo
	Assign(attrs ...field.AssignExpr) IUserVideoDo
	Joins(fields ...field.RelationField) IUserVideoDo
	Preload(fields ...field.RelationField) IUserVideoDo
	FirstOrInit() (*model.UserVideo, error)
	FirstOrCreate() (*model.UserVideo, error)
	FindByPage(offset int, limit int) (result []*model.UserVideo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserVideoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userVideoDo) Debug() IUserVideoDo {
	return u.withDO(u.DO.Debug())
}

func (u userVideoDo) WithContext(ctx context.Context) IUserVideoDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userVideoDo) ReadDB() IUserVideoDo {
	return u.Clauses(dbresolver.Read)
}

func (u userVideoDo) WriteDB() IUserVideoDo {
	return u.Clauses(dbresolver.Write)
}

func (u userVideoDo) Session(config *gorm.Session) IUserVideoDo {
	return u.withDO(u.DO.Session(config))
}

func (u userVideoDo) Clauses(conds ...clause.Expression) IUserVideoDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userVideoDo) Returning(value interface{}, columns ...string) IUserVideoDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userVideoDo) Not(conds ...gen.Condition) IUserVideoDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userVideoDo) Or(conds ...gen.Condition) IUserVideoDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userVideoDo) Select(conds ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userVideoDo) Where(conds ...gen.Condition) IUserVideoDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userVideoDo) Order(conds ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userVideoDo) Distinct(cols ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userVideoDo) Omit(cols ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userVideoDo) Join(table schema.Tabler, on ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userVideoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userVideoDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userVideoDo) Group(cols ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userVideoDo) Having(conds ...gen.Condition) IUserVideoDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userVideoDo) Limit(limit int) IUserVideoDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userVideoDo) Offset(offset int) IUserVideoDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userVideoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserVideoDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userVideoDo) Unscoped() IUserVideoDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userVideoDo) Create(values ...*model.UserVideo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userVideoDo) CreateInBatches(values []*model.UserVideo, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userVideoDo) Save(values ...*model.UserVideo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userVideoDo) First() (*model.UserVideo, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserVideo), nil
	}
}

func (u userVideoDo) Take() (*model.UserVideo, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserVideo), nil
	}
}

func (u userVideoDo) Last() (*model.UserVideo, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserVideo), nil
	}
}

func (u userVideoDo) Find() ([]*model.UserVideo, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserVideo), err
}

func (u userVideoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserVideo, err error) {
	buf := make([]*model.UserVideo, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userVideoDo) FindInBatches(result *[]*model.UserVideo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userVideoDo) Attrs(attrs ...field.AssignExpr) IUserVideoDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userVideoDo) Assign(attrs ...field.AssignExpr) IUserVideoDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userVideoDo) Joins(fields ...field.RelationField) IUserVideoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userVideoDo) Preload(fields ...field.RelationField) IUserVideoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userVideoDo) FirstOrInit() (*model.UserVideo, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserVideo), nil
	}
}

func (u userVideoDo) FirstOrCreate() (*model.UserVideo, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserVideo), nil
	}
}

func (u userVideoDo) FindByPage(offset int, limit int) (result []*model.UserVideo, count int64, err error) {
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

func (u userVideoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userVideoDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userVideoDo) Delete(models ...*model.UserVideo) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userVideoDo) withDO(do gen.Dao) *userVideoDo {
	u.DO = *do.(*gen.DO)
	return u
}
