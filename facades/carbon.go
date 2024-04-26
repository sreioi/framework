package facades

import (
	carbon2 "github.com/golang-module/carbon/v2"
	"github.com/sreioi/framework/support/carbon"
)

func Time() carbon2.Carbon {
	return carbon.NewTime()
}
