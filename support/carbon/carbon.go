package carbon

import (
	"github.com/golang-module/carbon/v2"
	"sync"
)

var Carbon carbon.Carbon

var once = sync.Once{}

func init() {
	once.Do(func() {
		carbon.SetDefault(carbon.Default{
			Layout:       carbon.DateTimeLayout,
			Timezone:     carbon.Shanghai,
			WeekStartsAt: carbon.Sunday,
			Locale:       "zh-CN",
		})
		Carbon = carbon.NewCarbon()
	})
}

func NewTime() carbon.Carbon {
	return Carbon
}
