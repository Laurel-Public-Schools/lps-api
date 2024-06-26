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

func newVerificationToken(db *gorm.DB, opts ...gen.DOOption) verificationToken {
	_verificationToken := verificationToken{}

	_verificationToken.verificationTokenDo.UseDB(db, opts...)
	_verificationToken.verificationTokenDo.UseModel(&model.VerificationToken{})

	tableName := _verificationToken.verificationTokenDo.TableName()
	_verificationToken.ALL = field.NewAsterisk(tableName)
	_verificationToken.Identifier = field.NewString(tableName, "identifier")
	_verificationToken.Token = field.NewString(tableName, "token")
	_verificationToken.Expires = field.NewTime(tableName, "expires")

	_verificationToken.fillFieldMap()

	return _verificationToken
}

type verificationToken struct {
	verificationTokenDo verificationTokenDo

	ALL        field.Asterisk
	Identifier field.String
	Token      field.String
	Expires    field.Time

	fieldMap map[string]field.Expr
}

func (v verificationToken) Table(newTableName string) *verificationToken {
	v.verificationTokenDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v verificationToken) As(alias string) *verificationToken {
	v.verificationTokenDo.DO = *(v.verificationTokenDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *verificationToken) updateTableName(table string) *verificationToken {
	v.ALL = field.NewAsterisk(table)
	v.Identifier = field.NewString(table, "identifier")
	v.Token = field.NewString(table, "token")
	v.Expires = field.NewTime(table, "expires")

	v.fillFieldMap()

	return v
}

func (v *verificationToken) WithContext(ctx context.Context) IVerificationTokenDo {
	return v.verificationTokenDo.WithContext(ctx)
}

func (v verificationToken) TableName() string { return v.verificationTokenDo.TableName() }

func (v verificationToken) Alias() string { return v.verificationTokenDo.Alias() }

func (v verificationToken) Columns(cols ...field.Expr) gen.Columns {
	return v.verificationTokenDo.Columns(cols...)
}

func (v *verificationToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *verificationToken) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 3)
	v.fieldMap["identifier"] = v.Identifier
	v.fieldMap["token"] = v.Token
	v.fieldMap["expires"] = v.Expires
}

func (v verificationToken) clone(db *gorm.DB) verificationToken {
	v.verificationTokenDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v verificationToken) replaceDB(db *gorm.DB) verificationToken {
	v.verificationTokenDo.ReplaceDB(db)
	return v
}

type verificationTokenDo struct{ gen.DO }

type IVerificationTokenDo interface {
	gen.SubQuery
	Debug() IVerificationTokenDo
	WithContext(ctx context.Context) IVerificationTokenDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVerificationTokenDo
	WriteDB() IVerificationTokenDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVerificationTokenDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVerificationTokenDo
	Not(conds ...gen.Condition) IVerificationTokenDo
	Or(conds ...gen.Condition) IVerificationTokenDo
	Select(conds ...field.Expr) IVerificationTokenDo
	Where(conds ...gen.Condition) IVerificationTokenDo
	Order(conds ...field.Expr) IVerificationTokenDo
	Distinct(cols ...field.Expr) IVerificationTokenDo
	Omit(cols ...field.Expr) IVerificationTokenDo
	Join(table schema.Tabler, on ...field.Expr) IVerificationTokenDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVerificationTokenDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVerificationTokenDo
	Group(cols ...field.Expr) IVerificationTokenDo
	Having(conds ...gen.Condition) IVerificationTokenDo
	Limit(limit int) IVerificationTokenDo
	Offset(offset int) IVerificationTokenDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVerificationTokenDo
	Unscoped() IVerificationTokenDo
	Create(values ...*model.VerificationToken) error
	CreateInBatches(values []*model.VerificationToken, batchSize int) error
	Save(values ...*model.VerificationToken) error
	First() (*model.VerificationToken, error)
	Take() (*model.VerificationToken, error)
	Last() (*model.VerificationToken, error)
	Find() ([]*model.VerificationToken, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VerificationToken, err error)
	FindInBatches(result *[]*model.VerificationToken, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.VerificationToken) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVerificationTokenDo
	Assign(attrs ...field.AssignExpr) IVerificationTokenDo
	Joins(fields ...field.RelationField) IVerificationTokenDo
	Preload(fields ...field.RelationField) IVerificationTokenDo
	FirstOrInit() (*model.VerificationToken, error)
	FirstOrCreate() (*model.VerificationToken, error)
	FindByPage(offset int, limit int) (result []*model.VerificationToken, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVerificationTokenDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v verificationTokenDo) Debug() IVerificationTokenDo {
	return v.withDO(v.DO.Debug())
}

func (v verificationTokenDo) WithContext(ctx context.Context) IVerificationTokenDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v verificationTokenDo) ReadDB() IVerificationTokenDo {
	return v.Clauses(dbresolver.Read)
}

