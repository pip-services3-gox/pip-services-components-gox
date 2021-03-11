package log

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/info"
)

/*
Abstract logger that captures and formats log messages. Child classes take the captured messages and write them to their specific destinations.

Configuration parameters
Parameters to pass to the configure method for component configuration:

  level: maximum log level to capture
  source: source (context) name
References
*:context-info:*:*:1.0 (optional) ContextInfo to detect the context id and specify counters source
*/
type ILogWriter interface {
	Write(level int, correlationId string, err error, message string)
}

type Logger struct {
	level  int
	source string
	writer ILogWriter
}

// Creates a new instance of the logger and inherite from ILogerWriter.
// Parameters:
//   - writer ILogWriter
//   inherite from
// Returns *Logger
func InheritLogger(writer ILogWriter) *Logger {
	return &Logger{
		level:  Info,
		source: "",
		writer: writer,
	}
}

// Gets the maximum log level. Messages with higher log level are filtered out.
// Returns int
// the maximum log level.
func (c *Logger) Level() int {
	return c.level
}

// Set the maximum log level.
// Parameters:
//   - value int
// a new maximum log level.
func (c *Logger) SetLevel(value int) {
	c.level = value
}

// Gets the source (context) name.
// Returns string
// the source (context) name.
func (c *Logger) Source() string {
	return c.source
}

// Sets the source (context) name.
// Parameters:
//   - value string
//   a new source (context) name.
func (c *Logger) SetSource(value string) {
	c.source = value
}

// Configures component by passing configuration parameters.
// Parameters:
//   - config ConfigParams
//   configuration parameters to be set.
func (c *Logger) Configure(cfg *config.ConfigParams) {
	c.level = LogLevelConverter.ToLogLevel(cfg.GetAsStringWithDefault("level", strconv.Itoa(c.level)))
	c.source = cfg.GetAsStringWithDefault("source", c.source)
}

// Sets references to dependent components.
// Parameters:
//   - references IReferences
//   references to locate the component dependencies.
func (c *Logger) SetReferences(references refer.IReferences) {
	contextInfo, ok := references.GetOneOptional(
		refer.NewDescriptor("pip-services", "context-info", "*", "*", "1.0")).(info.ContextInfo)
	if ok && c.source == "" {
		c.source = contextInfo.Name
	}
}

// Composes an human-readable error description
// Parameters:
//   - err error
//   an error to format.
// Returns string
// a human-reable error description.
func (c *Logger) ComposeError(err error) string {
	builder := strings.Builder{}

	appErr, ok := err.(*errors.ApplicationError)
	if ok {
		builder.WriteString(appErr.Message)
		if appErr.Cause != "" {
			builder.WriteString(" Caused by: ")
			builder.WriteString(appErr.Cause)
		}
		if appErr.StackTrace != "" {
			builder.WriteString(" Stack trace: ")
			builder.WriteString(appErr.StackTrace)
		}
	} else {
		builder.WriteString(err.Error())
	}

	return builder.String()
}

// Formats the log message and writes it to the logger destination.
// Parameters:
//   - level int
//   a log level.
//   - correlationId: string
//    transaction id to trace execution through call chain.
//   - err error
//   an error object associated with this message.
//   - message string
//   a human-readable message to log.
//   - args []interface{}
//   arguments to parameterize the message.
func (c *Logger) FormatAndWrite(level int, correlationId string, err error, message string, args []interface{}) {
	if args != nil && len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	if c.writer != nil {
		c.writer.Write(level, correlationId, err, message)
	}
}

// Logs a message at specified log level.
// Parameters:
//   - level int
//   a log level.
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - err error
//   an error object associated with this message.
//   - message string
//   a human-readable message to log.
//   - args ...interface{}
//   arguments to parameterize the message.
func (c *Logger) Log(level int, correlationId string, err error, message string, args ...interface{}) {
	c.FormatAndWrite(level, correlationId, err, message, args)
}

// Logs fatal (unrecoverable) message that caused the process to crash.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - err error
//   an error object associated with this message.
//   - message string
//   a human-readable message to log.
//   - args ...interface{}
//   arguments to parameterize the message.
func (c *Logger) Fatal(correlationId string, err error, message string, args ...interface{}) {
	c.FormatAndWrite(Fatal, correlationId, err, message, args)
}

// Logs recoverable application error.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - err error
//   an error object associated with this message.
//   - message string
//   a human-readable message to log.
//   - args ...interface{}
//   arguments to parameterize the message.
func (c *Logger) Error(correlationId string, err error, message string, args ...interface{}) {
	c.FormatAndWrite(Error, correlationId, err, message, args)
}

// Logs a warning that may or may not have a negative impact.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - message string
//   a human-readable message to log.
//   - args ...interface{}
//   arguments to parameterize the message
func (c *Logger) Warn(correlationId string, message string, args ...interface{}) {
	c.FormatAndWrite(Warn, correlationId, nil, message, args)
}

// Logs an important information message
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - message string
//   a human-readable message to log.
//   - args ...interface{}
//   arguments to parameterize the message
func (c *Logger) Info(correlationId string, message string, args ...interface{}) {
	c.FormatAndWrite(Info, correlationId, nil, message, args)
}

// Logs a high-level debug information for troubleshooting.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - message string
//   a human-readable message to log.
//   - args ...interface{}
//   arguments to parameterize the message
func (c *Logger) Debug(correlationId string, message string, args ...interface{}) {
	c.FormatAndWrite(Debug, correlationId, nil, message, args)
}

// Logs a low-level debug information for troubleshooting.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - message string
//   a human-readable message to log.
//   - args ...interface{}
//   arguments to parameterize the message
func (c *Logger) Trace(correlationId string, message string, args ...interface{}) {
	c.FormatAndWrite(Trace, correlationId, nil, message, args)
}
