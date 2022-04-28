package log

import (
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

// LogMessage Data object to store captured log messages.
// This object is used by CachedLogger.
type LogMessage struct {
	Time          time.Time               `json:"time"`
	Source        string                  `json:"source"`
	Level         LevelType               `json:"level"`
	CorrelationId string                  `json:"correlation_id"`
	Error         errors.ErrorDescription `json:"error"`
	Message       string                  `json:"message"`
}

// NewLogMessage create new log message object
//	Parameters:
//		- level LevelType a log level
//		- source string a source
//		- correlationId string transaction id to trace execution through call chain.
//		- err errors.ErrorDescription an error object associated with this message.
//		- message string a human-readable message to log.
//	Returns: LogMessage
func NewLogMessage(level LevelType, source string, correlationId string,
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
