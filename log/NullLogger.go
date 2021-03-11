package log

/*

Dummy implementation of logger that doesn't do anything.

It can be used in testing or in situations when logger is required but shall be disabled.
*/
type NullLogger struct{}

// Creates a new instance of the logger.
// Returns *NullLogger
func NewNullLogger() *NullLogger {
	c := &NullLogger{}
	return c
}

// Gets the maximum log level. Messages with higher log level are filtered out.
// Returns int
// the maximum log level.
func (c *NullLogger) Level() int {
	return None
}

// Set the maximum log level.
// Parameters:
//   - value int
// a new maximum log level.
func (c *NullLogger) SetLevel(value int) {
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
func (c *NullLogger) Log(level int, correlationId string, err error, message string, args ...interface{}) {
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
func (c *NullLogger) Fatal(correlationId string, err error, message string, args ...interface{}) {
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
func (c *NullLogger) Error(correlationId string, err error, message string, args ...interface{}) {
}

// Logs a warning that may or may not have a negative impact.
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - message string
//  a human-readable message to log.
//  - args ...interface{}
//  arguments to parameterize the message
func (c *NullLogger) Warn(correlationId string, message string, args ...interface{}) {
}

// Logs an important information message
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - message string
//   a human-readable message to log.
//   - args ...interface{}
//   arguments to parameterize the message
func (c *NullLogger) Info(correlationId string, message string, args ...interface{}) {
}

// Logs a high-level debug information for troubleshooting.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - message string
//   a human-readable message to log.
//   - args ...interface{}
//   arguments to parameterize the message
func (c *NullLogger) Debug(correlationId string, message string, args ...interface{}) {
}

// Logs a low-level debug information for troubleshooting.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - message string
//   a human-readable message to log.
//   - args ...interface{}
//   arguments to parameterize the message
func (c *NullLogger) Trace(correlationId string, message string, args ...interface{}) {
}
