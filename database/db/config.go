package db

import (
	"fmt"
	"github.com/google/wire"

	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/contracts/database"
	"github.com/sreioi/framework/contracts/database/orm"
)

var ConfigSet = wire.NewSet(NewConfigImpl, wire.Bind(new(Config), new(*ConfigImpl)))
var _ Config = &ConfigImpl{}

type Config interface {
	Reads() []database.Config
	Writes() []database.Config
}

type ConfigImpl struct {
	config     config.Config
	connection string
}

func NewConfigImpl(config config.Config, connection string) *ConfigImpl {
	return &ConfigImpl{
		config:     config,
		connection: connection,
	}
}

func (c *ConfigImpl) Reads() []database.Config {
	configs := c.config.Get(fmt.Sprintf("database.connections.%s.read", c.connection))
	if configs, ok := configs.([]database.Config); ok {
		return c.fillDefault(configs)
	}
	return []database.Config{}
}

func (c *ConfigImpl) Writes() []database.Config {
	configs := c.config.Get(fmt.Sprintf("database.connections.%s.write", c.connection))
	if configs, ok := configs.([]database.Config); ok {
		return c.fillDefault(configs)
	}
	return []database.Config{}
}

func (c *ConfigImpl) fillDefault(configs []database.Config) []database.Config {
	var newConfigs []database.Config
	driver := c.config.GetString(fmt.Sprintf("database.connections.%s.driver", c.connection))
	for _, item := range configs {
		if driver != orm.DriverSqlite.String() {
			if item.Host == "" {
				item.Host = c.config.GetString(fmt.Sprintf("database.connections.%s.host", c.connection))
			}
			if item.Port == 0 {
				item.Port = c.config.GetInt(fmt.Sprintf("database.connections.%s.port", c.connection))
			}
			if item.Username == "" {
				item.Username = c.config.GetString(fmt.Sprintf("database.connections.%s.username", c.connection))
			}
			if item.Password == "" {
				item.Password = c.config.GetString(fmt.Sprintf("database.connections.%s.password", c.connection))
			}
		}
		if item.Database == "" {
			item.Database = c.config.GetString(fmt.Sprintf("database.connections.%s.database", c.connection))
		}
		newConfigs = append(newConfigs, item)
	}
	return newConfigs
}
