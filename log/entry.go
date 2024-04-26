package log

import (
	"context"
	"github.com/sreioi/framework/contracts/log"
	"time"
)

type Entry struct {
	ctx     context.Context
	level   log.Level
	time    time.Time
	message string
}

func (r *Entry) Context() context.Context {
	return r.ctx
}

func (r *Entry) Level() log.Level {
	return r.level
}

func (r *Entry) Time() time.Time {
	return r.time
}

func (r *Entry) Message() string {
	return r.message
}
