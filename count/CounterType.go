package count

/*
Types of counters that measure different types of metrics
Interval: = 0 Counters that measure execution time intervals
LastValue: = 1 Counters that keeps the latest measured value
Statistics: = 2 Counters that measure min/average/max statistics
Timestamp: = 3 Counter that record timestamps
Increment: = 4 Counter that increment counters
*/
const (
	Interval   = 0
	LastValue  = 1
	Statistics = 2
	Timestamp  = 3
	Increment  = 4
)

// TypeToString method converting counter type to string
func TypeToString(t int) string {
	name := ""

	switch t {
	case Interval:
		name = "interval"
	case LastValue:
		name = "lastvalue"
	case Statistics:
		name = "statistics"
	case Timestamp:
		name = "timestamp"
	case Increment:
		name = "increment"

	}

	return name
}
