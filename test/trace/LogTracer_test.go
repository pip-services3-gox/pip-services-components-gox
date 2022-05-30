package test_tracer

import (
	"context"
	"errors"
	"testing"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	clog "github.com/pip-services3-gox/pip-services3-components-gox/log"
	"github.com/pip-services3-gox/pip-services3-components-gox/trace"
	ctrace "github.com/pip-services3-gox/pip-services3-components-gox/trace"
)

func newLogTracer() *ctrace.LogTracer {
	tracer := trace.NewLogTracer()
	ctx := context.Background()
	tracer.SetReferences(
		ctx,
		cref.NewReferencesFromTuples(
			ctx,
			cref.NewDescriptor("pip-services", "logger", "null", "default", "1.0"),
			clog.NewNullLogger()))
	return tracer
}

func TestSimpleTracing(t *testing.T) {
	tracer := newLogTracer()
	tracer.Trace(context.Background(), "123", "mycomponent", "mymethod", 123456)
	tracer.Failure(context.Background(), "123", "mycomponent", "mymethod", errors.New("Test error"), 123456)
}

func TestTraceTiming(t *testing.T) {
	tracer := newLogTracer()
	var timing = tracer.BeginTrace(context.Background(), "123", "mycomponent", "mymethod")
	timing.EndTrace(context.Background())

	timing = tracer.BeginTrace(context.Background(), "123", "mycomponent", "mymethod")
	timing.EndFailure(context.Background(), errors.New("Test error"))
}
