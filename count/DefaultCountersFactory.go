package count

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/build"
)

/*
Creates ICounters components by their descriptors.
*/
var NullCountersDescriptor = refer.NewDescriptor("pip-services", "counters", "null", "*", "1.0")
var LogCountersDescriptor = refer.NewDescriptor("pip-services", "counters", "log", "*", "1.0")
var CompositeCountersDescriptor = refer.NewDescriptor("pip-services", "counters", "composite", "*", "1.0")

// Create a new instance of the factory.
// Returns *build.Factory
func NewDefaultCountersFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(NullCountersDescriptor, NewNullCounters)
	factory.RegisterType(LogCountersDescriptor, NewLogCounters)
	factory.RegisterType(CompositeCountersDescriptor, NewCompositeCounters)

	return factory
}
