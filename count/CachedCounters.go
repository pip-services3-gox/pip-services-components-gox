package count

import (
	"math"
	"sync"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

/*
Abstract implementation of performance counters that measures and stores counters in memory. Child classes implement saving of the counters into various destinations.

Configuration parameters
  options:
    interval: interval in milliseconds to save current counters measurements (default: 5 mins)
    reset_timeout: timeout in milliseconds to reset the counters. 0 disables the reset (default: 0)
*/
type CachedCounters struct {
	cache         map[string]*Counter
	updated       bool
	lastDumpTime  time.Time
	lastResetTime time.Time
	mux           sync.Mutex
	interval      int64
	resetTimeout  int64
	saver         ICountersSaver
}

type ICountersSaver interface {
	Save(counters []*Counter) error
}

// Inherit cache counters from saver
// Parameters:
//  - save ICountersSaver
// Returns *CachedCounters
func InheritCacheCounters(saver ICountersSaver) *CachedCounters {
	return &CachedCounters{
		cache:         map[string]*Counter{},
		updated:       false,
		lastDumpTime:  time.Now(),
		lastResetTime: time.Now(),
		interval:      300000,
		resetTimeout:  0,
		saver:         saver,
	}
}

// Configures component by passing configuration parameters.
// Parameters:
//   - config *config.ConfigParams
//   configuration parameters to be set.
func (c *CachedCounters) Configure(config *config.ConfigParams) {
	c.interval = config.GetAsLongWithDefault("interval", c.interval)
	c.resetTimeout = config.GetAsLongWithDefault("reset_timeout", c.resetTimeout)
}

// Clears (resets) a counter specified by its name.
// Parameters:
//   - name string
//   a counter name to clear.
func (c *CachedCounters) Clear(name string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	delete(c.cache, name)
}

//Clears (resets) all counters.
func (c *CachedCounters) ClearAll() {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache = map[string]*Counter{}
}

// Dumps (saves) the current values of counters.
func (c *CachedCounters) Dump() error {
	if !c.updated {
		return nil
	}

	counters := c.GetAll()
	err := c.saver.Save(counters)
	if err != nil {
		return err
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	c.updated = false
	c.lastDumpTime = time.Now()

	return nil
}

func (c *CachedCounters) update() error {
	c.updated = true
	newDumpTime := c.lastDumpTime.Add(time.Duration(c.interval) * time.Millisecond)
	if time.Now().After(newDumpTime) {
		return c.Dump()
	}
	return nil
}

func (c *CachedCounters) resetIfNeeded() {
	if c.resetTimeout == 0 {
		return
	}

	newResetTime := c.lastResetTime.Add(time.Duration(c.resetTimeout) * time.Millisecond)
	if time.Now().After(newResetTime) {
		c.cache = map[string]*Counter{}
		c.updated = false
		c.lastDumpTime = time.Now()
	}
}

//Gets all captured counters.
//Returns []*Counter
func (c *CachedCounters) GetAll() []*Counter {
	c.mux.Lock()
	defer c.mux.Unlock()

	result := []*Counter{}
	for _, v := range c.cache {
		result = append(result, v)
	}

	return result
}

// Gets a counter specified by its name. It counter does not exist or its type doesn't match the specified type it creates a new one.
// Parameters:
//   - name string
//   a counter name to retrieve.
//   - typ int
//   a counter type.
// Returns *Counter
// an existing or newly created counter of the specified type.
func (c *CachedCounters) Get(name string, typ int) *Counter {
	if name == "" {
		panic("Counter name cannot be nil")
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	c.resetIfNeeded()

	counter, ok := c.cache[name]
	if !ok || counter.Type != typ {
		counter = NewCounter(name, typ)
		c.cache[name] = counter
	}

	return counter
}

func (c *CachedCounters) calculateStats(counter *Counter, value float32) {
	if counter == nil {
		panic("Counter cannot be nil")
	}

	counter.Last = value
	counter.Count++
	counter.Max = float32(math.Max(float64(counter.Max), float64(value)))
	counter.Min = float32(math.Min(float64(counter.Min), float64(value)))
	counter.Average = ((counter.Average * float32(counter.Count-1)) + value) / float32(counter.Count)
}

// Begins measurement of execution time interval. It returns Timing object which has to be called at Timing.endTiming to end the measurement and update the counter.
// Parameters
//   - name string
//   a counter name of Interval type.
// Returns *Timing
// a Timing callback object to end timing.
func (c *CachedCounters) BeginTiming(name string) *Timing {
	return NewTiming(name, c)
}

// Ends measurement of execution elapsed time and updates specified counter.
// see
// Timing.endTiming
// Parameters:
//   - name string
//   a counter name
//   elapsed float32
//   execution elapsed time in milliseconds to update the counter.
func (c *CachedCounters) EndTiming(name string, elapsed float32) {
	counter := c.Get(name, Interval)
	c.calculateStats(counter, elapsed)
	c.update()
}

// Calculates min/average/max statistics based on the current and previous values.
// Parameters:
//   - name string
//   a counter name of Statistics type
//   - value float32
//   a value to update statistics
func (c *CachedCounters) Stats(name string, value float32) {
	counter := c.Get(name, Statistics)
	c.calculateStats(counter, value)
	c.update()
}

// Records the last calculated measurement value.
// Usually this method is used by metrics calculated externally.
// Parameters:
//   - name string
//   a counter name of Last type.
//   - value number
//   a last value to record.
func (c *CachedCounters) Last(name string, value float32) {
	counter := c.Get(name, LastValue)
	counter.Last = value
	c.update()
}

// Records the current time as a timestamp.
// Parameters:
//   - name string
//   a counter name of Timestamp type.
func (c *CachedCounters) TimestampNow(name string) {
	c.Timestamp(name, time.Now())
}

// Records the given timestamp.
// Parameters:
//   - name string
//   a counter name of Timestamp type.
//   value time.Time
//   a timestamp to record.
func (c *CachedCounters) Timestamp(name string, value time.Time) {
	counter := c.Get(name, Timestamp)
	counter.Time = value
	c.update()
}

// Increments counter by 1.
// Parameters:
//   - name string
//   a counter name of Increment type.
func (c *CachedCounters) IncrementOne(name string) {
	c.Increment(name, 1)
}

// Increments counter by given value.
// Parameters:
//   - name string
//   a counter name of Increment type.
//   - value int
//   a value to add to the counter.
func (c *CachedCounters) Increment(name string, value int) {
	counter := c.Get(name, Increment)
	counter.Count = counter.Count + value
	c.update()
}
