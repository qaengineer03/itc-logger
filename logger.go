package logger

import (
	"context"
	"log"
)

type Logger interface {
	Info(ctx context.Context, msg string)
	Error(ctx context.Context, err error)
	Warning(ctx context.Context, msg string)
	Debug(ctx context.Context, msg string)
}

type loggerImpl struct {
	logLevel LoggingLevel
}

func New(logLevel LoggingLevel) Logger {
	return &loggerImpl{
		logLevel: logLevel,
	}
}

const userHeader = "X-Username"

func (l *loggerImpl) Info(ctx context.Context, msg string) {
	if l.logLevel < InfoLevel {
		return
	}

	log.Println(ctx.Value(userHeader), msg)
}

func (l *loggerImpl) Error(ctx context.Context, err error) {
	if l.logLevel < ErrorLevel {
		return
	}

	log.Println(ctx.Value(userHeader), err)
}

func (l *loggerImpl) Warning(ctx context.Context, msg string) {
	if l.logLevel < WarningLevel {
		return
	}

	log.Println(ctx.Value(userHeader), msg)
}

func (l *loggerImpl) Debug(ctx context.Context, msg string) {
	if l.logLevel < DebugLevel {
		return
	}

	log.Println(ctx.Value(userHeader), msg)
}
