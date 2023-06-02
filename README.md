# Logger

Simple logger written in Go

#### Example

```go
package main

import "github.com/ironfang-ltd/logger-go"

func main() {
	l := logger.New()

	// Error,Warn,Info
	l.SetLogLevel(logger.Silent)

	// Optionally set the sink, defaults to console (stdout)
	l.SetSink(logger.NewConsoleSink())

	// Write some logs
	l.Info("this is an info log")
	l.Warn("this is an warn log")
	l.Error("this is an error log")
}
```
