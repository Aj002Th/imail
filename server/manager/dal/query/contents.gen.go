// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/Aj002Th/imail/server/manager/dal/model"
)

func newContent(db *gorm.DB, opts ...gen.DOOption) content {
	_content := content{}

	_content.contentDo.UseDB(db, opts...)
	_content.contentDo.UseModel(&model.Content{})

	tableName := _content.contentDo.TableName()
	_content.ALL = field.NewAsterisk(tableName)
	_content.ID = field.NewUint(tableName, "id")
	_content.CreatedAt = field.NewTime(tableName, "created_at")
	_content.UpdatedAt = field.NewTime(tableName, "updated_at")
	_content.DeletedAt = field.NewField(tableName, "deleted_at")
	_content.Title = field.NewString(tableName, "title")
	_content.Time = field.NewTime(tableName, "time")
	_content.Description = field.NewString(tableName, "description")
	_content.Cover = field.NewString(tableName, "cover")
	_content.Link = field.NewString(tableName, "link")
	_content.Author = field.NewString(tableName, "author")
	_content.Source = field.NewString(tableName, "source")
	_content.Category = field.NewString(tableName, "category")
	_content.Sended = field.NewBool(tableName, "sended")

	_content.fillFieldMap()

	return _content
}

type content struct {
	contentDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Title       field.String
	Time        field.Time
	Description field.String
	Cover       field.String
	Link        field.String
	Author      field.String
	Source      field.String
	Category    field.String
	Sended      field.Bool

	fieldMap map[string]field.Expr
}

func (c content) Table(newTableName string) *content {
	c.contentDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c content) As(alias string) *content {
	c.contentDo.DO = *(c.contentDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *content) updateTableName(table string) *content {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Title = field.NewString(table, "title")
	c.Time = field.NewTime(table, "time")
	c.Description = field.NewString(table, "description")
	c.Cover = field.NewString(table, "cover")
	c.Link = field.NewString(table, "link")
	c.Author = field.NewString(table, "author")
	c.Source = field.NewString(table, "source")
	c.Category = field.NewString(table, "category")
	c.Sended = field.NewBool(table, "sended")

	c.fillFieldMap()

	return c
}

func (c *content) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *content) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 13)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["title"] = c.Title
	c.fieldMap["time"] = c.Time
	c.fieldMap["description"] = c.Description
	c.fieldMap["cover"] = c.Cover
	c.fieldMap["link"] = c.Link
	c.fieldMap["author"] = c.Author
	c.fieldMap["source"] = c.Source
	c.fieldMap["category"] = c.Category
	c.fieldMap["sended"] = c.Sended
}

func (c content) clone(db *gorm.DB) content {
	c.contentDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c content) replaceDB(db *gorm.DB) content {
	c.contentDo.ReplaceDB(db)
	return c
}

type contentDo struct{ gen.DO }

type IContentDo interface {
	gen.SubQuery
	Debug() IContentDo
	WithContext(ctx context.Context) IContentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IContentDo
	WriteDB() IContentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IContentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IContentDo
	Not(conds ...gen.Condition) IContentDo
	Or(conds ...gen.Condition) IContentDo
	Select(conds ...field.Expr) IContentDo
	Where(conds ...gen.Condition) IContentDo
	Order(conds ...field.Expr) IContentDo
	Distinct(cols ...field.Expr) IContentDo
	Omit(cols ...field.Expr) IContentDo
	Join(table schema.Tabler, on ...field.Expr) IContentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IContentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IContentDo
	Group(cols ...field.Expr) IContentDo
	Having(conds ...gen.Condition) IContentDo
	Limit(limit int) IContentDo
	Offset(offset int) IContentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IContentDo
	Unscoped() IContentDo
	Create(values ...*model.Content) error
	CreateInBatches(values []*model.Content, batchSize int) error
	Save(values ...*model.Content) error
	First() (*model.Content, error)
	Take() (*model.Content, error)
	Last() (*model.Content, error)
	Find() ([]*model.Content, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Content, err error)
	FindInBatches(result *[]*model.Content, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Content) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IContentDo
	Assign(attrs ...field.AssignExpr) IContentDo
	Joins(fields ...field.RelationField) IContentDo
	Preload(fields ...field.RelationField) IContentDo
	FirstOrInit() (*model.Content, error)
	FirstOrCreate() (*model.Content, error)
	FindByPage(offset int, limit int) (result []*model.Content, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IContentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FindBySourceAuthorLink(source string, author string, link string) (result model.Content, err error)
}

// SELECT * FROM @@table WHERE source = @source and author = @author and link = @link and deleted_at is null
func (c contentDo) FindBySourceAuthorLink(source string, author string, link string) (result model.Content, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, source)
	params = append(params, author)
	params = append(params, link)
	generateSQL.WriteString("SELECT * FROM contents WHERE source = ? and author = ? and link = ? and deleted_at is null ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c contentDo) Debug() IContentDo {
	return c.withDO(c.DO.Debug())
}

func (c contentDo) WithContext(ctx context.Context) IContentDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c contentDo) ReadDB() IContentDo {
	return c.Clauses(dbresolver.Read)
}

func (c contentDo) WriteDB() IContentDo {
	return c.Clauses(dbresolver.Write)
}

func (c contentDo) Session(config *gorm.Session) IContentDo {
	return c.withDO(c.DO.Session(config))
}

func (c contentDo) Clauses(conds ...clause.Expression) IContentDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c contentDo) Returning(value interface{}, columns ...string) IContentDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c contentDo) Not(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c contentDo) Or(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c contentDo) Select(conds ...field.Expr) IContentDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c contentDo) Where(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c contentDo) Order(conds ...field.Expr) IContentDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c contentDo) Distinct(cols ...field.Expr) IContentDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c contentDo) Omit(cols ...field.Expr) IContentDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c contentDo) Join(table schema.Tabler, on ...field.Expr) IContentDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c contentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IContentDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c contentDo) RightJoin(table schema.Tabler, on ...field.Expr) IContentDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c contentDo) Group(cols ...field.Expr) IContentDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c contentDo) Having(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c contentDo) Limit(limit int) IContentDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c contentDo) Offset(offset int) IContentDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c contentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IContentDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c contentDo) Unscoped() IContentDo {
	return c.withDO(c.DO.Unscoped())
}

func (c contentDo) Create(values ...*model.Content) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c contentDo) CreateInBatches(values []*model.Content, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c contentDo) Save(values ...*model.Content) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c contentDo) First() (*model.Content, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) Take() (*model.Content, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) Last() (*model.Content, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) Find() ([]*model.Content, error) {
	result, err := c.DO.Find()
	return result.([]*model.Content), err
}

func (c contentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Content, err error) {
	buf := make([]*model.Content, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c contentDo) FindInBatches(result *[]*model.Content, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c contentDo) Attrs(attrs ...field.AssignExpr) IContentDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c contentDo) Assign(attrs ...field.AssignExpr) IContentDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c contentDo) Joins(fields ...field.RelationField) IContentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c contentDo) Preload(fields ...field.RelationField) IContentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c contentDo) FirstOrInit() (*model.Content, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) FirstOrCreate() (*model.Content, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) FindByPage(offset int, limit int) (result []*model.Content, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c contentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c contentDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c contentDo) Delete(models ...*model.Content) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *contentDo) withDO(do gen.Dao) *contentDo {
	c.DO = *do.(*gen.DO)
	return c
}
