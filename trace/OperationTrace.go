package trace

import (
	"time"

	cerr "github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

// OperationTrace data object to store captured operation traces.
// This object is used by CachedTracer.
type OperationTrace struct {
	// The time when operation was executed
	Time time.Time
	// The source (context name)
	Source string `json:"source"`
	// The name of component
	Component string `json:"component"`
	// The name of the executed operation
	Operation string `json:"operation"`
	// The transaction id to trace execution through call chain.
	CorrelationId string `json:"correlation_id"`
	// The duration of the operation in milliseconds
	Duration int64 `json:"duration"`

	//The description of the captured error
	//ErrorDescription
	//ApplicationException
	Error cerr.ErrorDescription `json:"error"`
}
