// logger is a package that logs based on five levels, Debug, Info, Warn, Error,
// and Fatal. Logrus is the underlying logger package. In production, this
// package logs in JSON format.
package logger

import (
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

// A Logger interface must support logging at Debug, Info, Warn, Error, and
// Fatal levels.
//
// Debug should only be used during development. If any debug calls are used, it
// should be disabled or removed before going to production.
//
// Info should log events that are happening, that may be critical for the
// application to succeed, but is not an error. For example, logging that a
// connection to a database was successfully established.
//
// Warn should almost never be used. It's like saying, "There may be an error
// here, but we're not sure, so deal with it in the future". It adds needless
// overhead and complexity.
//
// Error logs events that should not have happened, but did. Might include how a
// an operation received the wrong inputs, or data was checked against some
// validation measures, and found to be wrong.
//
// Fatal logs events that cannot be recovered from. Prefer Error over fatal,
// except in cases where the application fails if an event does error out,
// because most events can be handled.
type Logger interface {
	Debug(msg string, data map[string]string)
	Info(msg string, data map[string]string)
	Warn(msg string, data map[string]string)
	Error(msg string, data map[string]string)
	Fatal(msg string, data map[string]string)
}

type logger struct {
	l *logrus.Logger
}

// NewLogger returns a logger with JSON formatting set in production,
// the output streaming to STDOUT, and logging level set to Info level by
// default, unless in development environment.
// The output io.Writer sets where logrus logs to, nominally should be set to
// Stdout.
func NewLogger(debug bool, output io.Writer) Logger {
	l := &logrus.Logger{}

	if !debug {
		l.Formatter = &logrus.JSONFormatter{}
	} else {
		l.Formatter = &logrus.TextFormatter{}
	}

	l.Out = output

	if debug {
		l.Level = logrus.DebugLevel
	} else {
		l.Level = logrus.InfoLevel
	}

	return &logger{
		l: l,
	}
}

func (log *logger) formatFields(data map[string]string) logrus.Fields {
	// add time
	now, _ := time.Now().MarshalText()
	f := logrus.Fields{
		"logtime": string(now),
	}

	if data != nil {
		for k, v := range data {
			f[k] = v
		}
	}

	return f
}

func (l *logger) Debug(msg string, data map[string]string) {
	f := l.formatFields(data)
	l.l.WithFields(f).Debug(msg)
}

func (l *logger) Info(msg string, data map[string]string) {
	f := l.formatFields(data)
	l.l.WithFields(f).Info(msg)
}

func (l *logger) Warn(msg string, data map[string]string) {
	f := l.formatFields(data)
	l.l.WithFields(f).Warn(msg)
}

func (l *logger) Error(msg string, data map[string]string) {
	f := l.formatFields(data)
	l.l.WithFields(f).Error(msg)
}

func (l *logger) Fatal(msg string, data map[string]string) {
	f := l.formatFields(data)
	l.l.WithFields(f).Fatal(msg)
}
