package test_lock

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-components-gox/lock"
	"github.com/stretchr/testify/assert"
)

const LOCK1 = "lock_1"
const LOCK2 = "lock_2"
const LOCK3 = "lock_3"

type LockFixture struct {
	locker lock.ILock
}

func NewLockFixture(locker lock.ILock) *LockFixture {
	return &LockFixture{
		locker: locker,
	}
}

func (c *LockFixture) TestTryAcquireLock(t *testing.T) {
	// Try to acquire lock for the first time
	result, err := c.locker.TryAcquireLock("", LOCK1, 3000)
	assert.Nil(t, err)
	assert.True(t, result)

	// Try to acquire lock for the second time
	result, err = c.locker.TryAcquireLock("", LOCK1, 3000)
	assert.Nil(t, err)
	assert.False(t, result)

	// Release the lock
	err = c.locker.ReleaseLock("", LOCK1)
	assert.Nil(t, err)

	// Try to acquire lock for the third time
	result, err = c.locker.TryAcquireLock("", LOCK1, 3000)
	assert.Nil(t, err)
	assert.True(t, result)

	err = c.locker.ReleaseLock("", LOCK1)
	assert.Nil(t, err)
}

func (c *LockFixture) TestAcquireLock(t *testing.T) {
	// Acquire lock for the first time
	c.locker.AcquireLock("", LOCK2, 3000, 1000)

	// Acquire lock for the second time
	err := c.locker.AcquireLock("", LOCK2, 3000, 1000)
	assert.NotNil(t, err)

	// Release the lock
	err = c.locker.ReleaseLock("", LOCK2)
	assert.Nil(t, err)

	// Acquire lock for the third time
	err = c.locker.AcquireLock("", LOCK2, 3000, 1000)
	assert.Nil(t, err)

	err = c.locker.ReleaseLock("", LOCK2)
	assert.Nil(t, err)
}

func (c *LockFixture) TestReleaseLock(t *testing.T) {
	// Acquire lock for the first time
	result, err := c.locker.TryAcquireLock("", LOCK3, 3000)
	assert.Nil(t, err)
	assert.True(t, result)

	// Release the lock for the first time
	err = c.locker.ReleaseLock("", LOCK3)
	assert.Nil(t, err)

	// Release the lock for the second time
	err = c.locker.ReleaseLock("", LOCK3)
	assert.Nil(t, err)
}
