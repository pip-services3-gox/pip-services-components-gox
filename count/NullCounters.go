package count

import "time"

/*
Dummy implementation of performance counters that doesn't do anything.

It can be used in testing or in situations when counters is required but shall be disabled.
*/
type NullCounters struct{}

// Creates a new instance of the counter.
// Returns *NullCounters
func NewNullCounters() *NullCounters {
	return &NullCounters{}
}

// Begins measurement of execution time interval. It returns Timing object which has to be called at Timing.endTiming to end the measurement and update the counter.
// Parameters:
//   - name string
//   a counter name of Interval type.
// Returns *Timing
// a Timing callback object to end timing.
func (c *NullCounters) BeginTiming(name string) *Timing {
	return NewEmptyTiming()
}

// Calculates min/average/max statistics based on the current and previous values.
// Parameters:
//   - name string
//   a counter name of Statistics type
//   - value float32
//   a value to update statistics
func (c *NullCounters) Stats(name string, value float32) {}

// Records the last calculated measurement value.
// Usually this method is used by metrics calculated externally.
// Parameters:
//   - name string
//   a counter name of Last type.
//   - value float32
//   a last value to record.
func (c *NullCounters) Last(name string, value float32) {}

// Records the current time as a timestamp.
// Parameters:
//   - name string
//   a counter name of Timestamp type.
func (c *NullCounters) TimestampNow(name string) {}

// Records the given timestamp.
// Parameters:
//   - name string
//   a counter name of Timestamp type.
//   - value time.Time
//   a timestamp to record.
func (c *NullCounters) Timestamp(name string, value time.Time) {}

// Increments counter by 1.
// Parameters:
//   - name string
//   a counter name of Increment type.
func (c *NullCounters) IncrementOne(name string) {}

// Increments counter by given value.
// Parameters:
//   - name string
//   a counter name of Increment type.
//   - value float32
//   a value to add to the counter.
func (c *NullCounters) Increment(name string, value float32) {}
