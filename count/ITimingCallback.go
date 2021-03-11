package count

/*
Interface for a callback to end measurement of execution elapsed time.
*/

// Ends measurement of execution elapsed time and updates specified counter.
// see
// Timing.endTiming
// Parameters:
//   - name string
//   a counter name
//   - elapsed float32
//   execution elapsed time in milliseconds to update the counter.
type ITimingCallback interface {
	EndTiming(name string, elapsed float32)
}
