package count

import (
	"context"
	"time"
)

// Timing callback object returned by ICounters.beginTiming to end timing of
// execution block and update the associated counter.
//	Example:
//		timing := counters.BeginTiming("mymethod.exec_time");
//		defer  timing.EndTiming();
type Timing struct {
	start    time.Time
	callback ITimingCallback
	counter  string
}

// NewEmptyTiming creates a new instance of the timing callback object.
//	Returns: *Timing
func NewEmptyTiming() *Timing {
	return &Timing{
		start: time.Now(),
	}
}

// NewTiming creates a new instance of the timing callback object.
//	Parameters:
//		- counter string an associated counter name
//		- callback ITimingCallback a callback that shall be called when EndTiming is called.
//	Returns: *Timing
func NewTiming(counter string, callback ITimingCallback) *Timing {
	return &Timing{
		start:    time.Now(),
		callback: callback,
		counter:  counter,
	}
}

// EndTiming ends timing of an execution block, calculates
// elapsed time and updates the associated counter.
func (c *Timing) EndTiming(ctx context.Context) {
	if c.callback == nil {
		return
	}

	elapsed := time.Since(c.start).Seconds() * 1000
	c.callback.EndTiming(ctx, c.counter, elapsed)
}
