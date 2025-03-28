// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/realcaishen/utils-go/dal/model"
)

func newTGappedBlock(db *gorm.DB, opts ...gen.DOOption) tGappedBlock {
	_tGappedBlock := tGappedBlock{}

	_tGappedBlock.tGappedBlockDo.UseDB(db, opts...)
	_tGappedBlock.tGappedBlockDo.UseModel(&model.TGappedBlock{})

	tableName := _tGappedBlock.tGappedBlockDo.TableName()
	_tGappedBlock.ALL = field.NewAsterisk(tableName)
	_tGappedBlock.Chainid = field.NewInt32(tableName, "chainid")
	_tGappedBlock.Appid = field.NewInt32(tableName, "appid")
	_tGappedBlock.BlockNumber = field.NewInt64(tableName, "block_number")
	_tGappedBlock.UpdateTimestamp = field.NewTime(tableName, "update_timestamp")
	_tGappedBlock.InsertTimestamp = field.NewTime(tableName, "insert_timestamp")
	_tGappedBlock.IsProcessed = field.NewInt32(tableName, "is_processed")

	_tGappedBlock.fillFieldMap()

	return _tGappedBlock
}

type tGappedBlock struct {
	tGappedBlockDo tGappedBlockDo

	ALL             field.Asterisk
	Chainid         field.Int32
	Appid           field.Int32
	BlockNumber     field.Int64
	UpdateTimestamp field.Time
	InsertTimestamp field.Time
	IsProcessed     field.Int32

	fieldMap map[string]field.Expr
}

func (t tGappedBlock) Table(newTableName string) *tGappedBlock {
	t.tGappedBlockDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tGappedBlock) As(alias string) *tGappedBlock {
	t.tGappedBlockDo.DO = *(t.tGappedBlockDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tGappedBlock) updateTableName(table string) *tGappedBlock {
	t.ALL = field.NewAsterisk(table)
	t.Chainid = field.NewInt32(table, "chainid")
	t.Appid = field.NewInt32(table, "appid")
	t.BlockNumber = field.NewInt64(table, "block_number")
	t.UpdateTimestamp = field.NewTime(table, "update_timestamp")
	t.InsertTimestamp = field.NewTime(table, "insert_timestamp")
	t.IsProcessed = field.NewInt32(table, "is_processed")

	t.fillFieldMap()

	return t
}

func (t *tGappedBlock) WithContext(ctx context.Context) ITGappedBlockDo {
	return t.tGappedBlockDo.WithContext(ctx)
}

func (t tGappedBlock) TableName() string { return t.tGappedBlockDo.TableName() }

func (t tGappedBlock) Alias() string { return t.tGappedBlockDo.Alias() }

func (t tGappedBlock) Columns(cols ...field.Expr) gen.Columns {
	return t.tGappedBlockDo.Columns(cols...)
}

func (t *tGappedBlock) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tGappedBlock) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["chainid"] = t.Chainid
	t.fieldMap["appid"] = t.Appid
	t.fieldMap["block_number"] = t.BlockNumber
	t.fieldMap["update_timestamp"] = t.UpdateTimestamp
	t.fieldMap["insert_timestamp"] = t.InsertTimestamp
	t.fieldMap["is_processed"] = t.IsProcessed
}

func (t tGappedBlock) clone(db *gorm.DB) tGappedBlock {
	t.tGappedBlockDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tGappedBlock) replaceDB(db *gorm.DB) tGappedBlock {
	t.tGappedBlockDo.ReplaceDB(db)
	return t
}

type tGappedBlockDo struct{ gen.DO }

type ITGappedBlockDo interface {
	gen.SubQuery
	Debug() ITGappedBlockDo
	WithContext(ctx context.Context) ITGappedBlockDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITGappedBlockDo
	WriteDB() ITGappedBlockDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITGappedBlockDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITGappedBlockDo
	Not(conds ...gen.Condition) ITGappedBlockDo
	Or(conds ...gen.Condition) ITGappedBlockDo
	Select(conds ...field.Expr) ITGappedBlockDo
	Where(conds ...gen.Condition) ITGappedBlockDo
	Order(conds ...field.Expr) ITGappedBlockDo
	Distinct(cols ...field.Expr) ITGappedBlockDo
	Omit(cols ...field.Expr) ITGappedBlockDo
	Join(table schema.Tabler, on ...field.Expr) ITGappedBlockDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITGappedBlockDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITGappedBlockDo
	Group(cols ...field.Expr) ITGappedBlockDo
	Having(conds ...gen.Condition) ITGappedBlockDo
	Limit(limit int) ITGappedBlockDo
	Offset(offset int) ITGappedBlockDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITGappedBlockDo
	Unscoped() ITGappedBlockDo
	Create(values ...*model.TGappedBlock) error
	CreateInBatches(values []*model.TGappedBlock, batchSize int) error
	Save(values ...*model.TGappedBlock) error
	First() (*model.TGappedBlock, error)
	Take() (*model.TGappedBlock, error)
	Last() (*model.TGappedBlock, error)
	Find() ([]*model.TGappedBlock, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TGappedBlock, err error)
	FindInBatches(result *[]*model.TGappedBlock, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TGappedBlock) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITGappedBlockDo
	Assign(attrs ...field.AssignExpr) ITGappedBlockDo
	Joins(fields ...field.RelationField) ITGappedBlockDo
	Preload(fields ...field.RelationField) ITGappedBlockDo
	FirstOrInit() (*model.TGappedBlock, error)
	FirstOrCreate() (*model.TGappedBlock, error)
	FindByPage(offset int, limit int) (result []*model.TGappedBlock, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITGappedBlockDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tGappedBlockDo) Debug() ITGappedBlockDo {
	return t.withDO(t.DO.Debug())
}

func (t tGappedBlockDo) WithContext(ctx context.Context) ITGappedBlockDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tGappedBlockDo) ReadDB() ITGappedBlockDo {
	return t.Clauses(dbresolver.Read)
}

