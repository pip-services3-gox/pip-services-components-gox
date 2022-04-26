package auth

import (
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/build"
)

// MemoryCredentialStoreDescriptor Creates ICredentialStore components by their descriptors.
var MemoryCredentialStoreDescriptor = refer.NewDescriptor(
	"pip-services",
	"credential-store",
	"memory",
	"*",
	"1.0",
)

// NewDefaultCredentialStoreFactory create a new instance of the factory.
//	Returns: *build.Factory
func NewDefaultCredentialStoreFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(MemoryCredentialStoreDescriptor, NewEmptyMemoryCredentialStore)

	return factory
}
