package build

/*
Error raised when factory is not able to create requested component.
*/
import (
	"fmt"

	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

// Creates an error instance and assigns its values.
// Parameters:
//   - correlationId string
//   - message string
//   human-readable error of the component that cannot be created.
// Returns *errors.ApplicationError
func NewCreateError(correlationId string, message string) *errors.ApplicationError {
	e := errors.NewInternalError(correlationId, "CANNOT_CREATE", message)
	return e
}

// Creates an error instance and assigns its values.
// Parameters:
//   - correlationId string
//   - locator interface{}
//   human-readable locator of the component that cannot be created.
// Returns *errors.ApplicationError
func NewCreateErrorByLocator(correlationId string, locator interface{}) *errors.ApplicationError {
	message := fmt.Sprintf("Requested component %v cannot be created", locator)
	e := errors.NewInternalError(correlationId, "CANNOT_CREATE", message)
	e.WithDetails("locator", locator)
	return e
}
