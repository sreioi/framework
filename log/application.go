package log

import (
	"context"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/contracts/log"
)

type Application struct {
	instance *logrus.Logger
	log.Writer
}

var logsChannel = make(map[string]Application)

func NewApplication(config config.Config) *Application {
	//instance := logrus.New()
	//instance.SetLevel(logrus.DebugLevel)
	//
	//if config != nil {
	//	// 通过日志配置文件来初始化日志实例
	//	if logging := config.GetString("log.default"); logging != "" {
	//		if err := registerHook(config, instance, logging); err != nil {
	//			color.Redln("Init facades.Log error: " + err.Error())
	//			return nil
	//		}
	//	}
	//}
	//
	//return &Application{
	//	instance: instance,
	//	Writer:   NewWriter(instance.WithContext(context.Background())),
	//}

	initLogs(config)
	if logging := config.GetString("log.default"); logging != "" {
		if logger, ok := logsChannel[logging]; ok {
			return &logger
		}
	}
	return nil
}

func (r *Application) WithContext(ctx context.Context) log.Writer {
	switch r.Writer.(type) {
	case *Writer:
		return NewWriter(r.instance.WithContext(ctx))
	default:
		return r.Writer
	}
}

func (r *Application) Channel(channel string) *Application {
	logger, ok := logsChannel[channel]
	if !ok {
		color.Redln("Error log channel : " + channel)
		return nil
	}
	return &logger
}

func initLogs(config config.Config) {
	channels := config.Get("log.channels").(map[string]any)
	for channel, _ := range channels {
		if channel != "" {
			instance := logrus.New()
			instance.SetLevel(logrus.DebugLevel)

			if err := registerHook(config, instance, channel); err != nil {
				color.Redln("Init facades.Log error: " + err.Error())
			}

			logsChannel[channel] = Application{
				instance: instance,
				Writer:   NewWriter(instance.WithContext(context.Background())),
			}
		}
	}
}
