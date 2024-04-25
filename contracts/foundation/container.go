package foundation

type Container interface {
	// Bind registers a binding with the container.
	Bind(key any, callback func(app Application) (any, error))
	// Instance registers an existing instance as shared in the container.
	Instance(key, instance any)
	// Make resolves the given type from the container.
	Make(key any) (any, error)
	// Singleton registers a shared binding in the container.
	Singleton(key any, callback func(app Application) (any, error))
}
