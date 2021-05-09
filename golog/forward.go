package golog

import "github.com/rs/zerolog"

type logForwarder struct{}

var defaultLogForwarder logForwarder

func LogForwarder() logForwarder {
	return defaultLogForwarder
}

func (t logForwarder) Trace() *zerolog.Event {
	return Trace()
}

func (t logForwarder) Debug() *zerolog.Event {
	return Debug()
}

func (t logForwarder) Info() *zerolog.Event {
	return Info()
}

func (t logForwarder) Warn() *zerolog.Event {
	return Warn()
}

func (t logForwarder) Error() *zerolog.Event {
	return Error()
}

func (t logForwarder) Fatal() *zerolog.Event {
	return Fatal()
}

func (t logForwarder) Printf(format string, args ...interface{}) {
	Printf(format, args...)
}
