package foundation

import (
	"github.com/gookit/color"
	"github.com/sreioi/framework/config"
	Conconfig "github.com/sreioi/framework/contracts/config"
)

func (c *Container) MakeConfig() Conconfig.Config {
	instance, err := c.Make(config.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}
	return instance.(Conconfig.Config)
}
