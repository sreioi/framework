package foundation

import (
	"github.com/gookit/color"
	"github.com/sreioi/framework/config"
	Conconfig "github.com/sreioi/framework/contracts/config"
	Conlog "github.com/sreioi/framework/contracts/log"
	"github.com/sreioi/framework/log"
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
