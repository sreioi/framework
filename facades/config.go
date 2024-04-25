package facades

import (
	"github.com/sreioi/framework/contracts/config"
)

func Config() config.Config {
	return App().MakeConfig()
}
