package facades

import (
	"github.com/golang-module/carbon/v2"
	CarbonTime "github.com/sreioi/framework/support/support"
)

func Time() carbon.Carbon {
	return CarbonTime.NewTime()
}
