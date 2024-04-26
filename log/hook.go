package log

import (
	"github.com/sirupsen/logrus"
	"github.com/sreioi/framework/contracts/log"
)

type Hook struct {
	instance log.Hook
}

func (h *Hook) Levels() []logrus.Level {
	levels := h.instance.Levels()
	var logrusLevels []logrus.Level
	for _, item := range levels {
		logrusLevels = append(logrusLevels, logrus.Level(item))
	}

	return logrusLevels
}

func (h *Hook) Fire(entry *logrus.Entry) error {
	return h.instance.Fire(&Entry{
		ctx:     entry.Context,
		level:   log.Level(entry.Level),
		time:    entry.Time,
		message: entry.Message,
	})
}
