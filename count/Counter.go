package count

import (
	"math"
	"time"
)

/*
Data object to store measurement for a performance counter. This object is used by CachedCounters to store counters.
*/
type Counter struct {
	Name    string    `json:"name"`
	Type    int       `json:"type"`
	Last    float32   `json:"last"`
	Count   int       `json:"count"`
	Min     float32   `json:"min"`
	Max     float32   `json:"max"`
	Average float32   `json:"average"`
	Time    time.Time `json:"time"`
}

// Creates a instance of the data obejct
// Parameters:
//   - name string
//   a counter name.
//   - type CounterType
//   a counter type.
// Returns *Counter
func NewCounter(name string, typ int) *Counter {
	return &Counter{
		Name:    name,
		Type:    typ,
		Last:    0,
		Count:   0,
		Min:     math.MaxFloat32,
		Max:     -math.MaxFloat32,
		Average: 0,
	}
}
