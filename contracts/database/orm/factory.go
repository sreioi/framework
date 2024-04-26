package orm

type Factory interface {
	// Count 设置应生成的模型数量。
	Count(count int) Factory
	// Create 创建一个模型并将其持久化到数据库中。
	Create(value any, attributes ...map[string]any) error
	// CreateQuietly 会创建一个模型并将其持久化到数据库中，而不会触发任何模型事件。
	CreateQuietly(value any, attributes ...map[string]any) error
	// Make 会创建一个模型并返回，但不会将其持久化到数据库中。
	Make(value any, attributes ...map[string]any) error
}
