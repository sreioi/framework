package config

import (
	"github.com/sreioi/framework/contracts/foundation"
	"github.com/sreioi/framework/support"
)

const Binding = "config"

type ServiceProvider struct {
}

func (conf *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(support.EnvPath), nil
	})
}

func (conf *ServiceProvider) Boot(app foundation.Application) {

}
