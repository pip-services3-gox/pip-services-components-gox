package count

import (
	"time"
)

/*
Callback object returned by ICounters.beginTiming to end timing of execution block and update the associated counter.

Example
  timing := counters.BeginTiming("mymethod.exec_time");
  defer  timing.EndTiming();

*/
type Timing struct {
	start    time.Time
	callback ITimingCallback
	counter  string
}

// Creates a new instance of the timing callback object.
// Retruns *Timing
func NewEmptyTiming() *Timing {
	return &Timing{
		start: time.Now(),
	}
}

// Creates a new instance of the timing callback object.
// Parameters:
//   - counter string
//   an associated counter name
//   - callback ITimingCallback
//   a callback that shall be called when EndTiming is called.
// Retruns *Timing
func NewTiming(counter string, callback ITimingCallback) *Timing {
	return &Timing{
		start:    time.Now(),
		callback: callback,
		counter:  counter,
	}
}

// Ends timing of an execution block, calculates elapsed time and updates the associated counter.
func (c *Timing) EndTiming() {
	if c.callback == nil {
		return
	}

	elapsed := time.Since(c.start).Seconds() * 1000
	c.callback.EndTiming(c.counter, float32(elapsed))
}
