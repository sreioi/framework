package cache

import (
	"github.com/sreioi/framework/contracts/foundation"
)

const Binding = "cache"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		config := app.MakeConfig()
		log := app.MakeLog()
		store := config.GetString("cache.default")

		return NewApplication(config, log, store)
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
}
