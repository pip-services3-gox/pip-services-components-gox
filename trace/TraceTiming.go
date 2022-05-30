package trace

import (
	"context"
	"time"
)

// TraceTiming timing object returned by {ITracer.BeginTrace} to end timing
// of execution block and record the associated trace.
//
//	Example:
//		timing := tracer.BeginTrace("mymethod.exec_time");
//		...
//		timing.EndTrace();
//		if err != nil {
//			timing.EndFailure(err);
//		}
//
type TraceTiming struct {
	start         int64
	tracer        ITracer
	correlationId string
	component     string
	operation     string
}

// NewTraceTiming creates a new instance of the timing callback object.
//	Parameters:
//		- correlationId (optional) transaction id to trace execution through call chain.
//		- component 	an associated component name
//		- operation 	an associated operation name
//		- callback 		a callback that shall be called when endTiming is called.
func NewTraceTiming(correlationId string, component string, operation string, tracer ITracer) *TraceTiming {
	return &TraceTiming{
		correlationId: correlationId,
		component:     component,
		operation:     operation,
		tracer:        tracer,
		start:         time.Now().UTC().UnixNano(),
	}
}

// EndTrace ends timing of an execution block, calculates elapsed time
// and records the associated trace.
//	Parameters:
//		- ctx context.Context
func (c *TraceTiming) EndTrace(ctx context.Context) {
	if c.tracer != nil {
		elapsed := time.Now().UTC().UnixNano() - c.start
		c.tracer.Trace(ctx, c.correlationId, c.component, c.operation, elapsed/int64(time.Millisecond))
	}
}

// EndFailure ends timing of a failed block, calculates elapsed time
// and records the associated trace.
//	Parameters:
//		- ctx context.Context
func (c *TraceTiming) EndFailure(ctx context.Context, err error) {
	if c.tracer != nil {
		elapsed := time.Now().UTC().UnixNano() - c.start
		c.tracer.Failure(ctx, c.correlationId, c.component, c.operation, err, elapsed/int64(time.Millisecond))
	}
}
