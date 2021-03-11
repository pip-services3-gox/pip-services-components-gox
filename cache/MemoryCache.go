package cache

import (
	"encoding/json"
	"sync"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

/*
Cache that stores values in the process memory.

Remember: This implementation is not suitable for synchronization of distributed processes.

Configuration parameters
  options:
    timeout: default caching timeout in milliseconds (default: 1 minute)
    max_size: maximum number of values stored in this cache (default: 1000)
see
ICache

Example
  cache := NewMemoryCache();
  res, err := cache.Store("123", "key1", "ABC", 10000);
*/
type MemoryCache struct {
	cache   map[string]*CacheEntry
	lock    *sync.Mutex
	timeout int64
	maxSize int
}

// Creates a new instance of the cache.
// Returns *MemoryCache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache:   map[string]*CacheEntry{},
		lock:    &sync.Mutex{},
		timeout: 60000,
		maxSize: 1000,
	}
}

// Creates a new instance of the cache.
// Parameters
//   - cfg *config.ConfigParams
//   configuration parameters to be set.
// Returns *MemoryCache
func NewMemoryCacheFromConfig(cfg *config.ConfigParams) *MemoryCache {
	c := NewMemoryCache()
	c.Configure(cfg)
	return c
}

// Configures component by passing configuration parameters.
// Parameters:
//   - config *config.ConfigParams
//   configuration parameters to be set.
func (c *MemoryCache) Configure(cfg *config.ConfigParams) {
	c.timeout = cfg.GetAsLongWithDefault("timeout", c.timeout)
	c.maxSize = cfg.GetAsIntegerWithDefault("max_size", c.maxSize)
}

// Cleanup memory cache
func (c *MemoryCache) Cleanup() {
	var oldest *CacheEntry
	var keysToRemove = []string{}

	c.lock.Lock()
	defer c.lock.Unlock()

	for key, value := range c.cache {
		if value.IsExpired() {
			keysToRemove = append(keysToRemove, key)
		}
		if oldest == nil || oldest.Expiration().After(value.Expiration()) {
			oldest = value
		}
	}

	for _, key := range keysToRemove {
		delete(c.cache, key)
	}

	if len(c.cache) > c.maxSize && oldest != nil {
		delete(c.cache, oldest.Key())
	}
}

// Retrieves cached value from the cache using its key. If value is missing in the cache or expired it returns null.
// Parameters:
//   - correlationId string
//    transaction id to trace execution through call chain.
//   - key string
//   a unique value key.
// Returns interface{}, error
func (c *MemoryCache) Retrieve(correlationId string, key string) (interface{}, error) {
	if key == "" {
		panic("Key cannot be empty")
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	entry := c.cache[key]
	if entry != nil {
		if entry.IsExpired() {
			delete(c.cache, key)
			return nil, nil
		}
		var value interface{}
		err := json.Unmarshal((entry.Value()).([]byte), &value)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	return nil, nil
}

// Retrive cached value from the cache using its key and restore into reference object. If value is missing in the cache or expired it returns false.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - key string   a unique value key.
//   - refObj       pointer to object for restore
// Returns bool, error
func (c *MemoryCache) RetrieveAs(correlationId string, key string, refObj interface{}) (bool, error) {
	if key == "" {
		panic("Key cannot be empty")
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	entry := c.cache[key]
	if entry != nil {
		if entry.IsExpired() {
			delete(c.cache, key)
			return false, nil
		}
		err := json.Unmarshal((entry.Value()).([]byte), refObj)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// Stores value in the cache with expiration time, if success return stored value.
// Parameters:
//   - correlationId string
//    transaction id to trace execution through call chain.
//   - key string
//   a unique value key.
//   - value interface{}
//   a value to store.
//   - timeout int64
//   expiration timeout in milliseconds.
// Returns interface{}, error
func (c *MemoryCache) Store(correlationId string, key string, value interface{}, timeout int64) (interface{}, error) {
	if key == "" {
		panic("Key cannot be empty")
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	entry := c.cache[key]
	if timeout <= 0 {
		timeout = c.timeout
	}

	// if value == nil {
	// 	if entry != nil {
	// 		delete(c.cache, key)
	// 	}
	// 	return nil, nil
	// }

	jsonVal, err := json.Marshal(value)

	if err != nil {
		return nil, err
	}

	if entry != nil {
		entry.SetValue(jsonVal, timeout)
	} else {
		c.cache[key] = NewCacheEntry(key, jsonVal, timeout)
	}

	// cleanup
	if c.maxSize > 0 && len(c.cache) > c.maxSize {
		c.Cleanup()
	}

	return value, nil
}

// Removes a value from the cache by its key.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - key string
//   a unique value key.
// Returns error
func (c *MemoryCache) Remove(correlationId string, key string) error {
	if key == "" {
		panic("Key cannot be empty")
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.cache, key)

	return nil
}

// Clear a value from the cache.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
func (c *MemoryCache) Clear(correlationId string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache = map[string]*CacheEntry{}

	return nil
}
