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

func newTEventProcessedBlock(db *gorm.DB, opts ...gen.DOOption) tEventProcessedBlock {
	_tEventProcessedBlock := tEventProcessedBlock{}

	_tEventProcessedBlock.tEventProcessedBlockDo.UseDB(db, opts...)
	_tEventProcessedBlock.tEventProcessedBlockDo.UseModel(&model.TEventProcessedBlock{})

	tableName := _tEventProcessedBlock.tEventProcessedBlockDo.TableName()
	_tEventProcessedBlock.ALL = field.NewAsterisk(tableName)
	_tEventProcessedBlock.Chainid = field.NewInt32(tableName, "chainid")
	_tEventProcessedBlock.Appid = field.NewInt32(tableName, "appid")
	_tEventProcessedBlock.UpdateTimestamp = field.NewTime(tableName, "update_timestamp")
	_tEventProcessedBlock.InsertTimestamp = field.NewTime(tableName, "insert_timestamp")
	_tEventProcessedBlock.BlockNumber = field.NewInt64(tableName, "block_number")
	_tEventProcessedBlock.LatestBlockNumber = field.NewInt64(tableName, "latest_block_number")
	_tEventProcessedBlock.BacktrackBlockNumber = field.NewInt64(tableName, "backtrack_block_number")

	_tEventProcessedBlock.fillFieldMap()

	return _tEventProcessedBlock
}

type tEventProcessedBlock struct {
	tEventProcessedBlockDo tEventProcessedBlockDo

	ALL                  field.Asterisk
	Chainid              field.Int32
	Appid                field.Int32
	UpdateTimestamp      field.Time
	InsertTimestamp      field.Time
	BlockNumber          field.Int64
	LatestBlockNumber    field.Int64
	BacktrackBlockNumber field.Int64

	fieldMap map[string]field.Expr
}

func (t tEventProcessedBlock) Table(newTableName string) *tEventProcessedBlock {
	t.tEventProcessedBlockDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tEventProcessedBlock) As(alias string) *tEventProcessedBlock {
	t.tEventProcessedBlockDo.DO = *(t.tEventProcessedBlockDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tEventProcessedBlock) updateTableName(table string) *tEventProcessedBlock {
	t.ALL = field.NewAsterisk(table)
	t.Chainid = field.NewInt32(table, "chainid")
	t.Appid = field.NewInt32(table, "appid")
	t.UpdateTimestamp = field.NewTime(table, "update_timestamp")
	t.InsertTimestamp = field.NewTime(table, "insert_timestamp")
	t.BlockNumber = field.NewInt64(table, "block_number")
	t.LatestBlockNumber = field.NewInt64(table, "latest_block_number")
	t.BacktrackBlockNumber = field.NewInt64(table, "backtrack_block_number")

	t.fillFieldMap()

	return t
}

func (t *tEventProcessedBlock) WithContext(ctx context.Context) ITEventProcessedBlockDo {
	return t.tEventProcessedBlockDo.WithContext(ctx)
}

func (t tEventProcessedBlock) TableName() string { return t.tEventProcessedBlockDo.TableName() }

func (t tEventProcessedBlock) Alias() string { return t.tEventProcessedBlockDo.Alias() }

func (t tEventProcessedBlock) Columns(cols ...field.Expr) gen.Columns {
	return t.tEventProcessedBlockDo.Columns(cols...)
}

func (t *tEventProcessedBlock) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tEventProcessedBlock) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["chainid"] = t.Chainid
	t.fieldMap["appid"] = t.Appid
	t.fieldMap["update_timestamp"] = t.UpdateTimestamp
	t.fieldMap["insert_timestamp"] = t.InsertTimestamp
	t.fieldMap["block_number"] = t.BlockNumber
	t.fieldMap["latest_block_number"] = t.LatestBlockNumber
	t.fieldMap["backtrack_block_number"] = t.BacktrackBlockNumber
}

func (t tEventProcessedBlock) clone(db *gorm.DB) tEventProcessedBlock {
	t.tEventProcessedBlockDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tEventProcessedBlock) replaceDB(db *gorm.DB) tEventProcessedBlock {
	t.tEventProcessedBlockDo.ReplaceDB(db)
	return t
}

type tEventProcessedBlockDo struct{ gen.DO }

