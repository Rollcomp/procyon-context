package context

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type LoggerProvider interface {
	GetLogger() Logger
}

type LogLevel uint32

const (
	PanicLevel LogLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type Logger interface {
	Trace(ctx Context, args ...interface{})
	Debug(ctx Context, args ...interface{})
	Info(ctx Context, args ...interface{})
	Print(ctx Context, args ...interface{})
	Warning(ctx Context, args ...interface{})
	Error(ctx Context, args ...interface{})
	Fatal(ctx Context, args ...interface{})
	Panic(ctx Context, args ...interface{})
}

type SimpleLogger struct {
	log *logrus.Logger
}

func NewSimpleLogger() *SimpleLogger {
	log := &SimpleLogger{
		&logrus.Logger{
			Out:       os.Stdout,
			Formatter: NewLogFormatter(),
			Level:     logrus.InfoLevel,
		},
	}
	return log
}

func (l *SimpleLogger) checkContext(ctx Context) {
	if ctx == nil {
		panic("Context must not be nil")
	}
}

func (l *SimpleLogger) Trace(ctx Context, args ...interface{}) {
	l.checkContext(ctx)
	l.log.WithFields(logrus.Fields{
		"context_id": ctx.GetContextId(),
	}).Trace(args...)
}

func (l *SimpleLogger) Debug(ctx Context, args ...interface{}) {
	l.checkContext(ctx)
	l.log.WithFields(logrus.Fields{
		"context_id": ctx.GetContextId(),
	}).Debug(args...)
}

func (l *SimpleLogger) Info(ctx Context, args ...interface{}) {
	l.checkContext(ctx)
	l.log.WithFields(logrus.Fields{
		"context_id": ctx.GetContextId(),
	}).Info(args...)
}

func (l *SimpleLogger) Print(ctx Context, args ...interface{}) {
	l.checkContext(ctx)
	l.log.WithFields(logrus.Fields{
		"context_id": ctx.GetContextId(),
	}).Print(args...)
}

func (l SimpleLogger) Warning(ctx Context, args ...interface{}) {
	l.checkContext(ctx)
	l.log.WithFields(logrus.Fields{
		"context_id": ctx.GetContextId(),
	}).Warning(args...)
}

func (l *SimpleLogger) Error(ctx Context, args ...interface{}) {
	l.checkContext(ctx)
	l.log.WithFields(logrus.Fields{
		"context_id": ctx.GetContextId(),
	}).Error(args...)
}

func (l *SimpleLogger) Fatal(ctx Context, args ...interface{}) {
	l.checkContext(ctx)
	l.log.WithFields(logrus.Fields{
		"context_id": ctx.GetContextId(),
	}).Fatal(args...)
}

func (l *SimpleLogger) Panic(ctx Context, args ...interface{}) {
	l.checkContext(ctx)
	l.log.WithFields(logrus.Fields{
		"context_id": ctx.GetContextId(),
	}).Panic(args...)
}

type LogFormatter struct {
	logrus.TextFormatter
}

func NewLogFormatter() *LogFormatter {
	formatter := &LogFormatter{}
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	return formatter
}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 37 // gray
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}

	logContextId := entry.Data["context_id"].(string)
	separatorIndex := strings.Index(logContextId, "-")
	logContextId = logContextId[:separatorIndex]

	return []byte(
		fmt.Sprintf("[%s] \x1b[%dm%-7s\x1b[0m %s : %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), logContextId, entry.Message)), nil
}