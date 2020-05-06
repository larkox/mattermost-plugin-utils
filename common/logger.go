package common

type LogContext map[string]interface{}

type Logger interface {
	With(LogContext) Logger
	Timed() Logger
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}
