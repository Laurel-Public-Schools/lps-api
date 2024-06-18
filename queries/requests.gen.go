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

func newRequest(db *gorm.DB, opts ...gen.DOOption) request {
	_request := request{}

	_request.requestDo.UseDB(db, opts...)
	_request.requestDo.UseModel(&model.Request{})

	tableName := _request.requestDo.TableName()
	_request.ALL = field.NewAsterisk(tableName)
	_request.ID = field.NewInt32(tableName, "id")
	_request.StudentID = field.NewString(tableName, "studentId")
	_request.StudentName = field.NewString(tableName, "studentName")
	_request.NewTeacher = field.NewString(tableName, "newTeacher")
	_request.NewTeacherName = field.NewString(tableName, "newTeacherName")
	_request.CurrentTeacher = field.NewString(tableName, "currentTeacher")
	_request.CurrentTeacherName = field.NewString(tableName, "currentTeacherName")
	_request.DateRequested = field.NewTime(tableName, "dateRequested")
	_request.Status = field.NewString(tableName, "status")
	_request.Arrived = field.NewBool(tableName, "arrived")
	_request.Timestamp = field.NewString(tableName, "timestamp")

	_request.fillFieldMap()

	return _request
}

type request struct {
	requestDo requestDo

	ALL                field.Asterisk
	ID                 field.Int32
	StudentID          field.String
	StudentName        field.String
	NewTeacher         field.String
	NewTeacherName     field.String
	CurrentTeacher     field.String
	CurrentTeacherName field.String
	DateRequested      field.Time
	Status             field.String
	Arrived            field.Bool
	Timestamp          field.String

	fieldMap map[string]field.Expr
}

func (r request) Table(newTableName string) *request {
	r.requestDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r request) As(alias string) *request {
	r.requestDo.DO = *(r.requestDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *request) updateTableName(table string) *request {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.StudentID = field.NewString(table, "studentId")
	r.StudentName = field.NewString(table, "studentName")
	r.NewTeacher = field.NewString(table, "newTeacher")
	r.NewTeacherName = field.NewString(table, "newTeacherName")
	r.CurrentTeacher = field.NewString(table, "currentTeacher")
	r.CurrentTeacherName = field.NewString(table, "currentTeacherName")
	r.DateRequested = field.NewTime(table, "dateRequested")
	r.Status = field.NewString(table, "status")
	r.Arrived = field.NewBool(table, "arrived")
	r.Timestamp = field.NewString(table, "timestamp")

	r.fillFieldMap()

	return r
}

func (r *request) WithContext(ctx context.Context) IRequestDo { return r.requestDo.WithContext(ctx) }

func (r request) TableName() string { return r.requestDo.TableName() }

func (r request) Alias() string { return r.requestDo.Alias() }

func (r request) Columns(cols ...field.Expr) gen.Columns { return r.requestDo.Columns(cols...) }

func (r *request) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *request) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 11)
	r.fieldMap["id"] = r.ID
	r.fieldMap["studentId"] = r.StudentID
	r.fieldMap["studentName"] = r.StudentName
	r.fieldMap["newTeacher"] = r.NewTeacher
	r.fieldMap["newTeacherName"] = r.NewTeacherName
	r.fieldMap["currentTeacher"] = r.CurrentTeacher
	r.fieldMap["currentTeacherName"] = r.CurrentTeacherName
	r.fieldMap["dateRequested"] = r.DateRequested
	r.fieldMap["status"] = r.Status
	r.fieldMap["arrived"] = r.Arrived
	r.fieldMap["timestamp"] = r.Timestamp
}

func (r request) clone(db *gorm.DB) request {
	r.requestDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r request) replaceDB(db *gorm.DB) request {
	r.requestDo.ReplaceDB(db)
	return r
}

type requestDo struct{ gen.DO }

