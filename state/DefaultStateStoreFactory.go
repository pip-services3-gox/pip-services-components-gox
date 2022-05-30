package state

import (
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/build"
)

// NewDefaultStateStoreFactory creates IStateStore components by their descriptors.
//	See Factory
//	See IStateStore
//	See MemoryStateStore
//	See NullStateStore
func NewDefaultStateStoreFactory() *build.Factory {
	factory := build.NewFactory()
	nullStateStoreDescriptor := refer.NewDescriptor("pip-services", "state-store", "null", "*", "1.0")
	memoryStateStoreDescriptor := refer.NewDescriptor("pip-services", "state-store", "memory", "*", "1.0")

	factory.RegisterType(nullStateStoreDescriptor, NewEmptyNullStateStore[any])
	factory.RegisterType(memoryStateStoreDescriptor, NewEmptyMemoryStateStore[any])

	return factory
}
