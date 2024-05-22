package filesystem

import (
	configcontract "github.com/sreioi/framework/contracts/config"
	filesystemcontract "github.com/sreioi/framework/contracts/filesystem"
	"github.com/sreioi/framework/contracts/foundation"
)

const Binding = "filesystem"

var ConfigFacade configcontract.Config
var StorageFacade filesystemcontract.Storage

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewStorage(app.MakeConfig()), nil
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	StorageFacade = app.MakeStorage()
}
