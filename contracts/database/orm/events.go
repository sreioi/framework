package orm

import "context"

type EventType string

const EventRetrieved EventType = "retrieved"
const EventCreating EventType = "creating"
const EventCreated EventType = "created"
const EventUpdating EventType = "updating"
const EventUpdated EventType = "Updated"
const EventSaving EventType = "saving"
const EventSaved EventType = "saved"
const EventDeleting EventType = "deleting"
const EventDeleted EventType = "deleted"
const EventForceDeleting EventType = "force_deleting"
const EventForceDeleted EventType = "force_deleted"

type Event interface {
	// Context 返回事件上下文。
	Context() context.Context
	// GetAttribute 返回给定键的属性值。
	GetAttribute(key string) any
	// GetOriginal 返回给定键的原始属性值。
	GetOriginal(key string, def ...any) any
	// IsDirty 如果给定列为 dirty，则返回 true。
	IsDirty(columns ...string) bool
	// IsClean 如果给定列是干净的，则返回 true。
	IsClean(columns ...string) bool
	// Query 返回查询实例。
	Query() Query
	// SetAttribute 设置给定键的属性值。
	SetAttribute(key string, value any)
}

type DispatchesEvents interface {
	// DispatchesEvents returns the event handlers.
	DispatchesEvents() map[EventType]func(Event) error
}
