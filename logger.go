package logger

import (
	"sync"
)

type LogLevel int

const (
	Silent LogLevel = iota + 1
	Error
	Warn
	Info
)

type Logger interface {
	Info(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
	SetLogLevel(level LogLevel)
	SetSink(sink LogSink)
}

type LogSink interface {
	Send(format string, v ...any)
}

type logger struct {
	sink  LogSink
	level LogLevel
	mu    sync.Mutex
}

func (l *logger) Info(format string, v ...any) {
	if l.level >= Info {
		l.sink.Send(format, v...)
	}
}

func (l *logger) Warn(format string, v ...any) {
	if l.level >= Warn {
		l.sink.Send(format, v...)
	}
}

func (l *logger) Error(format string, v ...any) {
	if l.level >= Error {
		l.sink.Send(format, v...)
	}
}

func (l *logger) SetLogLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *logger) SetSink(sink LogSink) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.sink = sink
}

func New() Logger {
	logger := &logger{
		sink:  NewConsoleSink(),
		level: Info,
	}

	return logger
}
