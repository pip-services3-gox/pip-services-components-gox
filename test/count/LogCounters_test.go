package test_count

import (
	"context"
	"testing"
	"time"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/count"
	"github.com/pip-services3-gox/pip-services3-components-gox/log"
)

func TestLogCountersSimpleCounters(t *testing.T) {
	counters := count.NewLogCounters()
	fixture := NewCountersFixture(&counters.CachedCounters)
	fixture.TestSimpleCounters(t)
}

func TestLogCountersMeasureElapsedTime(t *testing.T) {
	counters := count.NewLogCounters()
	fixture := NewCountersFixture(&counters.CachedCounters)
	fixture.TestMeasureElapsedTime(t)
}

func TestLogCountersSave(t *testing.T) {
	counters := count.NewLogCounters()
	logger := log.NewConsoleLogger()

	refernces := cref.NewReferences(context.Background(), []any{cref.NewDescriptor("pip-services", "logger", "console", "*", "1.0"), logger})

	config := cconf.NewConfigParamsFromTuples("interval", 500)
	counters.Configure(context.Background(), config)
	counters.SetReferences(context.Background(), refernces)

	counters.Last(context.Background(), "Test.LastValue", 123)
	counters.Last(context.Background(), "Test.LastValue", 123456)

	<-time.After(time.Second * 2)

	counters.Last(context.Background(), "Test.LastValue", 1234689)

}
