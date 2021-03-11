package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
)

/*
Logger that writes log messages to console.

Errors are written to standard err stream and all other messages to standard out stream.

Configuration parameters
  level: maximum log level to capture
  source: source (context) name
References
*:context-info:*:*:1.0 (optional) ContextInfo to detect the context id and specify counters source
see
Logger

Example
  logger = NewConsoleLogger();
  logger.SetLevel(LogLevel.Debug);
  logger.Error("123", ex, "Error occured: %s", ex.message);
  logger.Debug("123", "Everything is OK.");
*/
type ConsoleLogger struct {
	Logger
}

// Creates a new instance of the logger.
// Returns ConsoleLogger
func NewConsoleLogger() *ConsoleLogger {
	c := &ConsoleLogger{}
	c.Logger = *InheritLogger(c)
	return c
}

// Writes a log message to the logger destination.
// Parameters:
//   - level int
//   a log level.
//   correlationId string
//   transaction id to trace execution through call chain.
//   - err error
//   an error object associated with this message.
//   - message string
//   a human-readable message to log.
func (c *ConsoleLogger) Write(level int, correlationId string, err error, message string) {
	if c.Level() < level {
		return
	}

	if correlationId == "" {
		correlationId = "---"
	}
	levelStr := LogLevelConverter.ToString(level)
	dateStr := convert.StringConverter.ToString(time.Now().UTC())

	build := strings.Builder{}
	build.WriteString("[")
	build.WriteString(correlationId)
	build.WriteString(":")
	build.WriteString(levelStr)
	build.WriteString(":")
	build.WriteString(dateStr)
	build.WriteString("] ")

	build.WriteString(message)

	if err != nil {
		if len(message) == 0 {
			build.WriteString("Error: ")
		} else {
			build.WriteString(": ")
		}

		build.WriteString(c.ComposeError(err))
	}

	build.WriteString("\n")
	output := build.String()

	if level == Fatal || level == Error || level == Warn {
		fmt.Fprintf(os.Stderr, output)
	} else {
		fmt.Printf(output)
	}
}
