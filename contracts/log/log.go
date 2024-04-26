package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Level logrus.Level

const (
	StackDriver  = "stack"
	SingleDriver = "single"
	DailyDriver  = "daily"
	CustomDriver = "custom"
)

type Log interface {
	// WithContext adds a context to the logger.
	WithContext(ctx context.Context) Writer
	Writer
}

type Writer interface {
	// Debug logs a message at DebugLevel.
	Debug(args ...any)
	// Debugf is equivalent to Debug, but with support for fmt.Printf-style arguments.
	Debugf(format string, args ...any)
	// Info logs a message at InfoLevel.
	Info(args ...any)
	// Infof is equivalent to Info, but with support for fmt.Printf-style arguments.
	Infof(format string, args ...any)
	// Warning logs a message at WarningLevel.
	Warning(args ...any)
	// Warningf is equivalent to Warning, but with support for fmt.Printf-style arguments.
	Warningf(format string, args ...any)
	// Error logs a message at ErrorLevel.
	Error(args ...any)
	// Errorf is equivalent to Error, but with support for fmt.Printf-style arguments.
	Errorf(format string, args ...any)
	// Fatal logs a message at FatalLevel.
	Fatal(args ...any)
	// Fatalf is equivalent to Fatal, but with support for fmt.Printf-style arguments.
	Fatalf(format string, args ...any)
	// Panic logs a message at PanicLevel.
	Panic(args ...any)
	// Panicf is equivalent to Panic, but with support for fmt.Printf-style arguments.
	Panicf(format string, args ...any)
}

type Logger interface {
	// Handle pass a channel config path here
	Handle(channel string) (Hook, error)
}

type Hook interface {
	// Levels monitoring level
	Levels() []Level
	// Fire executes logic when trigger
	Fire(Entry) error
}

type Entry interface {
	// Context returns the context of the entry.
	Context() context.Context
	// Level returns the level of the entry.
	Level() Level
	// Time returns the timestamp of the entry.
	Time() time.Time
	// Message returns the message of the entry.
	Message() string
}
