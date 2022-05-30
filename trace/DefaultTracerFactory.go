package trace

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

// DefaultTracerFactory creates [[ITracer]] components by their descriptors.
//	See [[Factory]]
//	See [[NullTracer]]
//	See [[ConsoleTracer]]
//	See [[CompositeTracer]]
type DefaultTracerFactory struct {
	cbuild.Factory
	NullTracerDescriptor      *cref.Descriptor
	LogTracerDescriptor       *cref.Descriptor
	CompositeTracerDescriptor *cref.Descriptor
}

// NewDefaultTracerFactory create a new instance of the factory.
func NewDefaultTracerFactory() *DefaultTracerFactory {
	c := &DefaultTracerFactory{
		Factory:                   *cbuild.NewFactory(),
		NullTracerDescriptor:      cref.NewDescriptor("pip-services", "tracer", "null", "*", "1.0"),
		LogTracerDescriptor:       cref.NewDescriptor("pip-services", "tracer", "log", "*", "1.0"),
		CompositeTracerDescriptor: cref.NewDescriptor("pip-services", "tracer", "composite", "*", "1.0"),
	}

	c.RegisterType(c.NullTracerDescriptor, NewNullTracer)
	c.RegisterType(c.LogTracerDescriptor, NewLogTracer)
	c.RegisterType(c.CompositeTracerDescriptor, NewCompositeTracer)

	return c
}
