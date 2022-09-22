package components

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/count"
	"github.com/pip-services3-gox/pip-services3-components-gox/log"
)

// Abstract component that supportes configurable dependencies, logging
// and performance counters.
//
//	Configuration parameters
//		- dependencies:
//			- [dependency name 1]: Dependency 1 locator (descriptor)
//			- ...
//			- [dependency name N]: Dependency N locator (descriptor)
//	References
//		- *:counters:*:*:1.0     (optional) ICounters components to pass collected measurements
//		- *:logger:*:*:1.0       (optional) ILogger components to pass log messages
//		- *:tracer:*:*:1.0       (optional) ITracer components to trace executed operations
//		- ...                                    References must match configured dependencies.
type Component struct {
	dependencyResolver *refer.DependencyResolver
	logger             *log.CompositeLogger
	counters           *count.CompositeCounters
}

func InheritComponent() *Component {
	return &Component{
		dependencyResolver: refer.NewDependencyResolver(),
		logger:             log.NewCompositeLogger(),
		counters:           count.NewCompositeCounters(),
	}
}

// Configures component by passing configuration parameters.
//	Parameters:
//		- ctx context.Context
//		- config    configuration parameters to be set.
func (c *Component) Configure(ctx context.Context, config *config.ConfigParams) {
	c.dependencyResolver.Configure(ctx, config)
	c.logger.Configure(ctx, config)
}

// SetReferences sets the component references. References must match configured dependencies.
//	Parameters:
//		- ctx context.Context
//		- references IReferences references to set.
func (c *Component) SetReferences(ctx context.Context, references refer.IReferences) {
	c.dependencyResolver.SetReferences(ctx, references)
	c.logger.SetReferences(ctx, references)
	c.counters.SetReferences(ctx, references)
}
