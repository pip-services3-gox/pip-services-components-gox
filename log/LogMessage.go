package log

import (
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

/*
Data object to store captured log messages. This object is used by CachedLogger.
*/
type LogMessage struct {
	Time          time.Time               `json:"time"`
	Source        string                  `json:"source"`
	Level         int                     `json:"level"`
	CorrelationId string                  `json:"correlation_id"`
	Error         errors.ErrorDescription `json:"error"`
	Message       string                  `json:"message"`
}

// Create new log message object
// Parameters:
//   - level int
//   an log level
//   - source string
//   an source
//   - correletionId string
//   transaction id to trace execution through call chain.
//   - err errors.ErrorDescription
//   an error object associated with this message.
//   - message string
//   a human-readable message to log.
// Returns LogMessage
func NewLogMessage(level int, source string, correlationId string,
	err errors.ErrorDescription, message string) LogMessage {
	return LogMessage{
		Time:          time.Now().UTC(),
		Source:        source,
		Level:         level,
		CorrelationId: correlationId,
		Error:         err,
		Message:       message,
	}
}
