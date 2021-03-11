package log

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
)

/*
Helper class to convert log level values.
*/
type TLogLevelConverter struct{}

var LogLevelConverter *TLogLevelConverter = &TLogLevelConverter{}

// Converts numbers and strings to standard log level values.
// Parameters:
//   - value interface{}
//   a value to be converted
// Returns int
// converted log level
func (c *TLogLevelConverter) ToLogLevel(value interface{}) int {
	return LogLevelFromString(value)
}

// Converts log level to a string.
// see
// LogLevel
// Parameters:
//   - level int
//   a log level to convert
// Returns string
// log level name string.
func (c *TLogLevelConverter) ToString(level int) string {
	return LogLevelToString(level)
}

// Converts log level to a LogLevel.
// see
// LogLevel
// Parameters:
//   - value interface{}
//   a log level string to convert
// Returns int
// log level value.
func LogLevelFromString(value interface{}) int {
	if value == nil {
		return Info
	}

	str := convert.StringConverter.ToString(value)
	str = strings.ToUpper(str)
	if "0" == str || "NOTHING" == str || "NONE" == str {
		return None
	} else if "1" == str || "FATAL" == str {
		return Fatal
	} else if "2" == str || "ERROR" == str {
		return Error
	} else if "3" == str || "WARN" == str || "WARNING" == str {
		return Warn
	} else if "4" == str || "INFO" == str {
		return Info
	} else if "5" == str || "DEBUG" == str {
		return Debug
	} else if "6" == str || "TRACE" == str {
		return Trace
	} else {
		return Info
	}
}

// Converts log level to a string.
// see
// LogLevel
// Parameters:
//   - level int
//   a log level to convert
// Returns string
// log level name string.
func LogLevelToString(level int) string {
	switch level {
	case Fatal:
		return "FATAL"
	case Error:
		return "ERROR"
	case Warn:
		return "WARN"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	case Trace:
		return "TRACE"
	default:
		return "UNDEF"
	}
}
