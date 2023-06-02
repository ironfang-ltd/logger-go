package logger

import (
	"bytes"
	"fmt"
)

type bufferSink struct {
	buffer *bytes.Buffer
}

func (s *bufferSink) Send(format string, v ...any) {
	s.buffer.WriteString(fmt.Sprintf(format, v...))
}

func NewBufferSink(buffer *bytes.Buffer) LogSink {
	return &bufferSink{
		buffer: buffer,
	}
}
