package database

import (
	"fmt"
	"github.com/sreioi/framework/contracts/foundation"
	"golang.org/x/net/context"
)

const BindingOrm = "orm"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(BindingOrm, func(app foundation.Application) (any, error) {
		config := app.MakeConfig()
		defaultCon := config.GetString("database.default")

		orm, err := InitOrm(context.Background(), config, defaultCon)
		if err != nil {
			return nil, fmt.Errorf("[Orm] Init %s connection error: %v", defaultCon, err)
		}
		return orm, nil
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {

}
