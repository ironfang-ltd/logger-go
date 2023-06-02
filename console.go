package logger

import (
	"log"
	"os"
)

type consoleSink struct {
	output *log.Logger
}

func (s *consoleSink) Send(format string, v ...any) {
	s.output.Printf(format, v...)
}

func NewConsoleSink() LogSink {
	return &consoleSink{
		output: log.New(os.Stdout, "", log.LstdFlags),
	}
}
