package test_count

import (
	"testing"
	"time"

	"github.com/pip-services3-gox/pip-services3-components-gox/count"
	"github.com/stretchr/testify/assert"
)

type CountersFixture struct {
	counters *count.CachedCounters
}

func NewCountersFixture(counters *count.CachedCounters) *CountersFixture {
	return &CountersFixture{
		counters: counters,
	}
}

func (c *CountersFixture) TestSimpleCounters(t *testing.T) {
	c.counters.Last("Test.LastValue", 123)
	c.counters.Last("Test.LastValue", 123456)

	var counter = c.counters.Get("Test.LastValue", count.LastValue)
	assert.NotNil(t, counter)
	assert.Equal(t, float32(123456), counter.Last)

	c.counters.IncrementOne("Test.Increment")
	c.counters.Increment("Test.Increment", 3)

	counter = c.counters.Get("Test.Increment", count.Increment)
	assert.NotNil(t, counter)
	assert.Equal(t, 4, counter.Count)

	c.counters.TimestampNow("Test.Timestamp")
	c.counters.TimestampNow("Test.Timestamp")

	counter = c.counters.Get("Test.Timestamp", count.Timestamp)
	assert.NotNil(t, counter)
	assert.NotNil(t, counter.Time)

	c.counters.Stats("Test.Statistics", 1)
	c.counters.Stats("Test.Statistics", 2)
	c.counters.Stats("Test.Statistics", 3)

	counter = c.counters.Get("Test.Statistics", count.Statistics)
	assert.NotNil(t, counter)
	assert.Equal(t, float32(2), counter.Average)

	c.counters.Dump()
}

func (c *CountersFixture) TestMeasureElapsedTime(t *testing.T) {
	timing := c.counters.BeginTiming("Test.Elapsed")

	time.Sleep(100 * time.Millisecond)

	timing.EndTiming()

	counter := c.counters.Get("Test.Elapsed", count.Interval)
	assert.NotNil(t, counter)
	assert.True(t, counter.Last > 50)
	assert.True(t, counter.Last < 5000)

	c.counters.Dump()
}
