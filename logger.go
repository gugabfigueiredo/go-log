package log

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
)

var (
	ErrTagsOddCount     = errors.New("odd logger tag count")
	ErrTagsKeyNotString = errors.New("tag key is not a string")
)

type ILogger interface {
	With(tags ...any) ILogger
	Debug(message string, tags ...any)
	Fatal(message string, tags ...any)
	Error(message string, tags ...any)
	Warn(message string, tags ...any)
	Info(message string, tags ...any)
	Log(message string, tags ...any)
}

type Logger struct {
	*zerolog.Logger
	context map[string]any
}

func (l *Logger) with(tags ...any) (*Logger, error) {

	if len(tags) == 0 {
		return l, nil
	}

	if len(tags)%2 == 1 {
		return l, ErrTagsOddCount
	}

	logger := &Logger{
		Logger:  l.Logger,
		context: l.context,
	}

	for i := 0; i < len(tags); i += 2 {
		tag, ok := tags[i].(string)
		if !ok {
			return l, ErrTagsKeyNotString
		}
		logger.context[tag] = tags[i+1]
	}

	return logger, nil
}

func (l *Logger) logWith(e *zerolog.Event, message string, tags ...any) {

	logger, err := l.with(tags...)
	if err != nil {
		l.Error("unable to add tags to log line", "error", err, "tags", tags, "originalMessage", message)
	}

	for key, value := range logger.context {
		e = e.Str(key, fmt.Sprintf("%v", value))
	}

	e.Msg(message)
}

func (l *Logger) With(tags ...any) ILogger {
	logger, err := l.with(tags...)
	if err != nil {
		l.Error("unable to add tags to log line", "error", err, "tags", tags)
	}
	return logger
}

func (l *Logger) Debug(message string, tags ...any) {
	l.logWith(l.Logger.Debug(), message, tags...)
}

func (l *Logger) Fatal(message string, tags ...any) {
	l.logWith(l.Logger.Fatal(), message, tags...)
}

func (l *Logger) Error(message string, tags ...any) {
	l.logWith(l.Logger.Error(), message, tags...)
}

func (l *Logger) Warn(message string, tags ...any) {
	l.logWith(l.Logger.Warn(), message, tags...)
}

func (l *Logger) Log(message string, tags ...any) {
	l.logWith(l.Logger.Log(), message, tags...)
}

func (l *Logger) Info(message string, tags ...any) {
	l.logWith(l.Logger.Info(), message, tags...)
}
