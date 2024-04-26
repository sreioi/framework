package orm

type Observer interface {
	// Retrieved 从数据库检索模型后调用。
	Retrieved(Event) error
	// Creating 在创建模型时被调用。
	Creating(Event) error
	// Created 在创建模型后调用。
	Created(Event) error
	// Updating 在更新模型时调用。
	Updating(Event) error
	// Updated 在模型更新后调用。
	Updated(Event) error
	// Saving 在保存模型时调用。
	Saving(Event) error
	// Saved 在保存模型后调用。
	Saved(Event) error
	// Deleting 在删除模型时调用。
	Deleting(Event) error
	// Deleted 在删除模型后调用。
	Deleted(Event) error
	// ForceDeleting 在强制删除模型时调用。
	ForceDeleting(Event) error
	// ForceDeleted 在强制删除模型后调用。
	ForceDeleted(Event) error
}
