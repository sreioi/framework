package orm

import (
	"context"
	"database/sql"
)

type Orm interface {
	// Connection 从连接池中获取 Orm 实例。
	Connection(name string) Orm
	// DB 获取底层数据库连接。
	DB() (*sql.DB, error)
	// Query 获取一个新的查询创建器实例。
	Query() Query
	// Factory 为给定的模型名称获取一个新的工厂实例。
	Factory() Factory
	// Observe 向 Orm 注册观察员。
	Observe(model any, observer Observer)
	// Transaction 运行一个包裹在数据库事务中的回调。
	Transaction(txFunc func(tx Transaction) error) error
	// WithContext 设置 Orm 使用的上下文。
	WithContext(ctx context.Context) Orm
}

type Transaction interface {
	Query
	// Commit 提交事务。
	Commit() error
	// Rollback 回滚事务。
	Rollback() error
}

type Query interface {
	// Association 通过名称获取关联实例。
	Association(association string) Association
	// Begin 开始事务。
	Begin() (Transaction, error)
	// Driver 获取驱动程序。
	Driver() Driver
	// Count 获取记录数。
	Count(count *int64) error
	// Connection 获取一个新的查询创建器实例。
	Connection(databaseName string) Query
	// Create 创建记录。
	Create(value any) error
	// Cursor 返回一个游标，使用 scan() 历返回的记录。
	Cursor() (chan Cursor, error)
	// Delete 删除记录。
	Delete(value any, conds ...any) (*Result, error)
	// Distinct 指定要查询的不同字段。
	Distinct(args ...any) Query
	// Exec 执行原生 SQL
	Exec(sql string, values ...any) (*Result, error)
	// Find 查找符合给定条件的记录。
	Find(dest any, conds ...any) error
	// FindOrFail 查找符合给定条件的记录，如果找不到则抛出异常。
	FindOrFail(dest any, conds ...any) error
	// First 查找符合给定条件的第一个记录。
	First(dest any) error
	// FirstOrCreate 查找符合给定条件的第一个记录，如果找不到则创建。
	FirstOrCreate(dest any, conds ...any) error
	// FirstOr 查找符合给定条件的第一个记录，如果找不到则执行回调。
	FirstOr(dest any, callback func() error) error
	// FirstOrFail 查找符合给定条件的第一个记录，如果找不到则抛出异常。
	FirstOrFail(dest any) error
	// FirstOrNew 查找符合给定条件的第一个记录，如果找不到则返回一个用这些属性初始化的模型新实例。
	FirstOrNew(dest any, attributes any, values ...any) error
	// ForceDelete 强制删除记录。
	ForceDelete(value any, conds ...any) (*Result, error)
	// Get 检索数据库中的所有记录。
	Get(dest any) error
	// Group 分组指定查询的分组方法。
	Group(name string) Query
	// Having 指定查询的 HAVING 条件。
	Having(query any, args ...any) Query
	// Join 指定查询的 JOIN 条件。
	Join(query string, args ...any) Query
	// Limit 返回的数据条数。
	Limit(limit int) Query
	// Load 为模型加载关系。
	Load(dest any, relation string, args ...any) error
	// LoadMissing 为模型加载一个尚未加载的关系。
	LoadMissing(dest any, relation string, args ...any) error
	// LockForUpdate 锁定表格中的选定行，以便更新。
	// 它在底层 SQL 查询中使用 "FOR UPDATE"
	// SELECT * FROM TABLE WHERE ... FOR UPDATE 是一种SQL语句，用于锁定行，以便其他事务不能修改这些行。
	LockForUpdate() Query
	// Model 设置要查询的模型实例。
	Model(value any) Query
	// Offset 指定在开始返回记录之前要跳过的记录数。
	Offset(offset int) Query
	// Omit 指定查询中应省略的列。
	Omit(columns ...string) Query
	// Order 指定返回结果的顺序。
	Order(value any) Query
	// OrWhere 在查询中添加 "or where "子句。
	OrWhere(query any, args ...any) Query
	// Paginate 将给定的查询转化为一个简单的分页器。
	Paginate(page, limit int, dest any, total *int64) error
	// Pluck 从数据库中检索单列。
	Pluck(column string, dest any) error
	// Raw 创建原始查询。
	Raw(sql string, values ...any) Query
	// Save 更新数据库中的值
	Save(value any) error
	// SaveQuietly 在不触发事件的情况下更新数据库中的值
	SaveQuietly(value any) error
	// Scan 扫描查询结果并填充目标对象。
	Scan(dest any) error
	// Scopes 应用一个或多个查询范围。
	Scopes(funcs ...func(Query) Query) Query
	// Select 指定应从数据库中检索的字段
	Select(query any, args ...any) Query
	// SharedLock 锁定表格中的选定行。
	// 它在底层 SQL 查询中使用 "FOR SHARE"
	// SELECT * FROM TABLE WHERE ... FOR SHARE 是一种 SELECT 语句，它用于锁定所选行以进行读取，防止其他事务对其进行修改。
	SharedLock() Query
	// Sum 计算列值的总和，并填充目标对象。
	Sum(column string, dest any) error
	// Table 指定用于查询的表。
	Table(name string, args ...any) Query
	// Update 用给定列和值更新记录
	Update(column any, value ...any) (*Result, error)
	// UpdateOrCreate 查找符合给定属性的第一条记录,如果没有找到，则创建一个具有这些属性的新记录。
	UpdateOrCreate(dest any, attributes any, values any) error
	// Where 在查询中添加 "where "子句。
	Where(query any, args ...any) Query
	// WithoutEvents 禁用查询的事件触发。
	WithoutEvents() Query
	// WithTrashed 允许将软删除模型纳入结果中。
	WithTrashed() Query
	// With 返回一个新的查询实例，其中已急切加载给定的关系。
	With(query string, args ...any) Query
}

type Association interface {
	// Find 查找符合给定条件的记录。
	Find(out any, conds ...any) error
	// Append 将模型附加到关联中。
	Append(values ...any) error
	// Replace 用给定值替换关联。
	Replace(values ...any) error
	// Delete 会从关联中删除给定值。
	Delete(values ...any) error
	// Clear 清除关联。
	Clear() error
	// Count 返回关联中的记录数。
	Count() int64
}

type ConnectionModel interface {
	// Connection 获取模型的连接名称。
	Connection() string
}

type Cursor interface {
	// Scan 将当前行扫描到给定的目标行。
	Scan(value any) error
}

type Result struct {
	// RowsAffected 返回受影响的行数。
	RowsAffected int64
}
