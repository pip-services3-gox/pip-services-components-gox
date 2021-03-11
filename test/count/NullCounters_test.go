package test_count

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-components-gox/count"
)

func TestNullCountersSimpleCounters(t *testing.T) {
	counters := count.NewNullCounters()
	counters.Last("Test.LastValue", 123)
	counters.Increment("Test.Increment", 3)
	counters.Stats("Test.Statistics", 123)
}

func TestNullCountersMeasureElapsedTime(t *testing.T) {
	counters := count.NewNullCounters()
	timer := counters.BeginTiming("Test.Elapsed")
	timer.EndTiming()
}
