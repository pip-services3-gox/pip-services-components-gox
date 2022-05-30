package test_tracer

import (
	"context"
	"errors"
	"testing"

	ctrace "github.com/pip-services3-gox/pip-services3-components-gox/trace"
)

func newNullTracer() *ctrace.NullTracer {
	return ctrace.NewNullTracer()
}

func TestSimpleNullTracing(t *testing.T) {
	tracer := newNullTracer()
	tracer.Trace(context.Background(), "123", "mycomponent", "mymethod", 123456)
	tracer.Failure(context.Background(), "123", "mycomponent", "mymethod", errors.New("Test error"), 123456)
}

func TestTraceNullTiming(t *testing.T) {
	tracer := newNullTracer()
	timing := tracer.BeginTrace(context.Background(), "123", "mycomponent", "mymethod")
	timing.EndTrace(context.Background())

	timing = tracer.BeginTrace(context.Background(), "123", "mycomponent", "mymethod")
	timing.EndFailure(context.Background(), errors.New("Test error"))
}
