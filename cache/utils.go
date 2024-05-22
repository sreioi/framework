package cache

import (
	"github.com/sreioi/framework/contracts/config"
)

func prefix(config config.Config) string {
	return config.GetString("cache.prefix") + ":"
}
