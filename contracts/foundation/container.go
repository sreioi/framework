package foundation

import (
	"github.com/sreioi/framework/contracts/cache"
	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/contracts/database/orm"
	"github.com/sreioi/framework/contracts/filesystem"
	"github.com/sreioi/framework/contracts/http"
	"github.com/sreioi/framework/contracts/log"
	"github.com/sreioi/framework/contracts/route"
	"github.com/sreioi/framework/contracts/validation"
)

type Container interface {
	// Bind registers a binding with the container.
	Bind(key any, callback func(app Application) (any, error))
	// BindWith registers a binding with the container.
	BindWith(key any, callback func(app Application, parameters map[string]any) (any, error))
	// Instance registers an existing instance as shared in the container.
	Instance(key, instance any)
	// Singleton registers a shared binding in the container.
	Singleton(key any, callback func(app Application) (any, error))
	// Make resolves the given type from the container.
	Make(key any) (any, error)
	// MakeConfig resolves the config instance.
	MakeConfig() config.Config
	// MakeLog resolves the log instance.
	MakeLog() log.Log
	// MakeOrm resolves the orm instance.
	MakeOrm() orm.Orm
	// MakeRoute resolves the route instance.
	MakeRoute() route.Route
	// MakeStorage resolves the storage instance.
	MakeStorage() filesystem.Storage
	// MakeValidation resolves the validation instance.
	MakeValidation() validation.Validation
	// MakeView resolves the view instance.
	MakeView() http.View
	// MakeCache resolves the cache instance.
	MakeCache() cache.Cache
	// MakeRateLimiter resolves the rate limiter instance.
	MakeRateLimiter() http.RateLimiter
}
