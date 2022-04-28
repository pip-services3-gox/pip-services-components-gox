package log

import (
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/build"
)

// Creates ILogger components by their descriptors.

var NullLoggerDescriptor = refer.NewDescriptor("pip-services", "logger", "null", "*", "1.0")
var ConsoleLoggerDescriptor = refer.NewDescriptor("pip-services", "logger", "console", "*", "1.0")
var CompositeLoggerDescriptor = refer.NewDescriptor("pip-services", "logger", "composite", "*", "1.0")

// NewDefaultLoggerFactory create a new instance of the factory.
//	Returns: *build.Factory
func NewDefaultLoggerFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(NullLoggerDescriptor, NewNullLogger)
	factory.RegisterType(ConsoleLoggerDescriptor, NewConsoleLogger)
	factory.RegisterType(CompositeLoggerDescriptor, NewCompositeLogger)

	return factory
}
