package count

import "time"

/*
Interface for performance counters that measure execution metrics.

The performance counters measure how code is performing: how fast or slow, how many transactions performed, how many objects are stored, what was the latest transaction time and so on.

They are critical to monitor and improve performance, scalability and reliability of code in production.
*/

type ICounters interface {
	// Begins measurement of execution time interval. It returns Timing object which has to be called at Timing.endTiming to end the measurement and update the counter.
	BeginTiming(name string) *Timing

	// Calculates min/average/max statistics based on the current and previous values.
	Stats(name string, value float32)

	// Records the last calculated measurement value.
	// Usually this method is used by metrics calculated externally.
	Last(name string, value float32)

	// Records the given timestamp.
	TimestampNow(name string)

	// Records the current time as a timestamp.
	Timestamp(name string, value time.Time)

	// Increments counter by 1.
	IncrementOne(name string)

	// Increments counter by given value.
	Increment(name string, value int)
}
