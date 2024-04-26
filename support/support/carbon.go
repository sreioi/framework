package carbon

import (
	"github.com/golang-module/carbon/v2"
	"sync"
)

var Caron carbon.Carbon

var once sync.Once

func init() {
	once.Do(func() {
		carbon.SetDefault(carbon.Default{
			Layout:       carbon.DateTimeLayout,
			Timezone:     carbon.PRC,
			WeekStartsAt: carbon.Sunday,
			Locale:       "zh-CN",
		})
		Caron = carbon.NewCarbon()
	})
}

func NewTime() carbon.Carbon {
	return Caron
}
