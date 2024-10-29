package log

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/google/uuid"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	splitter = "core"

	defaultLevel = 4
)

const (
	ContextKey string = "log_fields"
)

type Level string

const LoggerInCtx = "logger_ctx"

const (
	LevelUnknown = Level("unknown")
	LevelDebug   = Level("debug")
	LevelInfo    = Level("info")
	LevelWarn    = Level("warn")
	LevelError   = Level("error")
	LevelPanic   = Level("panic")
)

var levels = map[string]Level{
	LevelDebug.String(): LevelDebug,
	LevelInfo.String():  LevelInfo,
	LevelWarn.String():  LevelWarn,
	LevelError.String(): LevelError,
	LevelPanic.String(): LevelPanic,
}

func UniqueStringIdentifier() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}

func ParseAndSetLevel(level string) Level {
	if l, ok := levels[level]; ok {
		return l
	}
	return LevelInfo
}

func (l Level) String() string {
	return string(l)
}

type Hook interface {
	logrus.Hook
}

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Panic(format string, args ...interface{})

	Debugf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	WithError(err error) Entry
	WithField(key string, value interface{}) Entry
	WithFields(fields map[string]interface{}) Entry

	SetLevel(level Level)
	SetWriter(w io.Writer)
	AddHook(h Hook)
}

type logger struct {
	l          *logrus.Logger
	stackLevel int
}

func New(level Level, stackLevel int) Logger {
	return newWithStackLevel(level, stackLevel)
}

func newWithStackLevel(level Level, stackLevel int) *logger {
	l := &logger{
		l:          logrus.New(),
		stackLevel: stackLevel,
	}
	l.l.Formatter = &logrus.JSONFormatter{}

	l.SetLevel(level)

	mw := io.MultiWriter(os.Stdout)
	l.l.SetOutput(mw)

	return l
}

func (l *logger) log(level Level, format string, args ...interface{}) {
	e := &entry{
		l,
		logrus.NewEntry(l.l),
		-1,
	}

	e.log(level, format, args...)
}

func (l *logger) Debug(format string, args ...interface{}) {
	l.log(LevelDebug, format, args...)
}

func (l *logger) Info(format string, args ...interface{}) {
	l.log(LevelInfo, format, args...)
}

func (l *logger) Warning(format string, args ...interface{}) {
	l.log(LevelWarn, format, args...)
}

func (l *logger) Error(format string, args ...interface{}) {
	l.log(LevelError, format, args...)
}

func (l *logger) Panic(format string, args ...interface{}) {
	l.log(LevelPanic, format, args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.log(LevelDebug, format, args...)
}

func (l *logger) Warningf(format string, args ...interface{}) {
	l.log(LevelWarn, format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.log(LevelError, format, args...)
}

func (l *logger) WithError(err error) Entry {
	return &entry{
		l,
		l.l.WithError(err),
		2,
	}
}

func (l *logger) WithField(key string, value interface{}) Entry {
	return &entry{
		l,
		l.l.WithField(key, value),
		2,
	}
}

func (l *logger) WithFields(fields map[string]interface{}) Entry {
	return &entry{
		l,
		l.l.WithFields(fields),
		2,
	}
}

func (l *logger) SetWriter(w io.Writer) {
	l.l.Out = w
}

func (l *logger) Writer() *io.PipeWriter {
	return l.l.Writer()
}

func (l *logger) SetLevel(logLevel Level) {
	level, err := logrus.ParseLevel(logLevel.String())
	if err == nil {
		l.l.SetLevel(level)
		return
	}
	l.Panic(err.Error())
}

func (l *logger) AddHook(h Hook) {
	l.l.AddHook(h)
}

type ContextData struct {
	Fields map[string]interface{}
}

func FromContext(ctx context.Context) ContextData {
	if data, ok := ctx.Value(ContextKey).(ContextData); ok {
		return data
	}
	return ContextData{}
}

func WithContextData(ctx context.Context, data ContextData) context.Context {
	// nolint:golint
	return context.WithValue(ctx, ContextKey, data)
}

func LoggerInContext(ctx context.Context) Entry {
	if data, ok := ctx.Value(LoggerInCtx).(Entry); ok {
		return data
	}
	return WithField("REQUEST_ID", utils.GenerateUUID())
}
