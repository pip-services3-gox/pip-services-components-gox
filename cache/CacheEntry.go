package cache

import (
	"time"
)

/*
Data object to store cached values with their keys used by MemoryCache
*/
type CacheEntry struct {
	key        string
	value      interface{}
	expiration time.Time
}

// Creates a new instance of the cache entry and assigns its values.
// Parameters:
//   - key string
//   a unique key to locate the value.
//   - value interface{]
//   a value to be stored.
//   - timeout int64
//   expiration timeout in milliseconds.
// Returns *CacheEntry
func NewCacheEntry(key string, value interface{}, timeout int64) *CacheEntry {
	return &CacheEntry{
		key:        key,
		value:      value,
		expiration: time.Now().Add(time.Duration(timeout) * time.Millisecond),
	}
}

// Gets the key to locate the cached value.
// Returns string
// the value key.
func (c *CacheEntry) Key() string {
	return c.key
}

// Gets the cached value.
// Returns interface{}
// the value object.
func (c *CacheEntry) Value() interface{} {
	return c.value
}

// Gets the expiration timeout.
// Returns time.Time
// the expiration timeout in milliseconds.
func (c *CacheEntry) Expiration() time.Time {
	return c.expiration
}

// Sets a new value and extends its expiration.
// Parameters:
//   - value interface{}
//   a new cached value.
//   - timeout int64
//   a expiration timeout in milliseconds.
func (c *CacheEntry) SetValue(value interface{}, timeout int64) {
	c.value = value
	c.expiration = time.Now().Add(time.Duration(timeout) * time.Millisecond)
}

// Checks if this value already expired.
// Returns bool
// true if the value already expires and false otherwise.
func (c *CacheEntry) IsExpired() bool {
	return time.Now().After(c.expiration)
}
