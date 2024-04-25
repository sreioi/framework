package foundation

import "github.com/sreioi/framework/contracts/config"

type Container interface {
	// Bind registers a binding with the container.
	Bind(key any, callback func(app Application) (any, error))
	// BindWith registers a binding with the container.
	BindWith(key any, callback func(app Application, parameters map[string]any) (any, error))
	// Instance registers an existing instance as shared in the container.
	Instance(key, instance any)
	// Make resolves the given type from the container.
	Make(key any) (any, error)
	// MakeConfig resolves the config instance.
	MakeConfig() config.Config
	// Singleton registers a shared binding in the container.
	Singleton(key any, callback func(app Application) (any, error))
}
