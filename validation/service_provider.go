package validation

import (
	"github.com/sreioi/framework/contracts/foundation"
)

const Binding = "validation"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewValidation(), nil
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {

}
