package lock

/*
Dummy lock implementation that doesn't do anything.
It can be used in testing or in situations when lock is required but shall be disabled.
*/
type NullLock struct{}

func NewNullLock() *NullLock {
	return &NullLock{}
}

// Makes a single attempt to acquire a lock by its key. It returns immediately a positive or negative result.
// Parameters:
//   - correlationId string
//    transaction id to trace execution through call chain.
//   - key string
//   a unique lock key to acquire.
//   - ttl int64
//   a lock timeout (time to live) in milliseconds.
// Returns bool, error
// true if locked. Error object
func (c *NullLock) TryAcquireLock(correlationId string,
	key string, ttl int) (bool, error) {
	return true, nil
}

// Makes multiple attempts to acquire a lock by its key within give time interval.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - key string
//   a unique lock key to acquire.
//   ttl int64
//   a lock timeout (time to live) in milliseconds.
//   timeout int64
//   a lock acquisition timeout.
// Returns error
func (c *NullLock) AcquireLock(correlationId string,
	key string, ttl int, timeout int) error {
	return nil
}

// Releases the lock with the given key.
// Parameters:
//   - correlationId string
//   not used.
//   - key string
//   the key of the lock that is to be released.
// Return error
func (c *NullLock) ReleaseLock(correlationId string,
	key string) error {
	return nil
}
