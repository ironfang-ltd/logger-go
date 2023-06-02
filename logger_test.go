package logger

import (
	"bytes"
	"testing"
)

func TestSilent(t *testing.T) {

	buffer := &bytes.Buffer{}
	sink := NewBufferSink(buffer)

	l := New()
	l.SetLogLevel(Silent)
	l.SetSink(sink)
	l.Info("this is an info log")

	result := buffer.String()
	expected := ""

	if result != expected {
		t.Errorf("log result failed, expected: %s, got: %s", expected, result)
	}
}

func TestInfo(t *testing.T) {

	buffer := &bytes.Buffer{}
	sink := NewBufferSink(buffer)

	l := New()
	l.SetSink(sink)
	l.Info("this is an info log")

	result := buffer.String()
	expected := "this is an info log"

	if result != expected {
		t.Errorf("log result failed, expected: %s, got: %s", expected, result)
	}
}

func TestWarn(t *testing.T) {

	buffer := &bytes.Buffer{}
	sink := NewBufferSink(buffer)

	l := New()
	l.SetSink(sink)
	l.Warn("this is an warn log")

	result := buffer.String()
	expected := "this is an warn log"

	if result != expected {
		t.Errorf("log result failed, expected: %s, got: %s", expected, result)
	}
}

func TestError(t *testing.T) {

	buffer := &bytes.Buffer{}
	sink := NewBufferSink(buffer)

	l := New()
	l.SetSink(sink)
	l.Error("this is an error log")

	result := buffer.String()
	expected := "this is an error log"

	if result != expected {
		t.Errorf("log result failed, expected: %s, got: %s", expected, result)
	}
}
