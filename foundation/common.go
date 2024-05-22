package foundation

import (
	"github.com/gookit/color"
	"github.com/sreioi/framework/cache"
	"github.com/sreioi/framework/config"
	cachecontract "github.com/sreioi/framework/contracts/cache"
	Conconfig "github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/contracts/database/orm"
	filesystemcontract "github.com/sreioi/framework/contracts/filesystem"
	httpcontract "github.com/sreioi/framework/contracts/http"
	Conlog "github.com/sreioi/framework/contracts/log"
	"github.com/sreioi/framework/contracts/route"
	validationcontract "github.com/sreioi/framework/contracts/validation"
	"github.com/sreioi/framework/database"
	"github.com/sreioi/framework/filesystem"
	"github.com/sreioi/framework/http"
	"github.com/sreioi/framework/log"
	route2 "github.com/sreioi/framework/route"
	"github.com/sreioi/framework/validation"
)

func (c *Container) MakeConfig() Conconfig.Config {
	instance, err := c.Make(config.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}
	return instance.(Conconfig.Config)
}

func (c *Container) MakeLog() Conlog.Log {
	instance, err := c.Make(log.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}
	return instance.(Conlog.Log)
}

func (c *Container) MakeOrm() orm.Orm {
	instance, err := c.Make(database.BindingOrm)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(orm.Orm)
}

func (c *Container) MakeRoute() route.Route {
	instance, err := c.Make(route2.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(route.Route)
}

func (c *Container) MakeStorage() filesystemcontract.Storage {
	instance, err := c.Make(filesystem.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(filesystemcontract.Storage)
}

func (c *Container) MakeValidation() validationcontract.Validation {
	instance, err := c.Make(validation.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(validationcontract.Validation)
}

func (c *Container) MakeView() httpcontract.View {
	instance, err := c.Make(http.BindingView)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(httpcontract.View)
}

func (c *Container) MakeRateLimiter() httpcontract.RateLimiter {
	instance, err := c.Make(http.BindingRateLimiter)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(httpcontract.RateLimiter)
}

func (c *Container) MakeCache() cachecontract.Cache {
	instance, err := c.Make(cache.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(cachecontract.Cache)
}
