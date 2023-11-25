package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
)

var (
	ErrTagsOddCount     = errors.New("odd logger tag count")
	ErrTagsKeyNotString = errors.New("tag key is not a string")
)

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
		l.logWith(l.Error(), "unable to add tags to log line", "error", err, "tags", tags, "originalMessage", message)
	}

	for key, value := range logger.context {
		e = e.Str(key, fmt.Sprintf("%v", value))
	}

	e.Msg(message)
}

func (l *Logger) With(tags ...any) *Logger {
	logger, err := l.with(tags...)
	if err != nil {
		l.logWith(l.Error(), "unable to add tags to log line", "error", err, "tags", tags)
	}
	return logger
}

func (l *Logger) D(message string, tags ...any) {
	l.logWith(l.Debug(), message, tags...)
}

func (l *Logger) F(message string, tags ...any) {
	l.logWith(l.Fatal(), message, tags...)
}

func (l *Logger) E(message string, tags ...any) {
	l.logWith(l.Error(), message, tags...)
}

func (l *Logger) W(message string, tags ...any) {
	l.logWith(l.Warn(), message, tags...)
}

func (l *Logger) L(message string, tags ...any) {
	l.logWith(l.Log(), message, tags...)
}

func (l *Logger) I(message string, tags ...any) {
	l.logWith(l.Info(), message, tags...)
}
