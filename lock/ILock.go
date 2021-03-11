package lock

/*
Interface for locks to synchronize work or parallel processes and to prevent collisions.

The lock allows to manage multiple locks identified by unique keys.
*/
type ILock interface {
	// Makes a single attempt to acquire a lock by its key. It returns immediately a positive or negative result.
	TryAcquireLock(correlationId string, key string, ttl int64) (bool, error)

	// Makes multiple attempts to acquire a lock by its key within give time interval.
	AcquireLock(correlationId string, key string, ttl int64, timeout int64) error

	// Releases prevously acquired lock by its key.
	ReleaseLock(correlationId string, key string) error
}
