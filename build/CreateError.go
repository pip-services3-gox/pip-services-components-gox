package build

// Error raised when factory is not able to create requested component.

import (
	"fmt"

	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

// NewCreateError creates an error instance and assigns its values.
//	Parameters:
//		- correlationId string
//		- message string human-readable error of the component that cannot be created.
//	Returns: *errors.ApplicationError
func NewCreateError(correlationId string, message string) *errors.ApplicationError {
	return errors.NewInternalError(correlationId, "CANNOT_CREATE", message)
}

// NewCreateErrorByLocator creates an error instance and assigns its values.
//	Parameters:
//		- correlationId string
//		- locator any human-readable locator of the component that cannot be created.
//	Returns: *errors.ApplicationError
func NewCreateErrorByLocator(correlationId string, locator any) *errors.ApplicationError {
	message := fmt.Sprintf("Requested component %v cannot be created", locator)
	return errors.NewInternalError(correlationId, "CANNOT_CREATE", message).
		WithDetails("locator", locator)
}
