package facades

import (
	conFoundation "github.com/sreioi/framework/contracts/foundation"
	"github.com/sreioi/framework/foundation"
)

func App() conFoundation.Application {
	return foundation.App
}
