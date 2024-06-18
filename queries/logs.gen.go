// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package queries

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/laurel-public-schools/lps-api/model"
)

func newLog(db *gorm.DB, opts ...gen.DOOption) log {
	_log := log{}

	_log.logDo.UseDB(db, opts...)
	_log.logDo.UseModel(&model.Log{})

	tableName := _log.logDo.TableName()
	_log.ALL = field.NewAsterisk(tableName)
	_log.ID = field.NewString(tableName, "id")
	_log.User = field.NewString(tableName, "user")
	_log.Type = field.NewString(tableName, "type")
	_log.Action = field.NewString(tableName, "action")

	_log.fillFieldMap()

	return _log
}

type log struct {
	logDo logDo

	ALL    field.Asterisk
	ID     field.String
	User   field.String
	Type   field.String
	Action field.String

	fieldMap map[string]field.Expr
}

func (l log) Table(newTableName string) *log {
	l.logDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l log) As(alias string) *log {
	l.logDo.DO = *(l.logDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *log) updateTableName(table string) *log {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewString(table, "id")
	l.User = field.NewString(table, "user")
	l.Type = field.NewString(table, "type")
	l.Action = field.NewString(table, "action")

	l.fillFieldMap()

	return l
}

func (l *log) WithContext(ctx context.Context) ILogDo { return l.logDo.WithContext(ctx) }

func (l log) TableName() string { return l.logDo.TableName() }

func (l log) Alias() string { return l.logDo.Alias() }

func (l log) Columns(cols ...field.Expr) gen.Columns { return l.logDo.Columns(cols...) }

func (l *log) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *log) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 4)
	l.fieldMap["id"] = l.ID
	l.fieldMap["user"] = l.User
	l.fieldMap["type"] = l.Type
	l.fieldMap["action"] = l.Action
}

func (l log) clone(db *gorm.DB) log {
	l.logDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l log) replaceDB(db *gorm.DB) log {
	l.logDo.ReplaceDB(db)
	return l
}

type logDo struct{ gen.DO }

type ILogDo interface {
	gen.SubQuery
	Debug() ILogDo
	WithContext(ctx context.Context) ILogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILogDo
	WriteDB() ILogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILogDo
	Not(conds ...gen.Condition) ILogDo
	Or(conds ...gen.Condition) ILogDo
	Select(conds ...field.Expr) ILogDo
	Where(conds ...gen.Condition) ILogDo
	Order(conds ...field.Expr) ILogDo
	Distinct(cols ...field.Expr) ILogDo
	Omit(cols ...field.Expr) ILogDo
	Join(table schema.Tabler, on ...field.Expr) ILogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILogDo
	Group(cols ...field.Expr) ILogDo
	Having(conds ...gen.Condition) ILogDo
	Limit(limit int) ILogDo
	Offset(offset int) ILogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILogDo
	Unscoped() ILogDo
	Create(values ...*model.Log) error
	CreateInBatches(values []*model.Log, batchSize int) error
	Save(values ...*model.Log) error
	First() (*model.Log, error)
	Take() (*model.Log, error)
	Last() (*model.Log, error)
	Find() ([]*model.Log, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Log, err error)
	FindInBatches(result *[]*model.Log, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Log) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILogDo
	Assign(attrs ...field.AssignExpr) ILogDo
	Joins(fields ...field.RelationField) ILogDo
	Preload(fields ...field.RelationField) ILogDo
	FirstOrInit() (*model.Log, error)
	FirstOrCreate() (*model.Log, error)
	FindByPage(offset int, limit int) (result []*model.Log, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l logDo) Debug() ILogDo {
	return l.withDO(l.DO.Debug())
}

func (l logDo) WithContext(ctx context.Context) ILogDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l logDo) ReadDB() ILogDo {
	return l.Clauses(dbresolver.Read)
}

func (l logDo) WriteDB() ILogDo {
	return l.Clauses(dbresolver.Write)
}

func (l logDo) Session(config *gorm.Session) ILogDo {
	return l.withDO(l.DO.Session(config))
}

func (l logDo) Clauses(conds ...clause.Expression) ILogDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l logDo) Returning(value interface{}, columns ...string) ILogDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l logDo) Not(conds ...gen.Condition) ILogDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l logDo) Or(conds ...gen.Condition) ILogDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l logDo) Select(conds ...field.Expr) ILogDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l logDo) Where(conds ...gen.Condition) ILogDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l logDo) Order(conds ...field.Expr) ILogDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l logDo) Distinct(cols ...field.Expr) ILogDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l logDo) Omit(cols ...field.Expr) ILogDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l logDo) Join(table schema.Tabler, on ...field.Expr) ILogDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l logDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILogDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l logDo) RightJoin(table schema.Tabler, on ...field.Expr) ILogDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l logDo) Group(cols ...field.Expr) ILogDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l logDo) Having(conds ...gen.Condition) ILogDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l logDo) Limit(limit int) ILogDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l logDo) Offset(offset int) ILogDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l logDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILogDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l logDo) Unscoped() ILogDo {
	return l.withDO(l.DO.Unscoped())
}

func (l logDo) Create(values ...*model.Log) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l logDo) CreateInBatches(values []*model.Log, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l logDo) Save(values ...*model.Log) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l logDo) First() (*model.Log, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Log), nil
	}
}

func (l logDo) Take() (*model.Log, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Log), nil
	}
}

func (l logDo) Last() (*model.Log, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Log), nil
	}
}

func (l logDo) Find() ([]*model.Log, error) {
	result, err := l.DO.Find()
	return result.([]*model.Log), err
}

func (l logDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Log, err error) {
	buf := make([]*model.Log, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l logDo) FindInBatches(result *[]*model.Log, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l logDo) Attrs(attrs ...field.AssignExpr) ILogDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l logDo) Assign(attrs ...field.AssignExpr) ILogDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l logDo) Joins(fields ...field.RelationField) ILogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l logDo) Preload(fields ...field.RelationField) ILogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l logDo) FirstOrInit() (*model.Log, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Log), nil
	}
}

func (l logDo) FirstOrCreate() (*model.Log, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Log), nil
	}
}

func (l logDo) FindByPage(offset int, limit int) (result []*model.Log, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l logDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l logDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l logDo) Delete(models ...*model.Log) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *logDo) withDO(do gen.Dao) *logDo {
	l.DO = *do.(*gen.DO)
	return l
}