type IRequestDo interface {
	gen.SubQuery
	Debug() IRequestDo
	WithContext(ctx context.Context) IRequestDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRequestDo
	WriteDB() IRequestDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRequestDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRequestDo
	Not(conds ...gen.Condition) IRequestDo
	Or(conds ...gen.Condition) IRequestDo
	Select(conds ...field.Expr) IRequestDo
	Where(conds ...gen.Condition) IRequestDo
	Order(conds ...field.Expr) IRequestDo
	Distinct(cols ...field.Expr) IRequestDo
	Omit(cols ...field.Expr) IRequestDo
	Join(table schema.Tabler, on ...field.Expr) IRequestDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRequestDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRequestDo
	Group(cols ...field.Expr) IRequestDo
	Having(conds ...gen.Condition) IRequestDo
	Limit(limit int) IRequestDo
	Offset(offset int) IRequestDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRequestDo
	Unscoped() IRequestDo
	Create(values ...*model.Request) error
	CreateInBatches(values []*model.Request, batchSize int) error
	Save(values ...*model.Request) error
	First() (*model.Request, error)
	Take() (*model.Request, error)
	Last() (*model.Request, error)
	Find() ([]*model.Request, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Request, err error)
	FindInBatches(result *[]*model.Request, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Request) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRequestDo
	Assign(attrs ...field.AssignExpr) IRequestDo
	Joins(fields ...field.RelationField) IRequestDo
	Preload(fields ...field.RelationField) IRequestDo
	FirstOrInit() (*model.Request, error)
	FirstOrCreate() (*model.Request, error)
	FindByPage(offset int, limit int) (result []*model.Request, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRequestDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r requestDo) Debug() IRequestDo {
	return r.withDO(r.DO.Debug())
}

func (r requestDo) WithContext(ctx context.Context) IRequestDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r requestDo) ReadDB() IRequestDo {
	return r.Clauses(dbresolver.Read)
}

func (r requestDo) WriteDB() IRequestDo {
	return r.Clauses(dbresolver.Write)
}

func (r requestDo) Session(config *gorm.Session) IRequestDo {
	return r.withDO(r.DO.Session(config))
}

func (r requestDo) Clauses(conds ...clause.Expression) IRequestDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r requestDo) Returning(value interface{}, columns ...string) IRequestDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r requestDo) Not(conds ...gen.Condition) IRequestDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r requestDo) Or(conds ...gen.Condition) IRequestDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r requestDo) Select(conds ...field.Expr) IRequestDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r requestDo) Where(conds ...gen.Condition) IRequestDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r requestDo) Order(conds ...field.Expr) IRequestDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r requestDo) Distinct(cols ...field.Expr) IRequestDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r requestDo) Omit(cols ...field.Expr) IRequestDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r requestDo) Join(table schema.Tabler, on ...field.Expr) IRequestDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r requestDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRequestDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r requestDo) RightJoin(table schema.Tabler, on ...field.Expr) IRequestDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r requestDo) Group(cols ...field.Expr) IRequestDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r requestDo) Having(conds ...gen.Condition) IRequestDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r requestDo) Limit(limit int) IRequestDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r requestDo) Offset(offset int) IRequestDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r requestDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRequestDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r requestDo) Unscoped() IRequestDo {
	return r.withDO(r.DO.Unscoped())
}

func (r requestDo) Create(values ...*model.Request) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r requestDo) CreateInBatches(values []*model.Request, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r requestDo) Save(values ...*model.Request) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r requestDo) First() (*model.Request, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Request), nil
	}
}

func (r requestDo) Take() (*model.Request, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Request), nil
	}
}

func (r requestDo) Last() (*model.Request, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Request), nil
	}
}

func (r requestDo) Find() ([]*model.Request, error) {
	result, err := r.DO.Find()
	return result.([]*model.Request), err
}

func (r requestDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Request, err error) {
	buf := make([]*model.Request, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r requestDo) FindInBatches(result *[]*model.Request, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r requestDo) Attrs(attrs ...field.AssignExpr) IRequestDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r requestDo) Assign(attrs ...field.AssignExpr) IRequestDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r requestDo) Joins(fields ...field.RelationField) IRequestDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r requestDo) Preload(fields ...field.RelationField) IRequestDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r requestDo) FirstOrInit() (*model.Request, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Request), nil
	}
}

func (r requestDo) FirstOrCreate() (*model.Request, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Request), nil
	}
}

func (r requestDo) FindByPage(offset int, limit int) (result []*model.Request, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r requestDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r requestDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r requestDo) Delete(models ...*model.Request) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *requestDo) withDO(do gen.Dao) *requestDo {
	r.DO = *do.(*gen.DO)
	return r
}
