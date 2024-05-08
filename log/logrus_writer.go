package log

import (
	"errors"
	"fmt"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/contracts/log"
	"github.com/sreioi/framework/log/formatter"
	"github.com/sreioi/framework/log/logger"
	"io"
)

type Writer struct {
	instance *logrus.Entry

	code string

	message string
	// stacktrace
	stackEnabled bool
	stacktrace   map[string]any
}

func NewWriter(instance *logrus.Entry) log.Writer {
	return &Writer{
		instance: instance,
	}
}

func (r *Writer) Debug(args ...any) {
	r.instance.WithField("root", r.toMap()).Debug(args...)
}

func (r *Writer) Debugf(format string, args ...any) {
	r.instance.WithField("root", r.toMap()).Debugf(format, args...)
}

func (r *Writer) Info(args ...any) {
	r.instance.WithField("root", r.toMap()).Info(args...)
}

func (r *Writer) Infof(format string, args ...any) {
	r.instance.WithField("root", r.toMap()).Infof(format, args...)
}

func (r *Writer) Warning(args ...any) {
	r.instance.WithField("root", r.toMap()).Warning(args...)
}

func (r *Writer) Warningf(format string, args ...any) {
	r.instance.WithField("root", r.toMap()).Warningf(format, args...)
}

func (r *Writer) Error(args ...any) {
	r.withStackTrace(fmt.Sprint(args...))
	r.instance.WithField("root", r.toMap()).Error(args...)
}

func (r *Writer) Errorf(format string, args ...any) {
	r.withStackTrace(fmt.Sprintf(format, args...))
	r.instance.WithField("root", r.toMap()).Errorf(format, args...)
}

func (r *Writer) Fatal(args ...any) {
	r.withStackTrace(fmt.Sprint(args...))
	r.instance.WithField("root", r.toMap()).Fatal(args...)
}

func (r *Writer) Fatalf(format string, args ...any) {
	r.withStackTrace(fmt.Sprintf(format, args...))
	r.instance.WithField("root", r.toMap()).Fatalf(format, args...)
}

func (r *Writer) Panic(args ...any) {
	r.withStackTrace(fmt.Sprint(args...))
	r.instance.WithField("root", r.toMap()).Panic(args...)
}

func (r *Writer) Panicf(format string, args ...any) {
	r.withStackTrace(fmt.Sprintf(format, args...))
	r.instance.WithField("root", r.toMap()).Panicf(format, args...)
}

func (r *Writer) withStackTrace(message string) {
	erisNew := eris.New(message)
	r.message = erisNew.Error()
	format := eris.NewDefaultJSONFormat(eris.FormatOptions{
		InvertOutput: true,
		WithTrace:    true,
		InvertTrace:  true,
	})
	r.stacktrace = eris.ToCustomJSON(erisNew, format)
	r.stackEnabled = true
}

// ToMap returns a map representation of the error.
func (r *Writer) toMap() map[string]any {
	payload := map[string]any{}

	if message := r.message; message != "" {
		payload["message"] = message
	}

	if code := r.code; code != "" {
		payload["code"] = code
	}

	if stacktrace := r.stacktrace; stacktrace != nil || r.stackEnabled {
		payload["stacktrace"] = stacktrace
	}

	return payload
}

func registerHook(config config.Config, instance *logrus.Logger, channel string) error {
	channelPath := "log.channels." + channel
	driver := config.GetString(channelPath + ".driver")

	var hook logrus.Hook
	var err error

	switch driver {
	case log.StackDriver: // 配置的所有通道
		for _, stackChannel := range config.Get(channelPath + ".channels").([]string) {
			if stackChannel == channel {
				return errors.New("stack drive can't include self channel")
			}

			if err = registerHook(config, instance, stackChannel); err != nil {
				return err
			}
		}
		return nil
	case log.SingleDriver: // 单个日志文件
		// 判断是否打印到控制台
		if !config.GetBool(channelPath + ".print") {
			instance.SetOutput(io.Discard)
		}
		// 创建单个日志文件实例
		hook, err = logger.NewSingle(config).Handle(channelPath)
		if err != nil {
			return err
		}
	case log.DailyDriver: // 按天分割日志文件
		if !config.GetBool(channelPath + ".print") {
			instance.SetOutput(io.Discard)
		}
		hook, err = logger.NewDaily(config).Handle(channelPath)
		if err != nil {
			return err
		}
	case log.CustomDriver: // 自定义通道
		logLogger := config.Get(channelPath + ".diy").(log.Logger)
		logHook, err := logLogger.Handle(channelPath)
		if err != nil {
			return err
		}

		hook = &Hook{logHook}
	default:
		return errors.New("Error logging channel: " + channel)
	}

	instance.SetFormatter(formatter.NewGeneral(config))
	instance.AddHook(hook)

	return nil
}
