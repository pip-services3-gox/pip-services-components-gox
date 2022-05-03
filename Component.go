package components

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/count"
	"github.com/pip-services3-gox/pip-services3-components-gox/log"
)

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

func (c *Component) Configure(ctx context.Context, config *config.ConfigParams) {
	c.dependencyResolver.Configure(ctx, config)
	c.logger.Configure(ctx, config)
}

func (c *Component) SetReferences(ctx context.Context, references refer.IReferences) {
	c.dependencyResolver.SetReferences(ctx, references)
	c.logger.SetReferences(ctx, references)
	c.counters.SetReferences(ctx, references)
}