type ITEventProcessedBlockDo interface {
	gen.SubQuery
	Debug() ITEventProcessedBlockDo
	WithContext(ctx context.Context) ITEventProcessedBlockDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITEventProcessedBlockDo
	WriteDB() ITEventProcessedBlockDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITEventProcessedBlockDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITEventProcessedBlockDo
	Not(conds ...gen.Condition) ITEventProcessedBlockDo
	Or(conds ...gen.Condition) ITEventProcessedBlockDo
	Select(conds ...field.Expr) ITEventProcessedBlockDo
	Where(conds ...gen.Condition) ITEventProcessedBlockDo
	Order(conds ...field.Expr) ITEventProcessedBlockDo
	Distinct(cols ...field.Expr) ITEventProcessedBlockDo
	Omit(cols ...field.Expr) ITEventProcessedBlockDo
	Join(table schema.Tabler, on ...field.Expr) ITEventProcessedBlockDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITEventProcessedBlockDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITEventProcessedBlockDo
	Group(cols ...field.Expr) ITEventProcessedBlockDo
	Having(conds ...gen.Condition) ITEventProcessedBlockDo
	Limit(limit int) ITEventProcessedBlockDo
	Offset(offset int) ITEventProcessedBlockDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITEventProcessedBlockDo
	Unscoped() ITEventProcessedBlockDo
	Create(values ...*model.TEventProcessedBlock) error
	CreateInBatches(values []*model.TEventProcessedBlock, batchSize int) error
	Save(values ...*model.TEventProcessedBlock) error
	First() (*model.TEventProcessedBlock, error)
	Take() (*model.TEventProcessedBlock, error)
	Last() (*model.TEventProcessedBlock, error)
	Find() ([]*model.TEventProcessedBlock, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TEventProcessedBlock, err error)
	FindInBatches(result *[]*model.TEventProcessedBlock, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TEventProcessedBlock) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITEventProcessedBlockDo
	Assign(attrs ...field.AssignExpr) ITEventProcessedBlockDo
	Joins(fields ...field.RelationField) ITEventProcessedBlockDo
	Preload(fields ...field.RelationField) ITEventProcessedBlockDo
	FirstOrInit() (*model.TEventProcessedBlock, error)
	FirstOrCreate() (*model.TEventProcessedBlock, error)
	FindByPage(offset int, limit int) (result []*model.TEventProcessedBlock, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITEventProcessedBlockDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tEventProcessedBlockDo) Debug() ITEventProcessedBlockDo {
	return t.withDO(t.DO.Debug())
}

func (t tEventProcessedBlockDo) WithContext(ctx context.Context) ITEventProcessedBlockDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tEventProcessedBlockDo) ReadDB() ITEventProcessedBlockDo {
	return t.Clauses(dbresolver.Read)
}

func (t tEventProcessedBlockDo) WriteDB() ITEventProcessedBlockDo {
	return t.Clauses(dbresolver.Write)
}

func (t tEventProcessedBlockDo) Session(config *gorm.Session) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Session(config))
}

func (t tEventProcessedBlockDo) Clauses(conds ...clause.Expression) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tEventProcessedBlockDo) Returning(value interface{}, columns ...string) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tEventProcessedBlockDo) Not(conds ...gen.Condition) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tEventProcessedBlockDo) Or(conds ...gen.Condition) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tEventProcessedBlockDo) Select(conds ...field.Expr) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tEventProcessedBlockDo) Where(conds ...gen.Condition) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tEventProcessedBlockDo) Order(conds ...field.Expr) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tEventProcessedBlockDo) Distinct(cols ...field.Expr) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tEventProcessedBlockDo) Omit(cols ...field.Expr) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tEventProcessedBlockDo) Join(table schema.Tabler, on ...field.Expr) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tEventProcessedBlockDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITEventProcessedBlockDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tEventProcessedBlockDo) RightJoin(table schema.Tabler, on ...field.Expr) ITEventProcessedBlockDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tEventProcessedBlockDo) Group(cols ...field.Expr) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tEventProcessedBlockDo) Having(conds ...gen.Condition) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tEventProcessedBlockDo) Limit(limit int) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tEventProcessedBlockDo) Offset(offset int) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tEventProcessedBlockDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tEventProcessedBlockDo) Unscoped() ITEventProcessedBlockDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tEventProcessedBlockDo) Create(values ...*model.TEventProcessedBlock) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tEventProcessedBlockDo) CreateInBatches(values []*model.TEventProcessedBlock, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tEventProcessedBlockDo) Save(values ...*model.TEventProcessedBlock) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tEventProcessedBlockDo) First() (*model.TEventProcessedBlock, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TEventProcessedBlock), nil
	}
}

func (t tEventProcessedBlockDo) Take() (*model.TEventProcessedBlock, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TEventProcessedBlock), nil
	}
}

func (t tEventProcessedBlockDo) Last() (*model.TEventProcessedBlock, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TEventProcessedBlock), nil
	}
}

func (t tEventProcessedBlockDo) Find() ([]*model.TEventProcessedBlock, error) {
	result, err := t.DO.Find()
	return result.([]*model.TEventProcessedBlock), err
}

func (t tEventProcessedBlockDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TEventProcessedBlock, err error) {
	buf := make([]*model.TEventProcessedBlock, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tEventProcessedBlockDo) FindInBatches(result *[]*model.TEventProcessedBlock, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tEventProcessedBlockDo) Attrs(attrs ...field.AssignExpr) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tEventProcessedBlockDo) Assign(attrs ...field.AssignExpr) ITEventProcessedBlockDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tEventProcessedBlockDo) Joins(fields ...field.RelationField) ITEventProcessedBlockDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tEventProcessedBlockDo) Preload(fields ...field.RelationField) ITEventProcessedBlockDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tEventProcessedBlockDo) FirstOrInit() (*model.TEventProcessedBlock, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TEventProcessedBlock), nil
	}
}

func (t tEventProcessedBlockDo) FirstOrCreate() (*model.TEventProcessedBlock, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TEventProcessedBlock), nil
	}
}

func (t tEventProcessedBlockDo) FindByPage(offset int, limit int) (result []*model.TEventProcessedBlock, count int64, err error) {
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

func (t tEventProcessedBlockDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tEventProcessedBlockDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tEventProcessedBlockDo) Delete(models ...*model.TEventProcessedBlock) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tEventProcessedBlockDo) withDO(do gen.Dao) *tEventProcessedBlockDo {
	t.DO = *do.(*gen.DO)
	return t
}
