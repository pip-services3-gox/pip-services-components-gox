package log

import (
	"context"
	"fmt"
	"strings"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/info"
)

// ILogWriter Abstract logger that captures and formats log messages.
// Child classes take the captured messages and write them to their specific destinations.
//	Configuration parameters to pass to the configure method for component configuration:
//		level: maximum log level to capture
//		source: source (context) name
//	References:
//		*:context-info:*:*:1.0 (optional) ContextInfo to detect the context id and specify counters source
type ILogWriter interface {
	Write(ctx context.Context, level LevelType, correlationId string, err error, message string)
}

type Logger struct {
	level  LevelType
	source string
	writer ILogWriter
}

// InheritLogger creates a new instance of the logger and inherit from ILogWriter.
//	Parameters: writer ILogWriter inherit from
//	Returns: *Logger
func InheritLogger(writer ILogWriter) *Logger {
	return &Logger{
		level:  LevelInfo,
		source: "",
		writer: writer,
	}
}

// Level gets the maximum log level. Messages with higher log level are filtered out.
//	Returns int the maximum log level.
func (c *Logger) Level() LevelType {
	return c.level
}

// SetLevel set the maximum log level.
//	Parameters: value int a new maximum log level.
func (c *Logger) SetLevel(value LevelType) {
	c.level = value
}

// Source gets the source (context) name.
//	Returns: string the source (context) name.
func (c *Logger) Source() string {
	return c.source
}

// SetSource sets the source (context) name.
//	Parameters: value string a new source (context) name.
func (c *Logger) SetSource(value string) {
	c.source = value
}

// Configure configures component by passing configuration parameters.
//	Parameters: config ConfigParams configuration parameters to be set.
func (c *Logger) Configure(cfg *config.ConfigParams) {
	c.level = LevelConverter.ToLogLevel(cfg.GetAsStringWithDefault("level", logLevelToString(c.level)))
	c.source = cfg.GetAsStringWithDefault("source", c.source)
}

// SetReferences to dependent components.
//	Parameters: references IReferences references to locate the component dependencies.
func (c *Logger) SetReferences(references refer.IReferences) {
	descr := refer.NewDescriptor("pip-services", "context-info", "*", "*", "1.0")
	if contextInfo, ok := references.GetOneOptional(descr).(info.ContextInfo); ok && c.source == "" {
		c.source = contextInfo.Name
	}
}

// ComposeError composes an human-readable error description
// Parameters: err error an error to format.
//	Returns string a human-readable error description.
func (c *Logger) ComposeError(err error) string {
	builder := strings.Builder{}

	if appErr, ok := err.(*errors.ApplicationError); ok {
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

// FormatAndWrite formats the log message and writes it to the logger destination.
// Parameters:
//		- ctx context.Context
//		- level LevelType a log level
//		- correlationId: string transaction id to trace execution through call chain
//		- err error an error object associated with this message
//		- message string a human-readable message to log
//		- args []any arguments to parameterize the message
func (c *Logger) FormatAndWrite(ctx context.Context, level LevelType,
	correlationId string, err error, message string, args []any) {

	if args != nil && len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	if c.writer != nil {
		c.writer.Write(ctx, level, correlationId, err, message)
	}
}

// Log a message at specified log level.
//	Parameters:
//		- ctx context.Context
//		- level LevelType a log level.
//		- correlationId string transaction id to trace execution through call chain.
//		- err error an error object associated with this message.
//		- message string a human-readable message to log.
//		- args ...any arguments to parameterize the message.
func (c *Logger) Log(ctx context.Context, level LevelType, correlationId string, err error, message string, args ...any) {
	c.FormatAndWrite(ctx, level, correlationId, err, message, args)
}

// Fatal logs fatal (unrecoverable) message that caused the process to crash.
//	Parameters:
//		- ctx context.Context
//		- correlationId string transaction id to trace execution through call chain.
//		- err error an error object associated with this message.
//		- message string a human-readable message to log.
//		- args ...any arguments to parameterize the message.
func (c *Logger) Fatal(ctx context.Context, correlationId string, err error, message string, args ...any) {
	c.FormatAndWrite(ctx, LevelFatal, correlationId, err, message, args)
}

// Logs recoverable application error.
//	Parameters:
//		- ctx context.Context
//		- correlationId string transaction id to trace execution through call chain.
//		- err error an error object associated with this message.
//		- message string a human-readable message to log.
//		- args ...any arguments to parameterize the message.
func (c *Logger) Error(ctx context.Context, correlationId string, err error, message string, args ...any) {
	c.FormatAndWrite(ctx, LevelError, correlationId, err, message, args)
}

// Warn logs a warning that may or may not have a negative impact.
//	Parameters:
//		- ctx context.Context
//		- correlationId string transaction id to trace execution through call chain.
//		- message string a human-readable message to log.
//		- args ...any arguments to parameterize the message.
func (c *Logger) Warn(ctx context.Context, correlationId string, message string, args ...any) {
	c.FormatAndWrite(ctx, LevelWarn, correlationId, nil, message, args)
}

// Info logs an important information message
//	Parameters:
//		- ctx context.Context
//		- correlationId string transaction id to trace execution through call chain.
//		- message string a human-readable message to log.
//		- args ...any arguments to parameterize the message.
func (c *Logger) Info(ctx context.Context, correlationId string, message string, args ...any) {
	c.FormatAndWrite(ctx, LevelInfo, correlationId, nil, message, args)
}

// Debug logs a high-level debug information for troubleshooting.
//	Parameters:
//		- ctx context.Context
//		- correlationId string transaction id to trace execution through call chain.
//		- message string a human-readable message to log.
//		- args ...any arguments to parameterize the message.
func (c *Logger) Debug(ctx context.Context, correlationId string, message string, args ...any) {
	c.FormatAndWrite(ctx, LevelDebug, correlationId, nil, message, args)
}

// Trace logs a low-level debug information for troubleshooting.
//	Parameters:
//		- ctx context.Context
//		- correlationId string transaction id to trace execution through call chain.
//		- message string a human-readable message to log.
//		- args ...any arguments to parameterize the message.
func (c *Logger) Trace(ctx context.Context, correlationId string, message string, args ...any) {
	c.FormatAndWrite(ctx, LevelTrace, correlationId, nil, message, args)
}
