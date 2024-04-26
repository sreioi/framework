package facades

import "github.com/sreioi/framework/contracts/log"

func Log() log.Log {
	return App().MakeLog()
}
