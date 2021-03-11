package count

import (
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/refer"
)

/*
Aggregates all counters from component references under a single component.

It allows to capture metrics and conveniently send them to multiple destinations.

References
*:counters:*:*:1.0 (optional) ICounters components to pass collected measurements
see
ICounters

Example
  type MyComponent {
      _counters CompositeCounters = new CompositeCounters();
  }
      func (mc * MyConponent)setReferences(references: IReferences) {
          mc._counters.SetReferences(references);
      }

      func (mc * MyConponent) myMethod() {
          mc._counters.Increment("mycomponent.mymethod.calls");
          timing := mc._counters.BeginTiming("mycomponent.mymethod.exec_time");
  		 defer timing.EndTiming();
  		// do something

      	}
  var mc MyComponent{};
  mc._counters = NewCompositeCounters();

*/
type CompositeCounters struct {
	counters []ICounters
}

// Creates a new instance of the counters.
// Returns *CompositeCounters
func NewCompositeCounters() *CompositeCounters {
	c := &CompositeCounters{
		counters: []ICounters{},
	}
	return c
}

// Creates a new instance of the counters.
// Parameters:
// 			- references refer.IReferences
// 			references to locate the component dependencies.
// Returns *CompositeCounters
func NewCompositeCountersFromReferences(references refer.IReferences) *CompositeCounters {
	c := NewCompositeCounters()
	c.SetReferences(references)
	return c
}

// Sets references to dependent components.
// Parameters:
// 			- references refer.IReferences
// 			references to locate the component dependencies.
func (c *CompositeCounters) SetReferences(references refer.IReferences) {
	if c.counters == nil {
		c.counters = []ICounters{}
	}

	counters := references.GetOptional(refer.NewDescriptor("*", "counters", "*", "*", "*"))
	for _, l := range counters {
		if l == c {
			continue
		}

		counter, ok := l.(ICounters)
		if ok {
			c.counters = append(c.counters, counter)
		}
	}
}

// Begins measurement of execution time interval. It returns Timing object which has to be called at Timing.endTiming to end the measurement and update the counter.
// Parameters:
// 			- name string
// 			a counter name of Interval type.
// Returns *Timing
// a Timing callback object to end timing.
func (c *CompositeCounters) BeginTiming(name string) *Timing {
	return NewTiming(name, c)
}

// Ends measurement of execution elapsed time and updates specified counter.
// see
// Timing.endTiming
// Parameters:
// 			- name string
// 			a counter name
// 			- elapsed float32
// 			execution elapsed time in milliseconds to update the counter.
func (c *CompositeCounters) EndTiming(name string, elapsed float32) {
	for _, counter := range c.counters {
		callback, ok := counter.(ITimingCallback)
		if ok {
			callback.EndTiming(name, elapsed)
		}
	}
}

// Calculates min/average/max statistics based on the current and previous values.
// Parameters:
// 			- name string
// 			a counter name of Statistics type
// 			- value float32
// 			a value to update statistics
func (c *CompositeCounters) Stats(name string, value float32) {
	for _, counter := range c.counters {
		counter.Stats(name, value)
	}
}

// Records the last calculated measurement value.
// Usually this method is used by metrics calculated externally.
// Parameters:
// 			- name string
// 			a counter name of Last type.
// 			- value float32
// 			a last value to record.
func (c *CompositeCounters) Last(name string, value float32) {
	for _, counter := range c.counters {
		counter.Last(name, value)
	}
}

// Records the current time as a timestamp.
// Parameters:
// 			- name string
// 			a counter name of Timestamp type.
func (c *CompositeCounters) TimestampNow(name string) {
	c.Timestamp(name, time.Now())
}

// Records the given timestamp.
// Parameters:
// 			- name string
// 			a counter name of Timestamp type.
// 			- value time.Time
// 			a timestamp to record.
func (c *CompositeCounters) Timestamp(name string, value time.Time) {
	for _, counter := range c.counters {
		counter.Timestamp(name, value)
	}
}

// Increments counter by 1.
// Parameters:
// 			- name string
// 			a counter name of Increment type.
func (c *CompositeCounters) IncrementOne(name string) {
	c.Increment(name, 1)
}

// Increments counter by given value.
// Parameters:
// 			- name string
// 			a counter name of Increment type.
// 			- value number
// 			a value to add to the counter.
func (c *CompositeCounters) Increment(name string, value int) {
	for _, counter := range c.counters {
		counter.Increment(name, value)
	}
}
