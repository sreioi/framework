package http

import (
	"github.com/sreioi/framework/contracts/cache"
	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/contracts/foundation"
	"github.com/sreioi/framework/contracts/http"
)

const BindingRateLimiter = "rate_limiter"
const BindingView = "view"

type ServiceProvider struct{}

var (
	CacheFacade       cache.Cache
	ConfigFacade      config.Config
	RateLimiterFacade http.RateLimiter
)

func (http *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(BindingRateLimiter, func(app foundation.Application) (any, error) {
		return NewRateLimiter(), nil
	})
	app.Singleton(BindingView, func(app foundation.Application) (any, error) {
		return NewView(), nil
	})
}

func (http *ServiceProvider) Boot(app foundation.Application) {
	CacheFacade = app.MakeCache()
	ConfigFacade = app.MakeConfig()
	RateLimiterFacade = app.MakeRateLimiter()
}
