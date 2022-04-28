package info

import (
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/build"
)

// Creates information components by their descriptors.

var ContextInfoDescriptor = refer.NewDescriptor("pip-services", "context-info", "default", "*", "1.0")
var ContainerInfoDescriptor = refer.NewDescriptor("pip-services", "container-info", "default", "*", "1.0")
var ContainerInfoDescriptor2 = refer.NewDescriptor("pip-services-container", "container-info", "default", "*", "1.0")

// NewDefaultInfoFactory create a new instance of the factory.
//	Returns: *build.Factory
func NewDefaultInfoFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(ContextInfoDescriptor, NewContextInfo)
	factory.RegisterType(ContainerInfoDescriptor, NewContextInfo)
	factory.RegisterType(ContainerInfoDescriptor2, NewContextInfo)

	return factory
}
