package lock

import (
	"context"
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

// Lock abstract lock that implements default lock acquisition routine.
//	Configuration parameters:
//		options:
//		retry_timeout: timeout in milliseconds to retry lock acquisition. (Default: 100)
type Lock struct {
	retryTimeout int64
	locker       ILock
}

const (
	DefaultRetryTimeout            int64  = 100
	ConfigParamOptionsRetryTimeout string = "options.retry_timeout"
)

// InheritLock inherit lock fron ILock
//	Returns: *Lock
func InheritLock(locker ILock) *Lock {
	return &Lock{
		retryTimeout: DefaultRetryTimeout,
		locker:       locker,
	}
}

// Configure component by passing configuration parameters.
//	Parameters: config *config.ConfigParams configuration parameters to be set.
func (c *Lock) Configure(config *config.ConfigParams) {
	c.retryTimeout = config.GetAsLongWithDefault(ConfigParamOptionsRetryTimeout, c.retryTimeout)
}

// AcquireLock makes multiple attempts to acquire a lock by its key within give time interval.
//	Parameters:
//		- ctx context.Context
//		- correlationId string transaction id to trace execution through call chain.
//		- key string a unique lock key to acquire.
//		- ttl int64 a lock timeout (time to live) in milliseconds.
//		- timeout int64 a lock acquisition timeout.
// Returns error
func (c *Lock) AcquireLock(ctx context.Context, correlationId string,
	key string, ttl int64, timeout int64) error {

	expireTime := time.Now().Add(time.Duration(timeout) * time.Millisecond)

	// Repeat until time expires
	for time.Now().Before(expireTime) {
		// Try to get lock first
		locked, err := c.locker.TryAcquireLock(ctx, correlationId, key, ttl)
		if locked || err != nil {
			return err
		}

		// Sleep
		time.Sleep(time.Duration(c.retryTimeout) * time.Millisecond)
	}

	// Throw exception
	err := errors.NewConflictError(
		correlationId,
		"LOCK_TIMEOUT",
		"Acquiring lock "+key+" failed on timeout",
	).WithDetails("key", key)

	return err
}
