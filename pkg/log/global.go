package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

var Std = newWithStackLevel(LevelInfo, defaultLevel)

func Debug(format string, args ...interface{}) {
	Std.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	Std.Info(format, args...)
}

func Warning(format string, args ...interface{}) {
	Std.Warning(format, args...)
}

func Error(format string, args ...interface{}) {
	Std.Error(format, args...)
}

func Panic(format string, args ...interface{}) {
	Std.Panic(format, args...)
}

func Debugf(format string, args ...interface{}) {
	Std.Debug(format, args...)
}

func Warningf(format string, args ...interface{}) {
	Std.Warning(format, args...)
}

func Errorf(format string, args ...interface{}) {
	Std.Error(format, args...)
}

func WithField(key string, value interface{}) Entry {
	return Std.WithField(key, value)
}

func WithFields(fields map[string]interface{}) Entry {
	return Std.WithFields(fields)
}

func SetLevel(logLevel Level) {
	Std.SetLevel(logLevel)
}

func GetLevel() Level {
	switch Std.l.Level {
	case logrus.DebugLevel:
		return LevelDebug
	case logrus.InfoLevel:
		return LevelInfo
	case logrus.WarnLevel:
		return LevelWarn
	case logrus.ErrorLevel:
		return LevelError
	case logrus.PanicLevel:
		return LevelPanic
	}
	return LevelUnknown
}

func AddHook(h Hook) {
	Std.AddHook(h)
}

func SetWriter(w io.Writer) {
	Std.SetWriter(w)
}

func Writer() *io.PipeWriter {
	return Std.Writer()
}
