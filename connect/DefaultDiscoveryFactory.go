package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/build"
)

/*
Creates IDiscovery components by their descriptors.
*/
var MemoryDiscoveryDescriptor = refer.NewDescriptor("pip-services", "discovery", "memory", "*", "1.0")

// Create a new instance of the factory.
// Returns *build.Factory
func NewDefaultDiscoveryFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(MemoryDiscoveryDescriptor, NewEmptyMemoryDiscovery)

	return factory
}
