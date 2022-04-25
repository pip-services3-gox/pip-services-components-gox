package test

import (
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/build"
)

var ShutdownDescriptor = refer.NewDescriptor("pip-services", "shutdown", "*", "*", "1.0")

func NewDefaultTestFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(ShutdownDescriptor, NewShutdown)

	return factory
}
