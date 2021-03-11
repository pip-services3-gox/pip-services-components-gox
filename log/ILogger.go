package log

/*
Interface for logger components that capture execution log messages.
*/
type ILogger interface {

	// Gets the maximum log level. Messages with higher log level are filtered out.
	Level() int

	// Set the maximum log level.
	SetLevel(value int)

	// Logs a message at specified log level.
	Log(level int, correlationId string, err error, message string, args ...interface{})

	// Logs fatal (unrecoverable) message that caused the process to crash.
	Fatal(correlationId string, err error, message string, args ...interface{})

	// Logs recoverable application error.
	Error(correlationId string, err error, message string, args ...interface{})
	// Logs a warning that may or may not have a negative impact.
	Warn(correlationId string, message string, args ...interface{})

	// Logs an important information message
	Info(correlationId string, message string, args ...interface{})

	// Logs a high-level debug information for troubleshooting.
	Debug(correlationId string, message string, args ...interface{})

	// Logs a low-level debug information for troubleshooting.
	Trace(correlationId string, message string, args ...interface{})
}
