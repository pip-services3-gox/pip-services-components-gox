package test_count

import (
	"context"
	"testing"

	"github.com/pip-services3-gox/pip-services3-components-gox/count"
)

func TestNullCountersSimpleCounters(t *testing.T) {
	counters := count.NewNullCounters()
	counters.Last(context.Background(), "Test.LastValue", 123)
	counters.Increment(context.Background(), "Test.Increment", 3)
	counters.Stats(context.Background(), "Test.Statistics", 123)
}

func TestNullCountersMeasureElapsedTime(t *testing.T) {
	counters := count.NewNullCounters()
	timer := counters.BeginTiming(context.Background(), "Test.Elapsed")
	timer.EndTiming(context.Background())
}