func (v verificationTokenDo) WriteDB() IVerificationTokenDo {
	return v.Clauses(dbresolver.Write)
}

func (v verificationTokenDo) Session(config *gorm.Session) IVerificationTokenDo {
	return v.withDO(v.DO.Session(config))
}

func (v verificationTokenDo) Clauses(conds ...clause.Expression) IVerificationTokenDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v verificationTokenDo) Returning(value interface{}, columns ...string) IVerificationTokenDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v verificationTokenDo) Not(conds ...gen.Condition) IVerificationTokenDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v verificationTokenDo) Or(conds ...gen.Condition) IVerificationTokenDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v verificationTokenDo) Select(conds ...field.Expr) IVerificationTokenDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v verificationTokenDo) Where(conds ...gen.Condition) IVerificationTokenDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v verificationTokenDo) Order(conds ...field.Expr) IVerificationTokenDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v verificationTokenDo) Distinct(cols ...field.Expr) IVerificationTokenDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v verificationTokenDo) Omit(cols ...field.Expr) IVerificationTokenDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v verificationTokenDo) Join(table schema.Tabler, on ...field.Expr) IVerificationTokenDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v verificationTokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVerificationTokenDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v verificationTokenDo) RightJoin(table schema.Tabler, on ...field.Expr) IVerificationTokenDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v verificationTokenDo) Group(cols ...field.Expr) IVerificationTokenDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v verificationTokenDo) Having(conds ...gen.Condition) IVerificationTokenDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v verificationTokenDo) Limit(limit int) IVerificationTokenDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v verificationTokenDo) Offset(offset int) IVerificationTokenDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v verificationTokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVerificationTokenDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v verificationTokenDo) Unscoped() IVerificationTokenDo {
	return v.withDO(v.DO.Unscoped())
}

func (v verificationTokenDo) Create(values ...*model.VerificationToken) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v verificationTokenDo) CreateInBatches(values []*model.VerificationToken, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v verificationTokenDo) Save(values ...*model.VerificationToken) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v verificationTokenDo) First() (*model.VerificationToken, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VerificationToken), nil
	}
}

func (v verificationTokenDo) Take() (*model.VerificationToken, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VerificationToken), nil
	}
}

func (v verificationTokenDo) Last() (*model.VerificationToken, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VerificationToken), nil
	}
}

func (v verificationTokenDo) Find() ([]*model.VerificationToken, error) {
	result, err := v.DO.Find()
	return result.([]*model.VerificationToken), err
}

func (v verificationTokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VerificationToken, err error) {
	buf := make([]*model.VerificationToken, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v verificationTokenDo) FindInBatches(result *[]*model.VerificationToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v verificationTokenDo) Attrs(attrs ...field.AssignExpr) IVerificationTokenDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v verificationTokenDo) Assign(attrs ...field.AssignExpr) IVerificationTokenDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v verificationTokenDo) Joins(fields ...field.RelationField) IVerificationTokenDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v verificationTokenDo) Preload(fields ...field.RelationField) IVerificationTokenDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v verificationTokenDo) FirstOrInit() (*model.VerificationToken, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VerificationToken), nil
	}
}

func (v verificationTokenDo) FirstOrCreate() (*model.VerificationToken, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VerificationToken), nil
	}
}

func (v verificationTokenDo) FindByPage(offset int, limit int) (result []*model.VerificationToken, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v verificationTokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v verificationTokenDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v verificationTokenDo) Delete(models ...*model.VerificationToken) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *verificationTokenDo) withDO(do gen.Dao) *verificationTokenDo {
	v.DO = *do.(*gen.DO)
	return v
}