func (t tGappedBlockDo) WriteDB() ITGappedBlockDo {
	return t.Clauses(dbresolver.Write)
}

func (t tGappedBlockDo) Session(config *gorm.Session) ITGappedBlockDo {
	return t.withDO(t.DO.Session(config))
}

func (t tGappedBlockDo) Clauses(conds ...clause.Expression) ITGappedBlockDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tGappedBlockDo) Returning(value interface{}, columns ...string) ITGappedBlockDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tGappedBlockDo) Not(conds ...gen.Condition) ITGappedBlockDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tGappedBlockDo) Or(conds ...gen.Condition) ITGappedBlockDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tGappedBlockDo) Select(conds ...field.Expr) ITGappedBlockDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tGappedBlockDo) Where(conds ...gen.Condition) ITGappedBlockDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tGappedBlockDo) Order(conds ...field.Expr) ITGappedBlockDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tGappedBlockDo) Distinct(cols ...field.Expr) ITGappedBlockDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tGappedBlockDo) Omit(cols ...field.Expr) ITGappedBlockDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tGappedBlockDo) Join(table schema.Tabler, on ...field.Expr) ITGappedBlockDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tGappedBlockDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITGappedBlockDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tGappedBlockDo) RightJoin(table schema.Tabler, on ...field.Expr) ITGappedBlockDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tGappedBlockDo) Group(cols ...field.Expr) ITGappedBlockDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tGappedBlockDo) Having(conds ...gen.Condition) ITGappedBlockDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tGappedBlockDo) Limit(limit int) ITGappedBlockDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tGappedBlockDo) Offset(offset int) ITGappedBlockDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tGappedBlockDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITGappedBlockDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tGappedBlockDo) Unscoped() ITGappedBlockDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tGappedBlockDo) Create(values ...*model.TGappedBlock) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tGappedBlockDo) CreateInBatches(values []*model.TGappedBlock, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tGappedBlockDo) Save(values ...*model.TGappedBlock) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tGappedBlockDo) First() (*model.TGappedBlock, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TGappedBlock), nil
	}
}

func (t tGappedBlockDo) Take() (*model.TGappedBlock, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TGappedBlock), nil
	}
}

func (t tGappedBlockDo) Last() (*model.TGappedBlock, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TGappedBlock), nil
	}
}

func (t tGappedBlockDo) Find() ([]*model.TGappedBlock, error) {
	result, err := t.DO.Find()
	return result.([]*model.TGappedBlock), err
}

func (t tGappedBlockDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TGappedBlock, err error) {
	buf := make([]*model.TGappedBlock, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tGappedBlockDo) FindInBatches(result *[]*model.TGappedBlock, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tGappedBlockDo) Attrs(attrs ...field.AssignExpr) ITGappedBlockDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tGappedBlockDo) Assign(attrs ...field.AssignExpr) ITGappedBlockDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tGappedBlockDo) Joins(fields ...field.RelationField) ITGappedBlockDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tGappedBlockDo) Preload(fields ...field.RelationField) ITGappedBlockDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tGappedBlockDo) FirstOrInit() (*model.TGappedBlock, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TGappedBlock), nil
	}
}

func (t tGappedBlockDo) FirstOrCreate() (*model.TGappedBlock, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TGappedBlock), nil
	}
}

func (t tGappedBlockDo) FindByPage(offset int, limit int) (result []*model.TGappedBlock, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tGappedBlockDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tGappedBlockDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tGappedBlockDo) Delete(models ...*model.TGappedBlock) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tGappedBlockDo) withDO(do gen.Dao) *tGappedBlockDo {
	t.DO = *do.(*gen.DO)
	return t
}
