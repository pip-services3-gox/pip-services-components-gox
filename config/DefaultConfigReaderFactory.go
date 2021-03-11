package config

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/build"
)

/*
Creates IConfigReader components by their descriptors.
*/
var MemoryConfigReaderDescriptor = refer.NewDescriptor("pip-services", "config-reader", "memory", "*", "1.0")
var JsonConfigReaderDescriptor = refer.NewDescriptor("pip-services", "config-reader", "json", "*", "1.0")
var YamlConfigReaderDescriptor = refer.NewDescriptor("pip-services", "config-reader", "yaml", "*", "1.0")

//Create a new instance of the factory.
//Returns *build.Factory
func NewDefaultConfigReaderFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(MemoryConfigReaderDescriptor, NewEmptyMemoryConfigReader)
	factory.RegisterType(JsonConfigReaderDescriptor, NewEmptyJsonConfigReader)
	factory.RegisterType(YamlConfigReaderDescriptor, NewEmptyYamlConfigReader)

	return factory
}
